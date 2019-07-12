package submission

import (
	"github.com/labstack/echo/v4"

	"api/app/submission/middleware"
	"api/app/submission/controllers"
)

func Register(e *echo.Echo) {
	e.GET("submission/:id", controllers.FetchOne, middleware.IdExists)
	e.POST("submission", controllers.Submit, middleware.SubmissionIsValid)
	e.GET("submission", controllers.GetSubmit)
}