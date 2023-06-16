package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/russellcxl/go-elastic-iam/pkg/middlewares"
	"github.com/russellcxl/go-elastic-iam/pkg/video"
	gindump "github.com/tpkeeper/gin-dump"
)

var server *gin.Engine
var controller video.VideoController

func main() {
	service := video.NewService()
	controller = video.NewController(service)

	// initialise gin engine
	server = gin.New()
	server.Static("../templates/css", "./templates/css")
	server.LoadHTMLGlob("./templates/*.html")
	middlewares.SetLogOutput("logs/data.log") // writes to specified file and console

	// add middlewares
	server.Use(
		gin.Recovery(),
		middlewares.Logger(), // custom logger
		middlewares.Auth(),   // basic auth
		gindump.Dump(),       // more logging info
	)

	// specify routes
	handleApiRoutes()
	handleViewRoutes()
	server.Run(":8080")
}

func handleApiRoutes() {
	route := server.Group("/api")
	{
		route.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, controller.FindAll())
		})
	
		route.POST("/save", func(ctx *gin.Context) {
			v, err := controller.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H {
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, v)
		})
	}
}

func handleViewRoutes() {
	route := server.Group("/view")
	{
		route.GET("/videos", controller.ShowAll)
	}
}
