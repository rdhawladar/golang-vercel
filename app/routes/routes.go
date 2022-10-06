package routes

import (
	"golang-vercel/app/handler"

	_ "golang-vercel/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine) {
	app.NoRoute(ErrRouter)

	app.GET("/ping", handler.Ping)

	route := app.Group("/api")
	{
		route.GET("/hello/:name", handler.Hello)
		route.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func ErrRouter(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": "this page could not be found",
	})
}
