package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.New()
	// $ go get -u github.com/jteeuwen/go-bindata/...
	// $ go-bindata ./templates/...
	// $ go build
	// $ ./embedding-templates-into-app
	// html files are not used, you can delete the folder and run the example
	app.RegisterView(ion.HTML("./templates", ".html").Binary(Asset, AssetNames))
	app.Get("/", hi)

	// http://localhost:8080
	app.Run(ion.Addr(":8080"))
}

type page struct {
	Title, Name string
}

func hi(ctx context.Context) {
	ctx.ViewData("", page{Title: "Hi Page", Name: "ion"})
	ctx.View("hi.html")
}
