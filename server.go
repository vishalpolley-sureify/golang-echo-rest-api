package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"test-api/controllers"
)

func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Welcome mvc echo with mysql app using Golang")
	})

	e.GET("/employee", controllers.GetEmployees)

	e.GET("/employee/:id", controllers.GetEmployee)

	e.POST("/employee", controllers.PostEmployee)

	e.DELETE("/employee/delete/:id", controllers.DeleteEmployee)

	e.Logger.Fatal(e.Start(":8081"))
}