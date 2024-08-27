package manager

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"open-btm.com/observe"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"open-btm.com/configs"

	"github.com/spf13/cobra"
)

var (
	btmgraphprodechocli = &cobra.Command{
		Use:   "gprod",
		Short: "Run GraphQL Echo Production Server server ",
		Long:  `Run Production server`,
		Run: func(cmd *cobra.Command, args []string) {
			graph_prod_echo()
		},
	}
)

func graph_prod_echo() {
	//  loading dev env file first
	configs.AppConfig.SetEnv("prod")

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
	goFrame.AddCommand(btmgraphprodechocli)
}
