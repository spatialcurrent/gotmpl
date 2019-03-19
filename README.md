[![CircleCI](https://circleci.com/gh/spatialcurrent/gotmpl/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/gotmpl/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/gotmpl)](https://goreportcard.com/report/spatialcurrent/gotmpl)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/gotmpl?status.svg)](https://godoc.org/github.com/spatialcurrent/gotmpl) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/gotmpl/blob/master/LICENSE)

# gotmpl

# Description

**gotmpl** is a super simple command line program for rendering templates.  **gotmpl** uses environment variables as its context and [go-adaptive-functions](https://github.com/spatialcurrent/go-adaptive-functions) for its functions.  **gotmpl** uses [go-simple-serializer](https://github.com/spatialcurrent/go-simple-serializer) for parsing text into data, such as in the "Render Time Table" example below.

# Installation

No installation is required.  Just grab a [release](https://github.com/spatialcurrent/gotmpl/releases).  You might want to rename your binary to just `gotmpl` for convenience.

If you do have go already installed, you can just run using `go run main.go` or install with `bash scripts/install.sh`

# Usage

See the few examples below.

**Note**: Since Go templates add the piped value to the end of the positional argument array, **gotmpl** reorders the piped value to the beginning of the argument array, so **go-adaptive-functions** can be used.  This leads to a more seamless pattern, particularly for functions such as `split`, `join`, etc.

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

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/gotmpl/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
