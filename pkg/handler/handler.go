package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/russellcxl/go-elastic-iam/pkg/author"
	"github.com/russellcxl/go-elastic-iam/pkg/middlewares"
	"github.com/russellcxl/go-elastic-iam/pkg/video"
)

type Handler struct {
	server           *gin.Engine
	videoController  video.VideoController
	authorController author.AuthorController
}

func HandleRoutes(s *gin.Engine) {
	h := &Handler{
		server:           s,
		videoController:  video.NewController(),
		authorController: author.NewController(),
	}
	h.handleAPIRoutes()
	h.handlePublicRoutes()
}

func (h *Handler) handleAPIRoutes() {

	// basic auth only applied to /api routes
	route := h.server.Group("/api", middlewares.Auth())

	// video routes
	route.GET("/videos", h.videoController.FindAll)
	route.POST("/save", h.videoController.Save)

	// author routes
	route.GET("/authors", h.authorController.GetAll)
	route.GET("/author", h.authorController.Get)
	route.POST("/author", h.authorController.Save)

}

func (h *Handler) handlePublicRoutes() {
	h.server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	h.server.GET("/videos", h.videoController.ShowAll)
}
