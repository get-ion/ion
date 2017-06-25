package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.New() // defaults to these

	// - standard html  | view.HTML(...)
	// - django         | view.Django(...)
	// - pug(jade)      | view.Pug(...)
	// - handlebars     | view.Handlebars(...)
	// - amber          | view.Amber(...)

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
