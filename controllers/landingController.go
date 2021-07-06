package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/flosch/pongo2.v3"
	"github.com/nu7hatch/gouuid"
)

type LandingController struct {
}

func NewLandingController() *LandingController {
	return &LandingController{}
}

func (c *LandingController) ShowLanding(ctx *gin.Context) {
	uuid, _ := uuid.NewV4()

	ctx.HTML(http.StatusOK, "landing", pongo2.Context{
		"context": uuid.String(),
	})
}
