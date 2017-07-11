[![ion](https://user-images.githubusercontent.com/29665371/27510063-3b9587da-5912-11e7-89e4-d0c53fd09bd4.png)](https://get-invite.herokuapp.com)

<p align="center">
	<a href="https://travis-ci.org/get-ion/ion">
		<img src="https://img.shields.io/travis/get-ion/ion/master.svg?style=flat-square" alt="build status">
	</a>
	<a href="http://goreportcard.com/report/get-ion/ion">
		<img src="https://img.shields.io/badge/report%20card-a%2B-ff3333.svg?style=flat-square" alt="report card">
	</a>
	<a href="https://godoc.org/github.com/get-ion/ion">
		<img src="https://img.shields.io/badge/godocs-1.1.x-0366d6.svg?style=flat-square" alt="godocs">
	</a>
	<a href="https://github.com/get-ion/issues-v1/issues">
		<img src="https://img.shields.io/badge/get-support-cccc00.svg?style=flat-square" alt="get support">
	</a>
	<a href="https://github.com/get-ion/ion/tree/master/_examples">
		<img src="https://img.shields.io/badge/learn%20by-examples-0077b3.svg?style=flat-square" alt="view examples">
	</a>
	<a href="https://get-invite.herokuapp.com">
		<img src="https://get-invite.herokuapp.com/badge.svg?style=flat-square" alt="ion channel on slack">
	</a>
	<a href="http://webchat.freenode.net?channels=get-ion">
		<img src="https://img.shields.io/badge/irc-%23get--ion%20-61DAFB.svg?style=flat-square" alt="#get-ion on freenode">
	</a>
</p>

<br/>

<!--
[![build status](https://img.shields.io/travis/get-ion/ion/master.svg?style=flat-square)](https://travis-ci.org/get-ion/ion)
[![report card](https://img.shields.io/badge/report%20card-a%2B-ff3333.svg?style=flat-square)](http://goreportcard.com/report/get-ion/ion)
[![godocs](https://img.shields.io/badge/godocs-1.1.x-0366d6.svg?style=flat-square)](https://godoc.org/github.com/get-ion/ion)
[![get support](https://img.shields.io/badge/get-support-cccc00.svg?style=flat-square)](https://github.com/get-ion/issues-v1/issues)
[![view examples](https://img.shields.io/badge/learn%20by-examples-0077b3.svg?style=flat-square)](https://github.com/get-ion/ion/tree/master/_examples)
[![ion channel on slack](https://get-invite.herokuapp.com/badge.svg?style=flat-square)](https://get-invite.herokuapp.com)
[![#get-ion on freenode](https://img.shields.io/badge/irc-%23get--ion%20-61DAFB.svg?style=flat-square)](http://webchat.freenode.net?channels=get-ion)
-->

Ion is a fast, simple and efficient micro web framework for Go. It provides a beautifully expressive and easy to use foundation for your next website, API, or distributed app.

<!--| Feature | Awesome |
| -----------|-------------|
| [Core](_examples/) | &#10003; |
| [Custom Context Registry](_examples/#routing-grouping-dynamic-path-parameters-macros-and-custom-context) | &#10003; |
| [View Engine](_examples/#view) | &#10003; |
| [Sessions](https://github.com/get-ion/sessions) | &#10003; |
| [Websockets](https://github.com/get-ion/websocket) | &#10003; |
| [Caching](https://github.com/get-ion/cache) | &#10003; |
| [Typescript Tools](https://github.com/get-ion/typescript) | &#10003; |
| [Test Framework](_examples/#testing) | &#10003; |
| [STD `net/http` compatibility](_examples/#convert-httphandlerhandlerfunc) | &#10003; |-->


### 📑 Table of contents

* [Installation](#-installation)
* [Learn](#-learn)
	* [HTTP Listening](_examples/#http-listening)
	* [Configuration](_examples/#configuration)
	* [Routing, Grouping, Dynamic Path Parameters, "Macros" and Custom Context](_examples/#routing-grouping-dynamic-path-parameters-macros-and-custom-context)
	* [Subdomains](_examples/#subdomains)
	* [Wrap `http.Handler/HandlerFunc`](_examples/#convert-httphandlerhandlerfunc)
	* [View](_examples/#view)
	* [Authentication](_examples/#authentication)
	* [File Server](_examples/#file-server)
	* [How to Read from `context.Request() *http.Request`](_examples/#how-to-read-from-contextrequest-httprequest)
	* [How to Write to `context.ResponseWriter() http.ResponseWriter`](_examples/#how-to-write-to-contextresponsewriter-httpresponsewriter)
	* [Test](_examples/#testing)	
	* [Cache](cache/#table-of-contents)
	* [Sessions](sessions/#table-of-contents)
	* [Websockets](websocket/#table-of-contents)
	* [Miscellaneous](_examples/#miscellaneous)
	* [Typescript Automation Tools](typescript/#table-of-contents)
	* [Tutorial: Online Visitors](_examples/tutorial/online-visitors)
	* [Tutorial: URL Shortener using BoltDB](_examples/tutorial/url-shortener)
* [Middleware](middleware/)
* [Dockerize](https://github.com/hiveminded/cloud-native-go)
* [Support](#-support)
* [People](#-people)

### 🚀 Installation

The only requirement is the [Go Programming Language](https://golang.org/dl/), at least version 1.8.x

```sh
$ go get github.com/get-ion/ion
```

> _ion_ takes advantage of the [vendor directory](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo) feature. You get truly reproducible builds, as this method guards against upstream renames and deletes.

```go
// file: main.go
package main
import (
    "github.com/get-ion/ion"
    "github.com/get-ion/ion/context"
)
func main() {
    app := ion.New()
    // Load all templates from the "./templates" folder
    // where extension is ".html" and parse them
    // using the standard `html/template` package.
    app.RegisterView(ion.HTML("./templates", ".html"))

    // Method:    GET
    // Resource:  http://localhost:8080
    app.Get("/", func(ctx context.Context) {
        // Bind: {{.message}} with "Hello world!"
        ctx.ViewData("message", "Hello world!")
        // Render template file: ./templates/hello.html
        ctx.View("hello.html")
    })

    // Start the server using a network address and block.
    app.Run(ion.Addr(":8080"))
}
```
```html
<!-- file: ./templates/hello.html -->
<html>
<head>
    <title>Hello Page</title>
</head>
<body>
    <h1>{{.message}}</h1>
</body>
</html>
```

```sh 
$ go run main.go
> Now listening on: http://localhost:8080
> Application started. Press CTRL+C to shut down.
```

<details>
<summary>Hello World with Go 1.9</summary>

If you've installed Go 1.9 then you can omit the `github.com/get-ion/ion/context` package from the imports statement.

```go
// +build go1.9

package main

import "github.com/get-ion/ion"

func main() {
	app := ion.New()
	app.RegisterView(ion.HTML("./templates", ".html"))
	
	app.Get("/", func(ctx ion.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("hello.html")
	})

	app.Run(ion.Addr(":8080"))
}
```

We expect Go version 1.9 to be released in August, however you can install Go 1.9 beta today.

### Installing Go 1.9beta2
 
1. Go to https://golang.org/dl/#go1.9beta2
2. Download a compatible, with your OS, archieve, i.e `go1.9beta2.windows-amd64.zip`
3. Unzip the contents of `go1.9beta2.windows-amd64.zip/go` folder to your $GOROOT, i.e `C:\Go`
4. Open a terminal and execute `go version`, it should output the go1.9beta2 version, i.e:
```sh
C:\Users\hiveminded>go version
go version go1.9beta2 windows/amd64
```

</details>

<details>
<summary>Why another new web framework?</summary>

_ion_ is easy, it has a familiar API while in the same has far more features than [Gin](https://github.com/gin-gonic/gin) or [Martini](https://github.com/go-martini/martini).

You own your code —it will never generate (unfamiliar) code for you, like [Beego](https://github.com/astaxie/beego), [Revel](https://github.com/revel/revel) and [Buffalo](https://github.com/gobuffalo/buffalo) do.

It's not just-another-router but its overall performance is equivalent with something like [httprouter](https://github.com/julienschmidt/httprouter).

Unlike [fasthttp](https://github.com/valyala/fasthttp), ion provides full HTTP/2 support for free.

Compared to the rest open source projects, this one is very active and you get answers almost immediately.

</details>

### 👥 Community

Join the welcoming community of fellow _ion_ developers in [slack](https://get-invite.herokuapp.com).

### 🏫 Learn

The awesome _ion_ community is always adding new examples, [_examples](_examples/) is a great place to get started!

Read the [godocs](https://godoc.org/github.com/get-ion/ion) for a better understanding.

### 💙 Support

- [Post](https://github.com/get-ion/issues-v1/issues) a feature request or report a bug
- :star: and watch [the public repository](https://github.com/get-ion/ion/stargazers), will keep you up to date
- :earth_americas: publish [an article](https://medium.com/search?q=ionframework) or share a [tweet](https://twitter.com/hashtag/ionframework) about your personal experience with ion

### 🥇 People

The current lead maintainer is [Bill Qeras, Jr.](https://github.com/hiveminded)

[List of all contributors](https://github.com/get-ion/ion/graphs/contributors)