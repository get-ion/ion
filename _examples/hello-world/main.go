// +build go.1.8

package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.Default()

	// Method: GET
	// Resource: http://localhost:8080/
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.HTML("<b>Hello world!</b>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method: GET
	// Resource: http://context:8080/ping
	app.Get("/ping", func(ctx ion.Context) {
		ctx.WriteString("pong")
	})

	// Method: GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx context.Context) {
		ctx.JSON(ion.Map{"message": "Hello ion web framework."})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(ion.Addr(":8080"))
}
