package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

type Company struct {
	Name  string
	City  string
	Other string
}

func MyHandler(ctx context.Context) {
	c := &Company{}
	if err := ctx.ReadJSON(c); err != nil {
		ctx.StatusCode(ion.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Writef("Received: %#v\n", c)
}

func main() {
	app := ion.New()

	app.Post("/", MyHandler)

	// use Postman or whatever to do a POST request
	// to the http://localhost:8080 with RAW BODY:
	/*
		{
			"Name": "ion-Go",
			"City": "New York",
			"Other": "Something here"
		}
	*/
	// and Content-Type to application/json
	//
	// The response should be:
	// Received: &main.Company{Name:"ion-Go", City:"New York", Other:"Something here"}
	app.Run(ion.Addr(":8080"))
}
