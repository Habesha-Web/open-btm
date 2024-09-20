package manager

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"open-btm.com/configs"
	"open-btm.com/database"
	"open-btm.com/graph"
	"open-btm.com/models"
	"open-btm.com/observe"
	"open-btm.com/users"

	"github.com/spf13/cobra"
)

var (
	env                string
	btmgraphdevechocli = &cobra.Command{
		Use:   "run",
		Short: "Run GraphQL Echo Development server ",
		Long:  `Run Gofr development server`,
		Run: func(cmd *cobra.Command, args []string) {
			switch env {
			case "":
				graph_echo_run("dev")
			default:
				graph_echo_run(env)
			}
		},
	}
)

func otelechospanstarter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		routeName := ctx.Path() + "_" + strings.ToLower(ctx.Request().Method)
		tracer, span := observe.EchoAppSpanner(ctx, fmt.Sprintf("%v-root", routeName))
		ctx.Set("tracer", &observe.RouteTracer{Tracer: tracer, Span: span})

		// Process request
		err := next(ctx)
		if err != nil {
			return err
		}

		span.End()
		return nil
	}
}

func dbsessioninjection(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		db, err := database.ReturnSession()
		if err != nil {
			return err
		}
		ctx.Set("db", db)

		nerr := next(ctx)
		if nerr != nil {
			return nerr
		}

		return nil
	}
}

func graph_echo_run(env string) {
	//  loading dev env file first
	configs.AppConfig.SetEnv(env)

	// Starting Otel Global tracer
	tp := observe.InitTracer()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// starting the app
	app := echo.New()

	// Recover incase of panic attacks
	app.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	//  prometheus metrics middleware
	app.Use(echoprometheus.NewMiddleware("echo_blue"))

	// Rate Limiting to throttle overload
	app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))

	_, ok := models.LoginBlueAdmin()
	if !ok {
		panic("No Auth for the app")
	}
	SetupRoutes(app)

	// starting on provided port
	go func(app *echo.Echo) {
		//  Http serving port
		HTTP_PORT := configs.AppConfig.Get("HTTP_PORT")
		app.Logger.Fatal(app.Start("0.0.0.0:" + HTTP_PORT))
		// log.Fatal(app.ListenTLS(":" + port_1, "server.pem", "server-key.pem"))
	}(app)

	c := make(chan os.Signal, 1)   // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")

	fmt.Println("Running cleanup tasks...")
	// Your cleanup tasks go here
	fmt.Println("btm was successful shutdown.")

}

func SetupRoutes(app *echo.Echo) {

	// the Otel spanner middleware
	app.Use(otelechospanstarter)

	app.Use(middleware.BodyDump(func(ctx echo.Context, reqBody, resBody []byte) {
		//  Geting tracer
		tracer := ctx.Get("tracer").(*observe.RouteTracer)
		tracer.Span.SetAttributes(attribute.String("request", string(reqBody)))
		tracer.Span.SetAttributes(attribute.String("response", string(resBody)))
	}))

	// db session injection
	app.Use(dbsessioninjection)

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	gapp := app.Group("/api/v1")

	// playgroundHandler := playground.Handler("GraphQL", "/query")
	cfg := graph.Config{
		Directives: graph.DirectiveRoot{
			HasRole: graph.HasRoleDirective,
		},
	}

	//  This is GraphQL project graphql route
	gapp.POST("/project/:project_id", func(ctx echo.Context) error {

		// Parsing project ID
		project_id, err := strconv.Atoi(ctx.Param("project_id"))
		if err != nil {
			return err
		}

		//  Getting Project Database Session
		project_db, err := database.ReturnSession()
		if err != nil {
			return nil
		}

		// Getting Project database Name
		var project models.Project
		if res := project_db.Model(&models.Project{}).Preload(clause.Associations).Where("id = ?", uint(project_id)).First(&project); res.Error != nil {
			return res.Error
		}

		//  Connecting to Databse
		db, err := database.ReturnSessionDatabase(project.DatabaseName)
		if err != nil {
			log.Errorf("Error Connecting to Database: %v\n", err)
			return err
		}

		//  Geting tracer
		tracer := ctx.Get("tracer").(*observe.RouteTracer)

		// Providing reslover trancer and database connection
		cfg.Resolvers = &graph.Resolver{DB: db, Tracer: tracer}

		//  Schema handler
		graphqlHandler := handler.NewDefaultServer(
			graph.NewExecutableSchema(cfg),
		)
		graphqlHandler.ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	})

	//  admin graphql  config
	pcfg := users.Config{
		Directives: users.DirectiveRoot{
			HasProjectRole: users.HasProjectRoleDirective,
		},
	}

	//  This is GraphQL project graphql route
	gapp.POST("/admin", func(ctx echo.Context) error {

		//  Geting dbsession
		db := ctx.Get("db").(*gorm.DB)

		//  Geting tracer
		tracer := ctx.Get("tracer").(*observe.RouteTracer)

		// Providing reslover trancer and database connection
		pcfg.Resolvers = &users.Resolver{DB: db, Tracer: tracer}

		//  Schema handler
		graphqlHandler := handler.NewDefaultServer(
			users.NewExecutableSchema(pcfg),
		)

		graphqlHandler.ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	})

}

func init() {
	btmgraphdevechocli.Flags().StringVar(&env, "env", "help", "Which environment to run for example prod or dev")
	goFrame.AddCommand(btmgraphdevechocli)
}
