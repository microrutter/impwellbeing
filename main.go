package main

import (
	"git.icod.de/dalu/ginpongo2v3"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/impwellbeing/controllers"
)

func main() {
	app := gin.Default()
	app.Use(static.Serve("/static", static.LocalFile("static", false)))
	app.HTMLRender = ginpongo2v3.New("default", "views", gin.IsDebugging())

	landingController := controllers.NewLandingController()
	resourceController := controllers.NewResourceController()

	app.GET("/", landingController.ShowLanding)
	app.GET("/resource", resourceController.ShowResource)

	_ = app.Run()

}
