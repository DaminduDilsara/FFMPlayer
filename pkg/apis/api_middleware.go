package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func EchoManager(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	default_ := e.Group("/mediaplayer")
	defaultSecured_ := e.Group("/mediaplayer")

	defaultSecured_.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})

	FN_Default(default_)
}

func FN_Default(g *echo.Group) {

	g.POST("/api/upload/file", UploadFileApiController)
	g.GET("/api/all/videos", GetAllFilesApiController)

}
