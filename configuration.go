package ion

import (
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"

	"github.com/get-ion/ion/context"
	"github.com/get-ion/ion/core/errors"
)

var errConfigurationDecode = errors.New("error while trying to decode configuration")

// YAML reads Configuration from a configuration.yml file.
//
// Accepts the absolute path of the configuration.yml.
// An error will be shown to the user via panic with the error message.
// Error may occur when the configuration.yml doesn't exists or is not formatted correctly.
//
// Usage:
// app := ion.Run(ion.Addr(":8080"), ion.WithConfiguration(ion.YAML("myconfig.yml")))
func YAML(filename string) Configuration {
	c := DefaultConfiguration()

	// get the abs
	// which will try to find the 'filename' from current workind dir too.
	yamlAbsPath, err := filepath.Abs(filename)
	if err != nil {
		panic(errConfigurationDecode.AppendErr(err))
	}

	// read the raw contents of the file
	data, err := ioutil.ReadFile(yamlAbsPath)
	if err != nil {
		panic(errConfigurationDecode.AppendErr(err))
	}

	// put the file's contents as yaml to the default configuration(c)
	if err := yaml.Unmarshal(data, &c); err != nil {
		panic(errConfigurationDecode.AppendErr(err))
	}

	return c
}

// TOML reads Configuration from a toml-compatible document file.
// Read more about toml's implementation at:
// https://github.com/toml-lang/toml
//
//
// Accepts the absolute path of the configuration file.
// An error will be shown to the user via panic with the error message.
// Error may occur when the file doesn't exists or is not formatted correctly.
//
// Usage:
// app := ion.Run(ion.Addr(":8080"), ion.WithConfiguration(ion.YAML("myconfig.tml")))
func TOML(filename string) Configuration {
	c := DefaultConfiguration()

	// get the abs
	// which will try to find the 'filename' from current workind dir too.
	tomlAbsPath, err := filepath.Abs(filename)
	if err != nil {
		panic(errConfigurationDecode.AppendErr(err))
	}

	// read the raw contents of the file
	data, err := ioutil.ReadFile(tomlAbsPath)
	if err != nil {
		panic(errConfigurationDecode.AppendErr(err))
	}

	// put the file's contents as toml to the default configuration(c)
	if _, err := toml.Decode(string(data), &c); err != nil {
		panic(errConfigurationDecode.AppendErr(err))
	}
	// Author's notes:
	// The toml's 'usual thing' for key naming is: the_config_key instead of TheConfigKey
	// but I am always prefer to use the specific programming language's syntax
	// and the original configuration name fields for external configuration files
	// so we do 'toml: "TheConfigKeySameAsTheConfigField" instead.
	return c
}

// Configurator is just an interface which accepts the framework instance.
//
// It can be used to register a custom configuration with `Configure` in order
// to modify the framework instance.
//
// Currently Configurator is being used to describe the configuration's fields values.
type Configurator func(*Application)

// variables for configurators don't need any receivers, functions
// for them that need (helps code editors to recognise as variables without parenthesis completion).

// WithoutStartupLog turns off the information send, once, to the terminal when the main server is open.
var WithoutStartupLog = func(app *Application) {
	app.config.DisableStartupLog = true
}

// WithoutBanner is a conversion for the `WithoutStartupLog` option.
//
// Turns off the information send, once, to the terminal when the main server is open.
var WithoutBanner = WithoutStartupLog

// WithoutInterruptHandler disables the automatic graceful server shutdown
// when control/cmd+C pressed.
var WithoutInterruptHandler = func(app *Application) {
	app.config.DisableInterruptHandler = true
}

// WithoutVersionCheck will disable the version checker and updater.
var WithoutVersionCheck = func(app *Application) {
	app.config.DisableVersionCheck = true
}

// WithoutPathCorrection disables the PathCorrection setting.
//
// See` Configuration`.
var WithoutPathCorrection = func(app *Application) {
	app.config.DisablePathCorrection = true
}

// WithoutBodyConsumptionOnUnmarshal disables BodyConsumptionOnUnmarshal setting.
//
// See` Configuration`.
var WithoutBodyConsumptionOnUnmarshal = func(app *Application) {
	app.config.DisableBodyConsumptionOnUnmarshal = true
}

// WithoutAutoFireStatusCode disables the AutoFireStatusCode setting.
//
// See` Configuration`.
var WithoutAutoFireStatusCode = func(app *Application) {
	app.config.DisableAutoFireStatusCode = true
}

// WithPathEscape enanbles the PathEscape setting.
//
// See` Configuration`.
var WithPathEscape = func(app *Application) {
	app.config.EnablePathEscape = true
}

// WithFireMethodNotAllowed enanbles the FireMethodNotAllowed setting.
//
// See` Configuration`.
var WithFireMethodNotAllowed = func(app *Application) {
	app.config.FireMethodNotAllowed = true
}

// WithTimeFormat sets the TimeFormat setting.
//
// See` Configuration`.
func WithTimeFormat(timeformat string) Configurator {
	return func(app *Application) {
		app.config.TimeFormat = timeformat
	}
}

// WithCharset sets the Charset setting.
//
// See` Configuration`.
func WithCharset(charset string) Configurator {
	return func(app *Application) {
		app.config.Charset = charset
	}
}

// WithRemoteAddrHeader enables or adds a new or existing request header name
// that can be used to validate the client's real IP.
//
// Existing values are:
// "X-Real-Ip":             false,
// "X-Forwarded-For":       false,
// "CF-Connecting-IP": false
//
// Look `context.RemoteAddr()` for more.
func WithRemoteAddrHeader(headerName string) Configurator {
	return func(app *Application) {
		if app.config.RemoteAddrHeaders == nil {
			app.config.RemoteAddrHeaders = make(map[string]bool, 0)
		}
		app.config.RemoteAddrHeaders[headerName] = true
	}
}

// WithoutRemoteAddrHeader disables an existing request header name
// that can be used to validate the client's real IP.
//
// Existing values are:
// "X-Real-Ip":             false,
// "X-Forwarded-For":       false,
// "CF-Connecting-IP": false
//
// Look `context.RemoteAddr()` for more.
func WithoutRemoteAddrHeader(headerName string) Configurator {
	return func(app *Application) {
		if app.config.RemoteAddrHeaders == nil {
			app.config.RemoteAddrHeaders = make(map[string]bool, 0)
		}
		app.config.RemoteAddrHeaders[headerName] = false
	}
}

// WithOtherValue adds a value based on a key to the Other setting.
//
// See` Configuration`.
func WithOtherValue(key string, val interface{}) Configurator {
	return func(app *Application) {
		if app.config.Other == nil {
			app.config.Other = make(map[string]interface{}, 0)
		}
		app.config.Other[key] = val
	}
}

// Configuration the whole configuration for an ion instance
// these can be passed via options also, look at the top of this file(configuration.go).
// Configuration is a valid OptionSetter.
type Configuration struct {
	// vhost is private and setted only with .Run method, it cannot be changed after the first set.
	// It can be retrieved by the context if needed (i.e router for subdomains)
	vhost string

	// DisableStartupLog if setted to true then it turns off the write banner on server startup.
	//
	// Defaults to false.
	DisableStartupLog bool `yaml:"DisableStartupLog" toml:"DisableStartupLog"`
	// DisableInterruptHandler if setted to true then it disables the automatic graceful server shutdown
	// when control/cmd+C pressed.
	// Turn this to true if you're planning to handle this by your own via a custom host.Task.
	//
	// Defaults to false.
	DisableInterruptHandler bool `yaml:"DisableInterruptHandler" toml:"DisableInterruptHandler"`

	// DisableVersionCheck if true then process will be not be notified for any available updates.
	//
	// Defaults to false.
	DisableVersionCheck bool `yaml:"DisableVersionCheck" toml:"DisableVersionCheck"`

	// DisablePathCorrection corrects and redirects the requested path to the registered path
	// for example, if /home/ path is requested but no handler for this Route found,
	// then the Router checks if /home handler exists, if yes,
	// (permant)redirects the client to the correct path /home
	//
	// Defaults to false.
	DisablePathCorrection bool `yaml:"DisablePathCorrection" toml:"DisablePathCorrection"`

	// EnablePathEscape when is true then its escapes the path, the named parameters (if any).
	// Change to false it if you want something like this https://github.com/get-ion/ion/issues/135 to work
	//
	// When do you need to Disable(false) it:
	// accepts parameters with slash '/'
	// Request: http://localhost:8080/details/Project%2FDelta
	// ctx.Param("project") returns the raw named parameter: Project%2FDelta
	// which you can escape it manually with net/url:
	// projectName, _ := url.QueryUnescape(c.Param("project").
	//
	// Defaults to false.
	EnablePathEscape bool `yaml:"EnablePathEscape" toml:"EnablePathEscape"`

	// FireMethodNotAllowed if it's true router checks for StatusMethodNotAllowed(405) and
	//  fires the 405 error instead of 404
	// Defaults to false.
	FireMethodNotAllowed bool `yaml:"FireMethodNotAllowed" toml:"FireMethodNotAllowed"`

	// DisableBodyConsumptionOnUnmarshal manages the reading behavior of the context's body readers/binders.
	// If setted to true then it
	// disables the body consumption by the `context.UnmarshalBody/ReadJSON/ReadXML`.
	//
	// By-default io.ReadAll` is used to read the body from the `context.Request.Body which is an `io.ReadCloser`,
	// if this field setted to true then a new buffer will be created to read from and the request body.
	// The body will not be changed and existing data before the
	// context.UnmarshalBody/ReadJSON/ReadXML will be not consumed.
	DisableBodyConsumptionOnUnmarshal bool `yaml:"DisableBodyConsumptionOnUnmarshal" toml:"DisableBodyConsumptionOnUnmarshal"`

	// DisableAutoFireStatusCode if true then it turns off the http error status code handler automatic execution
	// from "context.StatusCode(>=400)" and instead app should manually call the "context.FireStatusCode(>=400)".
	//
	// By-default a custom http error handler will be fired when "context.StatusCode(code)" called,
	// code should be >=400 in order to be received as an "http error handler".
	//
	// Developer may want this option to setted as true in order to manually call the
	// error handlers when needed via "context.FireStatusCode(>=400)".
	// HTTP Custom error handlers are being registered via "app.OnErrorCode(code, handler)".
	//
	// Defaults to false.
	DisableAutoFireStatusCode bool `yaml:"DisableAutoFireStatusCode" toml:"DisableAutoFireStatusCode"`

	// TimeFormat time format for any kind of datetime parsing
	// Defaults to  "Mon, 02 Jan 2006 15:04:05 GMT".
	TimeFormat string `yaml:"TimeFormat" toml:"TimeFormat"`

	// Charset character encoding for various rendering
	// used for templates and the rest of the responses
	// Defaults to "UTF-8".
	Charset string `yaml:"Charset" toml:"Charset"`

	//  +----------------------------------------------------+
	//  | Context's keys for values used on various featuers |
	//  +----------------------------------------------------+

	// Context values' keys for various features.
	//
	// TranslateLanguageContextKey & TranslateFunctionContextKey are used by i18n handlers/middleware
	// currently we have only one: https://github.com/get-ion/ion/tree/master/middleware/i18n.
	//
	// Defaults to "ion.translate" and "ion.language"
	TranslateFunctionContextKey string `yaml:"TranslateFunctionContextKey" toml:"TranslateFunctionContextKey"`
	// TranslateLanguageContextKey used for i18n.
	//
	// Defaults to "ion.language"
	TranslateLanguageContextKey string `yaml:"TranslateLanguageContextKey" toml:"TranslateLanguageContextKey"`

	// GetViewLayoutContextKey is the key of the context's user values' key
	// which is being used to set the template
	// layout from a middleware or the main handler.
	// Overrides the parent's or the configuration's.
	//
	// Defaults to "ion.ViewLayout"
	ViewLayoutContextKey string `yaml:"ViewLayoutContextKey" toml:"ViewLayoutContextKey"`
	// GetViewDataContextKey is the key of the context's user values' key
	// which is being used to set the template
	// binding data from a middleware or the main handler.
	//
	// Defaults to "ion.viewData"
	ViewDataContextKey string `yaml:"ViewDataContextKey" toml:"ViewDataContextKey"`
	// RemoteAddrHeaders returns the allowed request headers names
	// that can be valid to parse the client's IP based on.
	//
	// Defaults to:
	// "X-Real-Ip":             false,
	// "X-Forwarded-For":       false,
	// "CF-Connecting-IP": false
	//
	// Look `context.RemoteAddr()` for more.
	RemoteAddrHeaders map[string]bool `yaml:"RemoteAddrHeaders" toml:"RemoteAddrHeaders"`
	// Other are the custom, dynamic options, can be empty.
	// This field used only by you to set any app's options you want
	// or by custom adaptors, it's a way to simple communicate between your adaptors (if any)
	// Defaults to a non-nil empty map
	Other map[string]interface{} `yaml:"Other" toml:"Other"`
}

var _ context.ConfigurationReadOnly = &Configuration{}

// GetVHost returns the non-exported vhost config field.
//
// If original addr ended with :443 or :80, it will return the host without the port.
// If original addr was :https or :http, it will return localhost.
// If original addr was 0.0.0.0, it will return localhost.
func (c Configuration) GetVHost() string {
	return c.vhost
}

// GetDisablePathCorrection returns the configuration.DisablePathCorrection,
// DisablePathCorrection corrects and redirects the requested path to the registered path
// for example, if /home/ path is requested but no handler for this Route found,
// then the Router checks if /home handler exists, if yes,
// (permant)redirects the client to the correct path /home.
func (c Configuration) GetDisablePathCorrection() bool {
	return c.DisablePathCorrection
}

// GetEnablePathEscape is the configuration.EnablePathEscape,
// returns true when its escapes the path, the named parameters (if any).
func (c Configuration) GetEnablePathEscape() bool {
	return c.EnablePathEscape
}

// GetFireMethodNotAllowed returns the configuration.FireMethodNotAllowed.
func (c Configuration) GetFireMethodNotAllowed() bool {
	return c.FireMethodNotAllowed
}

// GetDisableBodyConsumptionOnUnmarshal returns the configuration.GetDisableBodyConsumptionOnUnmarshal,
// manages the reading behavior of the context's body readers/binders.
// If returns true then the body consumption by the `context.UnmarshalBody/ReadJSON/ReadXML`
// is disabled.
//
// By-default io.ReadAll` is used to read the body from the `context.Request.Body which is an `io.ReadCloser`,
// if this field setted to true then a new buffer will be created to read from and the request body.
// The body will not be changed and existing data before the
// context.UnmarshalBody/ReadJSON/ReadXML will be not consumed.
func (c Configuration) GetDisableBodyConsumptionOnUnmarshal() bool {
	return c.DisableBodyConsumptionOnUnmarshal
}

// GetDisableAutoFireStatusCode returns the configuration.DisableAutoFireStatusCode.
// Returns true when the http error status code handler automatic execution turned off.
func (c Configuration) GetDisableAutoFireStatusCode() bool {
	return c.DisableAutoFireStatusCode
}

// GetTimeFormat returns the configuration.TimeFormat,
// format for any kind of datetime parsing.
func (c Configuration) GetTimeFormat() string {
	return c.TimeFormat
}

// GetCharset returns the configuration.Charset,
// the character encoding for various rendering
// used for templates and the rest of the responses.
func (c Configuration) GetCharset() string {
	return c.Charset
}

// GetTranslateFunctionContextKey returns the configuration's TranslateFunctionContextKey value,
// used for i18n.
func (c Configuration) GetTranslateFunctionContextKey() string {
	return c.TranslateFunctionContextKey
}

// GetTranslateLanguageContextKey returns the configuration's TranslateLanguageContextKey value,
// used for i18n.
func (c Configuration) GetTranslateLanguageContextKey() string {
	return c.TranslateLanguageContextKey
}

// GetViewLayoutContextKey returns the key of the context's user values' key
// which is being used to set the template
// layout from a middleware or the main handler.
// Overrides the parent's or the configuration's.
func (c Configuration) GetViewLayoutContextKey() string {
	return c.ViewLayoutContextKey
}

// GetViewDataContextKey returns the key of the context's user values' key
// which is being used to set the template
// binding data from a middleware or the main handler.
func (c Configuration) GetViewDataContextKey() string {
	return c.ViewDataContextKey
}

// GetRemoteAddrHeaders returns the allowed request headers names
// that can be valid to parse the client's IP based on.
//
// Defaults to:
// "X-Real-Ip":             false,
// "X-Forwarded-For":       false,
// "CF-Connecting-IP": false
//
// Look `context.RemoteAddr()` for more.
func (c Configuration) GetRemoteAddrHeaders() map[string]bool {
	return c.RemoteAddrHeaders
}

// GetOther returns the configuration.Other map.
func (c Configuration) GetOther() map[string]interface{} {
	return c.Other
}

// WithConfiguration sets the "c" values to the framework's configurations.
//
// Usage:
// app.Run(ion.Addr(":8080"), ion.WithConfiguration(ion.Configuration{/* fields here */ }))
// or
// ion.WithConfiguration(ion.YAML("./cfg/ion.yml"))
// or
// ion.WithConfiguration(ion.TOML("./cfg/ion.tml"))
func WithConfiguration(c Configuration) Configurator {
	return func(app *Application) {
		main := app.config

		if v := c.DisableStartupLog; v {
			main.DisableStartupLog = v
		}

		if v := c.DisableInterruptHandler; v {
			main.DisableInterruptHandler = v
		}

		if v := c.DisablePathCorrection; v {
			main.DisablePathCorrection = v
		}

		if v := c.EnablePathEscape; v {
			main.EnablePathEscape = v
		}

		if v := c.FireMethodNotAllowed; v {
			main.FireMethodNotAllowed = v
		}

		if v := c.DisableBodyConsumptionOnUnmarshal; v {
			main.DisableBodyConsumptionOnUnmarshal = v
		}

		if v := c.DisableAutoFireStatusCode; v {
			main.DisableAutoFireStatusCode = v
		}

		if v := c.TimeFormat; v != "" {
			main.TimeFormat = v
		}

		if v := c.Charset; v != "" {
			main.Charset = v
		}

		if v := c.TranslateFunctionContextKey; v != "" {
			main.TranslateFunctionContextKey = v
		}

		if v := c.TranslateLanguageContextKey; v != "" {
			main.TranslateLanguageContextKey = v
		}

		if v := c.ViewLayoutContextKey; v != "" {
			main.ViewLayoutContextKey = v
		}

		if v := c.ViewDataContextKey; v != "" {
			main.ViewDataContextKey = v
		}

		if v := c.RemoteAddrHeaders; len(v) > 0 {
			if main.RemoteAddrHeaders == nil {
				main.RemoteAddrHeaders = make(map[string]bool, 0)
			}
			for key, value := range v {
				main.RemoteAddrHeaders[key] = value
			}
		}

		if v := c.Other; len(v) > 0 {
			if main.Other == nil {
				main.Other = make(map[string]interface{}, 0)
			}
			for key, value := range v {
				main.Other[key] = value
			}
		}
	}
}

// DefaultConfiguration returns the default configuration for an ion station, fills the main Configuration
func DefaultConfiguration() Configuration {
	return Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisableVersionCheck:               false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, Jan 02 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
		TranslateFunctionContextKey:       "ion.translate",
		TranslateLanguageContextKey:       "ion.language",
		ViewLayoutContextKey:              "ion.viewLayout",
		ViewDataContextKey:                "ion.viewData",
		RemoteAddrHeaders: map[string]bool{
			"X-Real-Ip":        false,
			"X-Forwarded-For":  false,
			"CF-Connecting-IP": false,
		},
		Other: make(map[string]interface{}, 0),
	}
}
