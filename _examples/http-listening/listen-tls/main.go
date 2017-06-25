package main

import (
	"net/url"

	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"

	"github.com/get-ion/ion/core/host"
)

func main() {
	app := ion.New()

	app.Get("/", func(ctx context.Context) {
		ctx.Writef("Hello from the SECURE server")
	})

	app.Get("/mypath", func(ctx context.Context) {
		ctx.Writef("Hello from the SECURE server on path /mypath")
	})

	// to start a new server listening at :80 and redirects
	// to the secure address, then:
	target, _ := url.Parse("https://127.0.1:443")
	go host.NewProxy("127.0.0.1:80", target).ListenAndServe()

	// start the server (HTTPS) on port 443, this is a blocking func
	app.Run(ion.TLS("127.0.0.1:443", "mycert.cert", "mykey.key"))

}
