package ion

import (
	// std packages
	stdContext "context"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	// context for the handlers
	"github.com/get-ion/ion/context"
	// core packages, needed to build the application
	"github.com/get-ion/ion/core/errors"
	"github.com/get-ion/ion/core/host"
	"github.com/get-ion/ion/core/netutil"
	"github.com/get-ion/ion/core/router"
	// view
	"github.com/get-ion/ion/view"
	// middleware used in Default method
	requestLogger "github.com/get-ion/ion/middleware/logger"
	"github.com/get-ion/ion/middleware/recover"
)

const (

	// Version is the current version number of the ion micro web framework.
	Version = "1.0.0"
)

// Application is responsible to manage the state of the application.
// It contains and handles all the necessary parts to create a fast web server.
type Application struct {
	Scheduler host.Scheduler

	// routing embedded | exposing APIBuilder's and Router's public API.
	*router.APIBuilder
	*router.Router
	ContextPool *context.Pool

	// config contains the configuration fields
	// all fields defaults to something that is working, developers don't have to set it.
	config *Configuration

	// the logrus logger instance, defaults to "Info" level messages (all except "Debug")
	logger *logrus.Logger

	// view engine
	view view.View
	// used for build
	once sync.Once

	mu       sync.Mutex
	Shutdown func(stdContext.Context) error
}

// New creates and returns a fresh empty ion *Application instance.
func New() *Application {
	config := DefaultConfiguration()

	app := &Application{
		config:     &config,
		logger:     logrus.New(),
		APIBuilder: router.NewAPIBuilder(),
		Router:     router.NewRouter(),
	}

	app.ContextPool = context.New(func() context.Context {
		return context.NewContext(app)
	})

	return app
}

// Default returns a new Application instance which, unlike `New`,
// recovers on panics and logs the incoming http requests.
func Default() *Application {
	app := New()
	app.Use(recover.New())
	app.Use(requestLogger.New())

	return app
}

// Configure can called when modifications to the framework instance needed.
// It accepts the framework instance
// and returns an error which if it's not nil it's printed to the logger.
// See configuration.go for more.
//
// Returns itself in order to be used like `app:= New().Configure(...)`
func (app *Application) Configure(configurators ...Configurator) *Application {
	for _, cfg := range configurators {
		cfg(app)
	}

	return app
}

// ConfigurationReadOnly returns an object which doesn't allow field writing.
func (app *Application) ConfigurationReadOnly() context.ConfigurationReadOnly {
	return app.config
}

// These are the different logging levels. You can set the logging level to log
// on the application 's instance of logger, obtained with `app.Logger()`.
//
// These are conversions from logrus.
const (
	// NoLog level, logs nothing.
	// It's the logrus' `PanicLevel` but it never used inside ion so it will never log.
	NoLog = logrus.PanicLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = logrus.ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel = logrus.WarnLevel
)

// Logger returns the logrus logger instance(pointer) that is being used inside the "app".
func (app *Application) Logger() *logrus.Logger {
	return app.logger
}

var (
	// HTML view engine.
	// Conversion for the view.HTML.
	HTML = view.HTML
	// Django view engine.
	// Conversion for the view.Django.
	Django = view.Django
	// Handlebars view engine.
	// Conversion for the view.Handlebars.
	Handlebars = view.Handlebars
	// Pug view engine.
	// Conversion for the view.Pug.
	Pug = view.Pug
	// Amber view engine.
	// Conversion for the view.Amber.
	Amber = view.Amber
)

// NoLayout to disable layout for a particular template file
// Conversion for the view.NoLayout.
const NoLayout = view.NoLayout

// RegisterView should be used to register view engines mapping to a root directory
// and the template file(s) extension.
func (app *Application) RegisterView(viewEngine view.Engine) {
	app.view.Register(viewEngine)
}

// View executes and writes the result of a template file to the writer.
//
// First parameter is the writer to write the parsed template.
// Second parameter is the relative, to templates directory, template filename, including extension.
// Third parameter is the layout, can be empty string.
// Forth parameter is the bindable data to the template, can be nil.
//
// Use context.View to render templates to the client instead.
// Returns an error on failure, otherwise nil.
func (app *Application) View(writer io.Writer, filename string, layout string, bindingData interface{}) error {
	if app.view.Len() == 0 {
		err := errors.New("view engine is missing, use `RegisterView`")
		app.Logger().Errorln(err)
		return err
	}

	err := app.view.ExecuteWriter(writer, filename, layout, bindingData)
	if err != nil {
		app.Logger().Errorln(err)
	}
	return err
}

// SPA  accepts an "assetHandler" which can be the result of an
// app.StaticHandler or app.StaticEmbeddedHandler.
// It wraps the router and checks:
// if it;s an asset, if the request contains "." (this behavior can be changed via /core/router.NewSPABuilder),
// if the request is index, redirects back to the "/" in order to let the root handler to be executed,
// if it's not an asset then it executes the router, so the rest of registered routes can be
// executed without conflicts with the file server ("assetHandler").
//
// Use that instead of `StaticWeb` for root "/" file server.
//
// Example: https://github.com/get-ion/ion/tree/master/_examples/file-server/single-page-application
func (app *Application) SPA(assetHandler context.Handler) {
	s := router.NewSPABuilder(assetHandler)
	wrapper := s.BuildWrapper(app.ContextPool)
	app.Router.WrapRouter(wrapper)
}

// NewHost accepts a standar *http.Server object,
// completes the necessary missing parts of that "srv"
// and returns a new, ready-to-use, host (supervisor).
func (app *Application) NewHost(srv *http.Server) *host.Supervisor {
	app.mu.Lock()
	defer app.mu.Unlock()

	// set the server's handler to the framework's router
	if srv.Handler == nil {
		srv.Handler = app.Router
	}

	// check if different ErrorLog provided, if not bind it with the framework's logger
	if srv.ErrorLog == nil {
		srv.ErrorLog = log.New(app.logger.Out, "[HTTP Server] ", 0)
	}

	if srv.Addr == "" {
		srv.Addr = ":8080"
	}

	// create the new host supervisor
	// bind the constructed server and return it
	su := host.New(srv)

	if app.config.vhost == "" { // vhost now is useful for router subdomain on wildcard subdomains,
		// in order to correct decide what to do on:
		// mydomain.com -> invalid
		// localhost -> invalid
		// sub.mydomain.com -> valid
		// sub.localhost -> valid
		// we need the host (without port if 80 or 443) in order to validate these, so:
		app.config.vhost = netutil.ResolveVHost(srv.Addr)
	}
	// the below schedules some tasks that will run among the server

	// I was thinking to have them on Default or here and if user not wanted these, could use a custom core/host
	// but that's too much for someone to just disable the banner for example,
	// so I will bind them to a configuration field, although is not direct to the *Application,
	// host is de-coupled from *Application as the other features too, it took me 2 months for this design.

	// copy the registered schedule tasks from the scheduler, if any will be copied to this host supervisor's scheduler.
	app.Scheduler.CopyTo(&su.Scheduler)

	if !app.config.DisableStartupLog {
		// show the banner and the available keys to exit from app.
		su.Schedule(host.WriteStartupLog(app.logger.Out)) // app.logger.Writer -> Info
	}

	if !app.config.DisableInterruptHandler {
		// give 5 seconds to the server to wait for the (idle) connections.
		shutdownTimeout := 5 * time.Second

		// when CTRL+C/CMD+C pressed.
		su.Schedule(host.ShutdownOnInterruptTask(shutdownTimeout))
	}

	if app.Shutdown == nil {
		app.Shutdown = su.Shutdown
	}

	return su
}

// Runner is just an interface which accepts the framework instance
// and returns an error.
//
// It can be used to register a custom runner with `Run` in order
// to set the framework's server listen action.
//
// Currently Runner is being used to declare the built'n server listeners.
//
// See `Run` for more.
type Runner func(*Application) error

// Listener can be used as an argument for the `Run` method.
// It can start a server with a custom net.Listener via server's `Serve`.
//
// See `Run` for more.
func Listener(l net.Listener) Runner {
	return func(app *Application) error {
		app.config.vhost = netutil.ResolveVHost(l.Addr().String())
		return app.NewHost(new(http.Server)).
			Serve(l)
	}
}

// Server can be used as an argument for the `Run` method.
// It can start a server with a *http.Server.
//
// See `Run` for more.
func Server(srv *http.Server) Runner {
	return func(app *Application) error {
		return app.NewHost(srv).
			ListenAndServe()
	}
}

// Addr can be used as an argument for the `Run` method.
// It accepts a host address which is used to build a server
// and a listener which listens on that host and port.
//
// Addr should have the form of [host]:port, i.e localhost:8080 or :8080.
//
// See `Run` for more.
func Addr(addr string) Runner {
	return func(app *Application) error {
		return app.NewHost(&http.Server{Addr: addr}).
			ListenAndServe()
	}
}

// TLS can be used as an argument for the `Run` method.
// It will start the Application's secure server.
//
// Use it like you used to use the http.ListenAndServeTLS function.
//
// Addr should have the form of [host]:port, i.e localhost:443 or :443.
// CertFile & KeyFile should be filenames with their extensions.
//
// See `Run` for more.
func TLS(addr string, certFile, keyFile string) Runner {
	return func(app *Application) error {
		return app.NewHost(&http.Server{Addr: addr}).
			ListenAndServeTLS(certFile, keyFile)
	}
}

// AutoTLS can be used as an argument for the `Run` method.
// It will start the Application's secure server using
// certifications created on the fly by the "autocert" golang/x package,
// so localhost may not be working, use it at "production" machine.
//
// Addr should have the form of [host]:port, i.e mydomain.com:443.
//
// See `Run` for more.
func AutoTLS(addr string) Runner {
	return func(app *Application) error {
		return app.NewHost(&http.Server{Addr: addr}).
			ListenAndServeAutoTLS()
	}
}

// Raw can be used as an argument for the `Run` method.
// It accepts any (listen) function that returns an error,
// this function should be block and return an error
// only when the server exited or a fatal error caused.
//
// With this option you're not limited to the servers
// that ion can run by-default.
//
// See `Run` for more.
func Raw(f func() error) Runner {
	return func(*Application) error {
		return f()
	}
}

// Build sets up, once, the framework.
// It builds the default router with its default macros
// and the template functions that are very-closed to ion.
func (app *Application) Build() error {
	rp := errors.NewReporter()

	app.once.Do(func() {
		rp.Describe("api builder: %v", app.APIBuilder.GetReport())

		if !app.Router.Downgraded() {
			// router
			// create the request handler, the default routing handler
			routerHandler := router.NewDefaultHandler()

			rp.Describe("router: %v", app.Router.BuildRouter(app.ContextPool, routerHandler, app.APIBuilder))
			// re-build of the router from outside can be done with;
			// app.RefreshRouter()
		}

		if app.view.Len() > 0 {
			// view engine
			// here is where we declare the closed-relative framework functions.
			// Each engine has their defaults, i.e yield,render,render_r,partial, params...
			rv := router.NewRoutePathReverser(app.APIBuilder)
			app.view.AddFunc("urlpath", rv.Path)
			// app.view.AddFunc("url", rv.URL)
			rp.Describe("view: %v", app.view.Load())
		}
	})

	return rp.Return()
}

// ErrServerClosed is returned by the Server's Serve, ServeTLS, ListenAndServe,
// and ListenAndServeTLS methods after a call to Shutdown or Close.
//
// Conversion for the http.ErrServerClosed.
var ErrServerClosed = http.ErrServerClosed

// Run builds the framework and starts the desired `Runner` with or without configuration edits.
//
// Run should be called only once per Application instance, it blocks like http.Server.
//
// If more than one server needed to run on the same ion instance
// then create a new host and run it manually by `go NewHost(*http.Server).Serve/ListenAndServe` etc...
// or use an already created host:
// h := NewHost(*http.Server)
// Run(Raw(h.ListenAndServe), WithoutStartupLog, WithCharset("UTF-8"))
//
// The Application can go online with any type of server or ion's host with the help of
// the following runners:
// `Listener`, `Server`, `Addr`, `TLS`, `AutoTLS` and `Raw`.
func (app *Application) Run(serve Runner, withOrWithout ...Configurator) error {
	// first Build because it doesn't need anything from configuration,
	//  this give the user the chance to modify the router inside a configurator as well.
	if err := app.Build(); err != nil {
		return errors.PrintAndReturnErrors(err, app.logger.Errorf)
	}

	app.Configure(withOrWithout...)

	// this will block until an error(unless supervisor's DeferFlow called from a Task).
	err := serve(app)
	if err != nil {
		app.Logger().Errorln(err)
	}
	return err
}
