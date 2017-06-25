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

	// Good when you want to modify the whole configuration.
	app.Run(ion.Addr(":8080"), ion.WithConfiguration(ion.Configuration{ // default configuration:
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))

	// or before Run:
	// app.Configure(ion.WithConfiguration(ion.Configuration{...}))
}
