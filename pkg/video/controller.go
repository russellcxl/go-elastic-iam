package video

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/russellcxl/go-elastic-iam/pkg/types"
	"github.com/russellcxl/go-elastic-iam/pkg/validators"
)

type VideoController interface {
	Save(*gin.Context) (*types.Video, error)
	FindAll() []types.Video
}

type videoController struct {
	service VideoService
	validator *validator.Validate
}

func NewController() VideoController {

	// add all validations for custom json validation tags e.g. `validation:"is-title-ok"`
	validate := validator.New()
	for tag, v := range validators.Validations {
		validate.RegisterValidation(tag, v)
	}

	return &videoController{
		service: NewService(),
		validator: validate,
	}
}

func (c *videoController) Save(ctx *gin.Context) (*types.Video, error) {
	var v types.Video
	if err := ctx.ShouldBindJSON(&v); err != nil {
		return nil, err
	}
	if err := c.validator.Struct(v); err != nil {
		return nil, err
	}
	c.service.Save(v)
	return &v, nil
}

func (c *videoController) FindAll() []types.Video {
	return c.service.FindAll()
}
