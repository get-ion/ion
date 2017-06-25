// black-box testing
package handlerconv_test

import (
	"net/http"
	"testing"

	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
	"github.com/get-ion/ion/core/handlerconv"
	"github.com/get-ion/ion/httptest"
)

func TestFromStd(t *testing.T) {
	expected := "ok"
	std := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expected))
	}

	h := handlerconv.FromStd(http.HandlerFunc(std))

	hFunc := handlerconv.FromStd(std)

	app := ion.New()
	app.Get("/handler", h)
	app.Get("/func", hFunc)

	e := httptest.New(t, app)

	e.GET("/handler").
		Expect().Status(ion.StatusOK).Body().Equal(expected)

	e.GET("/func").
		Expect().Status(ion.StatusOK).Body().Equal(expected)
}

func TestFromStdWithNext(t *testing.T) {

	basicauth := "secret"
	passed := "ok"

	stdWNext := func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if username, password, ok := r.BasicAuth(); ok &&
			username == basicauth && password == basicauth {
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(ion.StatusForbidden)
	}

	h := handlerconv.FromStdWithNext(stdWNext)
	next := func(ctx context.Context) {
		ctx.WriteString(passed)
	}

	app := ion.New()
	app.Get("/handlerwithnext", h, next)

	e := httptest.New(t, app)

	e.GET("/handlerwithnext").
		Expect().Status(ion.StatusForbidden)

	e.GET("/handlerwithnext").WithBasicAuth(basicauth, basicauth).
		Expect().Status(ion.StatusOK).Body().Equal(passed)
}
