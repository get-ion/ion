package main

import (
	"os"
	"time"

	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

// get a filename based on the date, file logs works that way the most times
// but these are just a sugar.
func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFilename()
	// open an output file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

func main() {
	f := newLogFile()
	defer f.Close()

	app := ion.New()
	// attach the file as logger, remember, ion' app logger is just an io.Writer.
	app.Logger().Out = newLogFile()

	app.Get("/", func(ctx context.Context) {
		// for the sake of simplicity, in order see the logs at the ./_today_.txt
		ctx.Application().Logger().Infoln("Request path: " + ctx.Path())
		ctx.Writef("hello")
	})

	// navigate to http://localhost:8080
	// and open the ./logs.txt file
	if err := app.Run(ion.Addr(":8080"), ion.WithoutBanner); err != nil {
		if err != ion.ErrServerClosed {
			app.Logger().Warnln("Shutdown with error: " + err.Error())
		}
	}
}
