package video

import (
	"github.com/gin-gonic/gin"
	"github.com/russellcxl/go-elastic-iam/pkg/types"
)

type VideoController interface {
	Save(*gin.Context) (types.Video, error)
	FindAll() []types.Video
}

type videoController struct {
	service VideoService
}

func NewController(s VideoService) VideoController {
	return &videoController {
		service: s,
	}
}

func (c *videoController) Save(ctx *gin.Context) (types.Video, error) {
	var v types.Video
	if err := ctx.ShouldBindJSON(&v); err != nil {
		return v, err
	}
	c.service.Save(v)
	return v, nil
}

func (c *videoController) FindAll() []types.Video {
	return c.service.FindAll()
}