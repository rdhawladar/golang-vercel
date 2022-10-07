package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "golang-vercel/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	app *gin.Engine
)

// @title Golang Vercel Deployment
// @description API Documentation for Golang deployment in vercel serverless environment
// @version 1.0

// @schemes https http
// @host golang-vercel.vercel.app
func init() {
	app = gin.New()
	app.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.GET("/ping", Ping)
	app.GET("/hello/:name", Hello)
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

// @Tags        Welcome
// @Summary     Hello User
// @Description Endpoint to Welcome user and say Hello "Name"
// @Param       name query string true "Name in the URL param"
// @Accept      json
// @Produce     json
// @Success     200 {object} object "success"
// @Failure     400 {object} object "Request Error or parameter missing"
// @Failure     404 {object} object "When user not found"
// @Failure     500 {object} object "Server Error"
// @Router      /hello/:name [GET]
func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello %v", c.Param("name"))
}
