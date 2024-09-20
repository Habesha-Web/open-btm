package tests

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"open-btm.com/manager"
)

var (
	TestApp   *echo.Echo
	groupPath = "/api/v1"
)

func setupTestApp() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}
	TestApp = echo.New()
	manager.SetupRoutes(TestApp)
}
