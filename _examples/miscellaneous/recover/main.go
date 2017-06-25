package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"

	"github.com/get-ion/ion/middleware/recover"
)

func main() {
	app := ion.New()
	// use this recover(y) middleware
	app.Use(recover.New())

	i := 0
	// let's simmilate a panic every next request
	app.Get("/", func(ctx context.Context) {
		i++
		if i%2 == 0 {
			panic("a panic here")
		}
		ctx.Writef("Hello, refresh one time more to get panic!")
	})

	// http://localhost:8080, refresh it 5-6 times.
	app.Run(ion.Addr(":8080"))
}

// Note:
// app := ion.Default() instead of ion.New() makes use of the recovery middleware automatically.
