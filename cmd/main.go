package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
		gindump.Dump(),       // more logging info
	)

	// specify routes
	handleApiRoutes()
	handleViewRoutes()
	
	godotenv.Load(".env")
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "5000"
	}
	server.Run(":" + port)
}

func handleApiRoutes() {

	// basic auth only applied to /api routes
	route := server.Group("/api", middlewares.Auth())

	// define routes
	route.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, controller.FindAll())
	})
	route.POST("/save", func(ctx *gin.Context) {
		v, err := controller.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, v)
	})
}

func handleViewRoutes() {
	server.GET("/videos", controller.ShowAll)
}
