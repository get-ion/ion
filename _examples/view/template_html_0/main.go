package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.New() // defaults to these

	// - standard html  | ion.HTML(...)
	// - django         | ion.Django(...)
	// - pug(jade)      | ion.Pug(...)
	// - handlebars     | ion.Handlebars(...)
	// - amber          | ion.Amber(...)

	tmpl := ion.HTML("./templates", ".html")
	tmpl.Reload(true) // reload templates on each request (development mode)
	// default template funcs are:
	//
	// - {{ urlpath "mynamedroute" "pathParameter_ifneeded" }}
	// - {{ render "header.html" }}
	// - {{ render_r "header.html" }} // partial relative path to current page
	// - {{ yield }}
	// - {{ current }}
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})
	app.RegisterView(tmpl)

	app.Get("/", hi)

	// http://localhost:8080
	app.Run(ion.Addr(":8080"), ion.WithCharset("UTF-8")) // defaults to that but you can change it.
}

func hi(ctx context.Context) {
	ctx.ViewData("Title", "Hi Page")
	ctx.ViewData("Name", "ion") // {{.Name}} will render: ion
	// ctx.ViewData("", myCcustomStruct{})
	ctx.View("hi.html")
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
