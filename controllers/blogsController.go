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

func NewBlogController() *ResourceController {
	return &ResourceController{}
}

func (c *ResourceController) ShowBlogIndex(ctx *gin.Context) {
	uuid, _ := uuid.NewV4()
	ResourceList := graphcms.GetAllBlogs(ctx)

	ctx.HTML(http.StatusOK, "resource/index", pongo2.Context{
		"context":   uuid.String(),
		"resources": ResourceList.Blogs,
	})
}
