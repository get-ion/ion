[![ion](https://user-images.githubusercontent.com/29665371/27510063-3b9587da-5912-11e7-89e4-d0c53fd09bd4.png)](https://get-invite.herokuapp.com)

<p align="center">
	<a href="https://travis-ci.org/get-ion/ion">
		<img src="https://img.shields.io/travis/get-ion/ion/master.svg?style=flat-square" alt="build status">
	</a>
	<a href="http://goreportcard.com/report/get-ion/ion">
		<img src="https://img.shields.io/badge/report%20card-a%2B-ff3333.svg?style=flat-square" alt="report card">
	</a>
	<a href="https://godoc.org/github.com/get-ion/ion">
		<img src="https://img.shields.io/badge/godocs-1.0.x-0366d6.svg?style=flat-square" alt="godocs">
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
[![godocs](https://img.shields.io/badge/godocs-1.0.x-0366d6.svg?style=flat-square)](https://godoc.org/github.com/get-ion/ion)
[![get support](https://img.shields.io/badge/get-support-cccc00.svg?style=flat-square)](https://github.com/get-ion/issues-v1/issues)
[![view examples](https://img.shields.io/badge/learn%20by-examples-0077b3.svg?style=flat-square)](https://github.com/get-ion/ion/tree/master/_examples)
[![ion channel on slack](https://get-invite.herokuapp.com/badge.svg?style=flat-square)](https://get-invite.herokuapp.com)
[![#get-ion on freenode](https://img.shields.io/badge/irc-%23get--ion%20-61DAFB.svg?style=flat-square)](http://webchat.freenode.net?channels=get-ion)
-->

Ion is a fast, simple and efficient micro web framework for Go. It provides a beautifully expressive and easy to use foundation for your next website, API, or distributed app.

| Feature | Awesome |
| -----------|-------------|
| [Core](_examples/) | &#10003; |
| [Custom Context Registry](_examples/#basic-routing-grouping-dynamic-path-parameters-macros-and-custom-context) | &#10003; |
| [View Engine](_examples/#view) | &#10003; |
| [Sessions](https://github.com/get-ion/sessions) | &#10003; |
| [Websockets](https://github.com/get-ion/websocket) | &#10003; |
| [Caching](https://github.com/get-ion/cache) | &#10003; |
| [Typescript Tools](https://github.com/get-ion/typescript) | &#10003; |
| [Test Framework](_examples/#testing) | &#10003; |
| [STD `net/http` compatibility](_examples/#convert-httphandlerhandlerfunc) | &#10003; |

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
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.HTML("<b>Hello world!</b>")
	})
	app.Run(ion.Addr(":8080"))
}
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
	app.Handle("GET", "/", func(ctx ion.Context) {
		ctx.HTML("<b>Hello world!</b>")
	})
	app.Run(ion.Addr(":8080"))
}
```

We expect Go version 1.9 to be released in August, however you can install Go 1.9 beta today.

### [Installing Go 1.9 from source](https://golang.org/doc/install/source)
 
(Optional) Install a C compiler

To build a Go installation with cgo support, which permits Go programs to import C libraries, a C compiler such as gcc or clang must be installed first. Do this using whatever installation method is standard on the system.

To build without cgo, set the environment variable CGO_ENABLED=0 before running all.bash or make.bash. 


Rename your current Go root folder (i.e C:/go) to "go_1.8.3"

Set `GOROOT_BOOTSTRAP` to that go root path (i.e C:/go_1.8.3)

Execute the below commands

```sh
git clone https://go.googlesource.com/go
cd go
git checkout go1.9beta1
cd src
./all.bash
```

> To build under Unix use all.bash



If all goes well, it will finish by printing output like:
```
ALL TESTS PASSED

---
Installed Go for linux/amd64 in /home/you/go.
Installed commands in /home/you/go/bin.
*** You need to add /home/you/go/bin to your $PATH. ***
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

The awesome _ion_ community is always adding new examples, [_examples](/examples) is a great place to get started!

Read the [godocs](https://godoc.org/github.com/get-ion/ion) for a better understanding.

### 💙 Support

- [Post](https://github.com/get-ion/issues-v1/issues) a feature request or report a bug
- :star: and watch [the public repository](https://github.com/get-ion/ion/stargazers), will keep you up to date
- :earth_americas: publish [an article](https://medium.com/search?q=ionframework) or share a [tweet](https://twitter.com/hashtag/ionframework) about your personal experience with ion

### 🥇 People

The current lead maintainer is [Bill Qeras, Jr.](https://github.com/hiveminded)

[List of all contributors](https://github.com/get-ion/ion/graphs/contributors)