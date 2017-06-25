package main

import (
	"github.com/get-ion/ion"
)

func main() {
	app := ion.New()

	app.Favicon("./assets/favicon.ico")

	// first parameter is the request path
	// second is the system directory
	//
	// app.StaticWeb("/css", "./assets/css")
	// app.StaticWeb("/js", "./assets/js")
	//
	app.StaticWeb("/static", "./assets")

	// http://localhost:8080/static/css/main.css
	// http://localhost:8080/static/js/jquery-2.1.1.js
	// http://localhost:8080/static/favicon.ico
	app.Run(ion.Addr(":8080"))

	// Note:
	// Routing doesn't allows something .StaticWeb("/", "./assets")
	//
	// To see how you can wrap the router in order to achieve
	// wildcard on root path, see "single-page-application".
}
