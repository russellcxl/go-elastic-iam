package video

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/russellcxl/go-elastic-iam/pkg/types"
	"github.com/russellcxl/go-elastic-iam/pkg/validators"
)

type VideoController interface {
	Save(*gin.Context)
	FindAll(*gin.Context)
	ShowAll(*gin.Context)
}

type videoController struct {
	service   VideoService
	validator *validator.Validate
}

func NewController() VideoController {

	// add all validations for custom json validation tags e.g. `validation:"is-title-ok"`
	validate := validator.New()
	for tag, v := range validators.Validations {
		validate.RegisterValidation(tag, v)
	}

	return &videoController{
		service:   NewService(),
		validator: validate,
	}
}

func (c *videoController) Save(ctx *gin.Context) {
	var v types.Video
	if err := ctx.ShouldBindJSON(&v); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := c.validator.Struct(v); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newV, err := c.service.Save(v)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, newV)
}

func (c *videoController) FindAll(ctx *gin.Context) {
	videos, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, videos)
}

func (c *videoController) ShowAll(ctx *gin.Context) {
	videos, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
