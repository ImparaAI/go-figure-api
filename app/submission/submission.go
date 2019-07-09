package submission

import (
	"github.com/labstack/echo/v4"
	"api/app/submission/controllers"
	"api/app/submission/middleware"
)

func Register(e *echo.Echo) {
	e.GET("submission/:id", controllers.FetchOne, middleware.IdExists)
	e.POST("submission", controllers.Submit, middleware.SubmissionIsValid)
}