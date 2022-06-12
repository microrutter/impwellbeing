package main

import (
	"log"
	"os"

	"git.icod.de/dalu/ginpongo2v3"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/impwellbeing/controllers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := gin.New()
	app.Use(static.Serve("/static", static.LocalFile("static", false)))
	app.HTMLRender = ginpongo2v3.New("default", "views", gin.IsDebugging())

	landingController := controllers.NewLandingController()
	resourceController := controllers.NewResourceController()
	blogController := controllers.NewBlogController()

	app.GET("/", landingController.ShowLanding)
	app.GET("/resource", resourceController.ShowResource)
	app.GET("/blog", blogController.ShowBlogIndex)
	app.GET("/blog/:postId", blogController.GetPost)

	_ = app.Run(":" + port)

}
