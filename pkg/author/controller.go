package author

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/russellcxl/go-elastic-iam/pkg/types"
)

type AuthorController interface {
	Save(*gin.Context)
	Get(*gin.Context)
	GetAll(*gin.Context) 
}

type authorController struct {
	service AuthorService
}

func NewController() AuthorController {
	return &authorController{
		service: newService(),
	}
}

func (c *authorController) Save(ctx *gin.Context) {
	var p types.Author
	if err := ctx.ShouldBindJSON(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newP, err := c.service.Save(p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, newP)
}

func (c *authorController) Get(ctx *gin.Context) {
	var req types.GetAuthorRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	author, err := c.service.Get(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, author)
}

func (c *authorController) GetAll(ctx *gin.Context) {
	authors, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, authors)
}