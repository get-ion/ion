package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
	"github.com/get-ion/ion/view"
)

// same as embedded-single-page-application but without go-bindata, the files are "physical" stored in the
// current system directory.

var page = struct {
	Title string
}{"Welcome"}

func newApp() *ion.Application {
	app := ion.New()
	app.RegisterView(view.HTML("./public", ".html"))

	app.Get("/", func(ctx context.Context) {
		ctx.ViewData("Page", page)
		ctx.View("index.html")
	})

	// or just serve index.html as it is:
	// app.Get("/", func(ctx context.Context) {
	// 	ctx.ServeFile("index.html", false)
	// })

	assetHandler := app.StaticHandler("./public", false, false)
	app.SPA(assetHandler)

	return app
}

func main() {
	app := newApp()

	// http://localhost:8080
	// http://localhost:8080/index.html
	// http://localhost:8080/app.js
	// http://localhost:8080/css/main.css
	app.Run(ion.Addr(":8080"))
}
