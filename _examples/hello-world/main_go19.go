// +build go1.9

package main

import (
	"github.com/get-ion/ion"
)

func main() {
	app := ion.Default()
	app.Handle("GET", "/", func(ctx ion.Context) {
		ctx.HTML("<b>Hello world!</b>")
	})
	app.Run(ion.Addr(":8080"))
}
