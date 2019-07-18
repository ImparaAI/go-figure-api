package drawing

import (
	"github.com/labstack/echo/v4"

	"api/app/drawing/middleware"
	"api/app/drawing/controllers"
)

func Register(e *echo.Echo) {
	e.GET("drawing/:id", controllers.FetchOne, middleware.IdExists)
	e.POST("drawing", controllers.Submit, middleware.SubmissionIsValid)
}