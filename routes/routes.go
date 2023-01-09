package routes

import (
	"github.com/labstack/echo/v4"
	"learn-echo-mongo/controllers"
	"net/http"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hallo this is echo")
	})
	e.GET("/api/users", controllers.GetUsers)
	e.GET("/api/user", controllers.GetUserById)
	e.POST("/api/user", controllers.CreateUsers)

	return e
}
