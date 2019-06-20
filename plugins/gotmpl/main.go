// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// gss.so creates a shared library of Go that can be called by C, C++, or Python
//
//
//  - https://godoc.org/cmd/cgo
//  - https://blog.golang.org/c-go-cgo
//
package main

import (
	"C"
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/spatialcurrent/go-simple-serializer/pkg/serializer"
	"github.com/spatialcurrent/gotmpl/pkg/gotmpl"
)

var gitBranch string
var gitCommit string

func main() {}

//export Version
func Version() *C.char {
	var b strings.Builder
	if len(gitBranch) > 0 {
		b.WriteString(fmt.Sprintf("Branch: %q\n", gitBranch))
	}
	if len(gitCommit) > 0 {
		b.WriteString(fmt.Sprintf("Commit: %q\n", gitCommit))
	}
	return C.CString(b.String())
}

//export Render
func Render(tmpl *C.char, ctx *C.char, format *C.char, outputString **C.char) *C.char {
	c, err := serializer.New(C.GoString(format)).Limit(-1).Deserialize([]byte(C.GoString(ctx)))
	if err != nil {
		return C.CString(errors.Wrapf(err, "error deserializing context with format %q", C.GoString(format)).Error())
	}
	funcs := gotmpl.InitFunctions()
	t, err := template.New("main").Funcs(funcs).Parse(C.GoString(tmpl))
	if err != nil {
		return C.CString(errors.Wrapf(err, "error creating template %q", C.GoString(tmpl)).Error())
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, c)
	if err != nil {
		return C.CString(errors.Wrap(err, "error rendering template").Error())
	}
	*outputString = C.CString(buf.String())
	return nil
}
