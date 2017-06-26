package main

import (
	"sync"

	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

// Context is our custom context.
type Context struct {
	context.Context
}

// Bold will send a bold text to the client.
func (ctx *Context) Bold(text string) {
	ctx.HTML("<b>" + text + "</b>")
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

func acquire(original context.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original
	return ctx
}

func release(ctx *Context) {
	contextPool.Put(ctx)
}

// Handler will convert our handler of func(*Context) to an ion Handler,
// in order to be compatible with the HTTP API.
func Handler(h func(*Context)) context.Handler {
	return func(original context.Context) {
		ctx := acquire(original)
		h(ctx)
		release(ctx)
	}
}

func main() {
	app := ion.New()

	app.Get("/", Handler(func(ctx *Context) {
		ctx.Bold("Hello from our *Context")
	}))

	app.Run(ion.Addr(":8080"))
}
