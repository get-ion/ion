package main

import (
	"testing"

	"github.com/get-ion/ion/httptest"
)

func TestCustomContextNewImpl(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	e.GET("/").Expect().
		Status(httptest.StatusOK).
		ContentType("text/html").
		Body().Equal("<b>Hello from our *Context</b>")
}
