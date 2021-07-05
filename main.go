package impwellbeing

import (
	"git.icod.de/dalu/ginpongo2v3"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func StartImpWellbeing() error {
	app := gin.Default()
	app.Use(static.Serve("/static", static.LocalFile("selfservice/static", false)))
	app.HTMLRender = ginpongo2v3.New("default", "views", gin.IsDebugging())

	// app.GET("/auth", authController.ShowLogin)

	_ = app.Run()

	return nil
}
