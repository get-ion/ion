package main

import (
	stdContext "context"
	"time"

	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

// Before continue:
//
// Gracefully Shutdown on control+C/command+C or when kill command sent is ENABLED BY-DEFAULT.
//
// In order to manually manage what to do when app is interrupted,
// We have to disable the default behavior with the option `WithoutInterruptHandler`
// and register a new interrupt handler (globally, across all possible hosts).
func main() {
	app := ion.New()

	ion.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		// close all hosts
		app.Shutdown(ctx)
	})

	app.Get("/", func(ctx context.Context) {
		ctx.HTML(" <h1>hi, I just exist in order to see if the server is closed</h1>")
	})

	// http://localhost:8080
	app.Run(ion.Addr(":8080"), ion.WithoutInterruptHandler)
}
