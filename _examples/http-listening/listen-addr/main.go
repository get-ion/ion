package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

func main() {
	app := ion.New()

	app.Get("/", func(ctx context.Context) {
		ctx.HTML("<h1>Hello World!/</h1>")
	})

	if err := app.Run(ion.Addr(":8080")); err != nil {
		panic(err)
	}

}
