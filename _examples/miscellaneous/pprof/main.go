package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"

	"github.com/get-ion/ion/middleware/pprof"
)

func main() {
	app := ion.New()

	app.Get("/", func(ctx context.Context) {
		ctx.HTML("<h1> Please click <a href='/debug/pprof'>here</a>")
	})

	app.Any("/debug/pprof/{action:path}", pprof.New())
	//                              ___________
	app.Run(ion.Addr(":8080"))
}
