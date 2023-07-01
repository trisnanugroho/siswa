package routes

import (
	"siswa/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo {
	e.GET("/students", controllers.SiswaController)             // endpoint 1
	e.GET("/students/:id", controllers.DetailSiswaController)   // endpoint 2
	e.POST("/students", controllers.TambahSiswaController)      // endpoint 3
	e.DELETE("/students/:id", controllers.HapusSiswaController) // endpoint 4
	return e
}
