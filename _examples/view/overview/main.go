package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.New()

	// - standard html  | ion.HTML(...)
	// - django         | ion.Django(...)
	// - pug(jade)      | ion.Pug(...)
	// - handlebars     | ion.Handlebars(...)
	// - amber          | ion.Amber(...)
	// with default template funcs:
	//
	// - {{ urlpath "mynamedroute" "pathParameter_ifneeded" }}
	// - {{ render "header.html" }}
	// - {{ render_r "header.html" }} // partial relative path to current page
	// - {{ yield }}
	// - {{ current }}
	app.RegisterView(ion.HTML("./templates", ".html"))
	app.Get("/", func(ctx context.Context) {

		ctx.ViewData("Name", "ion") // the .Name inside the ./templates/hi.html
		ctx.Gzip(true)              // enable gzip for big files
		ctx.View("hi.html")         // render the template with the file name relative to the './templates'

	})

	// http://localhost:8080/
	app.Run(ion.Addr(":8080"))
}

/*
Note:

In case you're wondering, the code behind the view engines derives from the "github.com/get-ion/ion/view" package,
access to the engines' variables can be granded by "github.com/get-ion/ion" package too.

    ion.HTML(...) is a shortcut of view.HTML(...)
    ion.Django(...)     >> >>      view.Django(...)
    ion.Pug(...)        >> >>      view.Pug(...)
    ion.Handlebars(...) >> >>      view.Handlebars(...)
    ion.Amber(...)      >> >>      view.Amber(...)
*/
