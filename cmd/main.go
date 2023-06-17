package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/russellcxl/go-elastic-iam/pkg/handler"
	"github.com/russellcxl/go-elastic-iam/pkg/middlewares"
	gindump "github.com/tpkeeper/gin-dump"
)

var server *gin.Engine

func main() {

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

	// map routes
	handler.HandleRoutes(server)

	godotenv.Load(".env")
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "5000"
	}
	server.Run(":" + port)
}
