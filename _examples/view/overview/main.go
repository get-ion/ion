package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.New()

	// - standard html  | view.HTML(...)
	// - django         | view.Django(...)
	// - pug(jade)      | view.Pug(...)
	// - handlebars     | view.Handlebars(...)
	// - amber          | view.Amber(...)
	// with default template funcs:
	//
	// - {{ urlpath "mynamedroute" "pathParameter_ifneeded" }}
	// - {{ render "header.html" }}
	// - {{ render_r "header.html" }} // partial relative path to current page
	// - {{ yield }}
	// - {{ current }}
	//
	// ion.HTML is a conversion for ion/view/html.go#HTML
	app.RegisterView(ion.HTML("./templates", ".html"))
	app.Get("/", func(ctx context.Context) {

		ctx.ViewData("Name", "ion") // the .Name inside the ./templates/hi.html
		ctx.Gzip(true)              // enable gzip for big files
		ctx.View("hi.html")         // render the template with the file name relative to the './templates'

	})

	// http://localhost:8080/
	app.Run(ion.Addr(":8080"))
}
