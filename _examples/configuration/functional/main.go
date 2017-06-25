package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.New()
	app.Get("/", func(ctx context.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
	// [...]

	// Good when you want to change some of the configuration's field.
	// Prefix: "With", code editors will help you navigate through all
	// configuration options without even a glitch to the documentation.

	app.Run(ion.Addr(":8080"), ion.WithoutBanner, ion.WithCharset("UTF-8"))

	// or before run:
	// app.Configure(ion.WithoutBanner, ion.WithCharset("UTF-8"))
	// app.Run(ion.Addr(":8080"))
}
