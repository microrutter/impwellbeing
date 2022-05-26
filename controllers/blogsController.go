package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
	"gopkg.in/flosch/pongo2.v3"

	"github.com/impwellbeing/services/graphcms"
)

type BlogController struct {
}

func NewBlogController() *BlogController {
	return &BlogController{}
}

func (c *BlogController) ShowBlogIndex(ctx *gin.Context) {
	uuid, _ := uuid.NewV4()
	ResourceList := graphcms.GetAllBlogs(ctx)

	ctx.HTML(http.StatusOK, "blog/index", pongo2.Context{
		"context": uuid.String(),
		"blogs":   ResourceList.Blogs,
	})
}
