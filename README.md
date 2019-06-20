[![CircleCI](https://circleci.com/gh/spatialcurrent/gotmpl/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/gotmpl/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/gotmpl)](https://goreportcard.com/report/spatialcurrent/gotmpl)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/gotmpl?status.svg)](https://godoc.org/github.com/spatialcurrent/gotmpl) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/gotmpl/blob/master/LICENSE)

# gotmpl

# Description

**gotmpl** is a super simple command line program for rendering template.  **gotmpl** parses its context from environment variables and positional arguments.  **gotmpl** uses [go-adaptive-functions](https://github.com/spatialcurrent/go-adaptive-functions) for its functions and [go-simple-serializer](https://github.com/spatialcurrent/go-simple-serializer) for serializing data, such as in the "Render Time Table" example below.

# Installation

No installation is required.  Just grab a [release](https://github.com/spatialcurrent/gotmpl/releases).  You might want to rename your binary to just `gotmpl` for convenience.

If you do have go already installed, you can just run using `go run cmd/gotmpl/main.go` or install with `make install`

# Usage

**CLI**

The command line tool, `gotmpl`, can be used to render templates.  We currently support the following platforms.

| GOOS | GOARCH |
| ---- | ------ |
| darwin | amd64 |
| linux | amd64 |
| windows | amd64 |
| linux | arm64 |

Pull requests to support other platforms are welcome!  See the [examples](#examples) section below for usage.

**Note**: Since Go's native template engine behavior is to add the piped value to the end of the positional argument array, **gotmpl** reorders the piped value to the beginning of the argument array, so **go-adaptive-functions** can be used.  This leads to a more seamless pattern, particularly for functions such as `split`, `join`, etc.

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

Run tests with `make test` (or `bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [ineffassign](https://github.com/gordonklaus/ineffassign), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/gotmpl/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
