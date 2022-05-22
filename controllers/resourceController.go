package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/flosch/pongo2.v3"
	"github.com/nu7hatch/gouuid"

	"github.com/impwellbeing/services/graphcms"
)

type ResourceController struct {
}

func NewResourceController() *ResourceController {
	return &ResourceController{}
}

func (c *ResourceController) ShowResource(ctx *gin.Context) {
	uuid, _ := uuid.NewV4()
	ResourceList := graphcms.ResourceCards(ctx)

	ctx.HTML(http.StatusOK, "resource/index", pongo2.Context{
		"context": uuid.String(),
		"resources": ResourceList.Resources,
	})
}
