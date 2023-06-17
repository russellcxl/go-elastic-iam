package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/russellcxl/go-elastic-iam/pkg/middlewares"
	"github.com/russellcxl/go-elastic-iam/pkg/video"
)

type Handler struct {
	server *gin.Engine
	controller video.VideoController
}

func HandleRoutes(s *gin.Engine) {
	h := &Handler{
		server: s,
		controller: video.NewController(),
	}
	h.handleAPIRoutes()
	h.handleViewRoutes()
}

func (h *Handler) handleAPIRoutes() {

	// basic auth only applied to /api routes
	route := h.server.Group("/api", middlewares.Auth())

	// define routes
	route.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, h.controller.FindAll())
	})
	route.POST("/save", func(ctx *gin.Context) {
		v, err := h.controller.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, v)
	})
}

func (h *Handler) handleViewRoutes() {
	h.server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	h.server.GET("/videos", func(ctx *gin.Context) {
		videos := h.controller.FindAll()
		data := gin.H {
			"title": "Video Page",
			"videos": videos,
		}
		ctx.HTML(http.StatusOK, "index.html", data)
	})
}