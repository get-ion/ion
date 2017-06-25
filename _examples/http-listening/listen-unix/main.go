package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/core/nettools"
)

func main() {
	app := ion.New()

	l, err := nettools.UNIX("/tmpl/srv.sock", 0666) // see its code to see how you can manually create a new file listener, it's easy.
	if err != nil {
		panic(err)
	}

	app.Run(ion.Listener(l))
}
