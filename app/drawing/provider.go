package drawing

import (
	"github.com/labstack/echo/v4"

	"api/app/drawing/controllers"
	"api/app/drawing/middleware"
	"api/app/drawing/processing"
)

func Register(e *echo.Echo) {
	registerRoutes(e)

	processing.PrepareQueues()
}

func registerRoutes(e *echo.Echo) {
	e.GET("drawing/:id", controllers.FetchOne, middleware.IdExists)
	e.GET("drawings/recent", controllers.FetchRecent)
	e.POST("drawing", controllers.Submit, middleware.SubmissionIsValid)
}
