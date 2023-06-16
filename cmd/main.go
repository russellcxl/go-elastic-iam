package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/russellcxl/go-elastic-iam/pkg/middlewares"
	"github.com/russellcxl/go-elastic-iam/pkg/video"
	gindump "github.com/tpkeeper/gin-dump"
)

func main() {
	service := video.NewService()
	controller := video.NewController(service)
	server := gin.New()
	middlewares.SetLogOutput("logs/data.log") // writes to specified file and console
	server.Use(
		gin.Recovery(),
		middlewares.Logger(), // custom logger
		middlewares.Auth(),   // basic auth
		gindump.Dump(),       // more logging info
	)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, controller.FindAll())
	})

	server.POST("/save", func(ctx *gin.Context) {
		v, err := controller.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		ctx.JSON(http.StatusOK, v)
	})

	server.Run(":8080")
}
