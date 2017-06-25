package host

import (
	"fmt"
	"io"
	"runtime"
)

// WriteStartupLog is a task which accepts a logger(io.Writer)
// and logs the listening address
// by a generated message based on the host supervisor's server and writes it to the "w".
// This task runs on serve.
func WriteStartupLog(w io.Writer) TaskRunnerFunc {
	return func(proc TaskProcess) {
		listeningURI := proc.Host().HostURL()
		interruptkey := "CTRL"
		if runtime.GOOS == "darwin" {
			interruptkey = "CMD"
		}
		w.Write([]byte(fmt.Sprintf("Now listening on: %s\nApplication started. Press %s+C to shut down.\n",
			listeningURI, interruptkey)))
	}
}
