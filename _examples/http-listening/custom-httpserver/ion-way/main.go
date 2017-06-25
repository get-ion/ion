package main

import (
	"net/http"

	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.New()

	app.Get("/", func(ctx context.Context) {
		ctx.Writef("Hello from the server")
	})

	app.Get("/mypath", func(ctx context.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})

	// Any custom fields here. Handler and ErrorLog are setted to the server automatically
	srv := &http.Server{Addr: ":8080"}

	// http://localhost:8080/
	// http://localhost:8080/mypath
	app.Run(ion.Server(srv)) // same as app.Run(ion.Addr(":8080"))

	// More:
	// see "multi" if you need to use more than one server at the same app.
	//
	// for a custom listener use: ion.Listener(net.Listener) or
	// ion.TLS(cert,key) or ion.AutoTLS(), see "custom-listener" example for those.
}
