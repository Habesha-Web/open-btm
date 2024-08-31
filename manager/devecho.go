package manager

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.opentelemetry.io/otel/attribute"
	"open-btm.com/configs"
	"open-btm.com/database"
	"open-btm.com/graph"
	"open-btm.com/observe"

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

		span.SetAttributes(attribute.String("response", string(ctx.Response().Status)))
		span.End()
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

	// the Otel spanner middleware
	app.Use(otelechospanstarter)

	//  prometheus metrics middleware
	app.Use(echoprometheus.NewMiddleware("echo_blue"))

	// Rate Limiting to throttle overload
	app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))

	// Recover incase of panic attacks
	app.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	setupRoutes(app)
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

func init() {
	btmgraphdevechocli.Flags().StringVar(&env, "env", "help", "Which environment to run for example prod or dev")
	goFrame.AddCommand(btmgraphdevechocli)
}

func setupRoutes(app *echo.Echo) {
	gapp := app.Group("/api/v1")

	// playgroundHandler := playground.Handler("GraphQL", "/query")

	gapp.POST("/admin", func(ctx echo.Context) error {
		//  Connecting to Databse
		db, err := database.ReturnSession()
		if err != nil {
			log.Errorf("Error Connecting to Database: %v\n", err)
		}
		//  Geting tracer
		tracer := ctx.Get("tracer").(*observe.RouteTracer)

		//  Schema handler
		graphqlHandler := handler.NewDefaultServer(
			graph.NewExecutableSchema(
				graph.Config{Resolvers: &graph.Resolver{DB: db, Tracer: tracer}},
			),
		)
		graphqlHandler.ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	})

	// gapp.GET("/playground", func(ctx echo.Context) error {
	// 	playgroundHandler.ServeHTTP(ctx.Response(), ctx.Request())
	// 	return nil
	// })
}
