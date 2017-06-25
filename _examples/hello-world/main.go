package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.Default()
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.HTML("<b>Hello world!</b>")
	})
	app.Run(ion.Addr(":8080"))
}
