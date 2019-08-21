[![CircleCI](https://circleci.com/gh/spatialcurrent/gotmpl/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/gotmpl/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/gotmpl)](https://goreportcard.com/report/spatialcurrent/gotmpl)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/gotmpl?status.svg)](https://godoc.org/github.com/spatialcurrent/gotmpl) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/gotmpl/blob/master/LICENSE)

# gotmpl

# Description

**gotmpl** is a lightweight wrapper around Go's native templating engine ([text/template](https://godoc.org/text/template)) that includes many more [template functions](https://godoc.org/text/template#hdr-Functions).  gotmpl is delivered as a simple command line program and supports multiple languages through cross compilers.

gotmpl uses [go-adaptive-functions](https://github.com/spatialcurrent/go-adaptive-functions) for its functions and [go-simple-serializer](https://github.com/spatialcurrent/go-simple-serializer) for serializing data, such as in the "Render Time Table" example below.

Using cross compilers, this library can also be called by other languages, including `C`, `C++`, `Python`, and `JavaScript`.  This library is cross compiled into a Shared Object file (`*.so`), which can be called by `C`, `C++`, and `Python` on Linux machines.  This library is also compiled to pure `JavaScript` using [GopherJS](https://github.com/gopherjs/gopherjs), which can be called by [Node.js](https://nodejs.org) and loaded in the browser.  See the examples folder for patterns that you can use.

# Installation

**CLI**

No setup is required to run a gotmpl executable.  Just grab a [release](https://github.com/spatialcurrent/gotmpl/releases).  You might want to rename your binary to just `gotmpl` for convenience.

If you do have go already installed, you can run using `go run github.com/spatialcurrent/gotmpl/cmd/gotmpl` or install with `make install`

# Usage

See the [text/template](https://godoc.org/text/template) documentation for information on how to use go templates.  [hugo](https://gohugo.io/) uses the same template engine, but with some changes.  Since Go's native template engine behavior is to add the piped value to the end of the positional argument array, gotmpl reorders the piped value to the beginning of the argument array, so functions from go-adaptive-functions can be used.  This leads to a more seamless pattern, particularly for functions such as `split`, `join`, etc.

**CLI**

The command line tool, `gotmpl`, can be used to render templates.  We currently support the following platforms.

| GOOS | GOARCH |
| ---- | ------ |
| darwin | amd64 |
| linux | amd64 |
| windows | amd64 |
| linux | arm64 |

Pull requests to support other platforms are welcome!  **gotmpl** parses its context from environment variables and positional arguments.  The template is read from stdin.  See the [examples](#examples) section below for usage.

**Go**

You can install the gotmpl package with.


```shell
go get -u -d github.com/spatialcurrent/gotmpl/...
```

You can then import the main public API with `import "github.com/spatialcurrent/gotmpl/pkg/gotmpl"`.  The package only contains 2 functions.  If you want to load the [go-adaptive-functions](https://github.com/spatialcurrent/go-adaptive-functions) into your own template, you can use `InitFunctions` to return a map of functions.

See [gotmpl](https://godoc.org/github.com/spatialcurrent/gotmpl/pkg/gotmpl) in GoDoc for API documentation and examples.

**Node**

gotmpl is built as a module.  In modern JavaScript, the module can be imported using [destructuring assignment](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment).

```javascript
const { render } = require('./dist/gotmpl.global.min.js');
```

In legacy JavaScript, you can use the `gotmpl.main.js` file that simply adds `gotmpl` to the global scope.

# Examples

**Get Shell**

```shell
echo '{{ split .SHELL "/" | last }}' | gotmpl
```

**Clean Path**

```shell
echo '{{ .PATH | split ":" | set | array | sort | join ":" }}' | gotmpl
```

**Render Time Table**

```shell
echo '{{ with $items := .data | parse "csv" }}<table style="text-align:left;font-size:16px;"><tr><th>Time</th><th>Title</th>{{ range $items }}<tr><td>{{ .Time }}</td><td>{{ .Title }}</td></tr>{{ end }}</table>{{ end }}' | data=$(cat timetable.csv) gotmpl
```

# Testing

**Go**

To run Go tests use `make test_go` (or `bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [ineffassign](https://github.com/gordonklaus/ineffassign), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

**JavaScript**

To run JavaScript tests, first install [Jest](https://jestjs.io/) using `make deps_javascript`, use [Yarn](https://yarnpkg.com/en/), or another method.  Then, build the JavaScript module with `make build_javascript`.  To run tests, use `make test_javascript`.  You can also use the scripts in the `package.json`.

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/gotmpl/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
