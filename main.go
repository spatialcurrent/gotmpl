// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

import (
	"github.com/spatialcurrent/go-adaptive-functions/af"
)

func main() {

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 && !fi.Mode().IsRegular() {
		fmt.Println("Usage: TEMPLATE_TEXT | gotmpl")
		fmt.Println("Usage: gotmpl < TEMPLATE_FILE")
		return
	}

	text, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	ctx := map[string]string{}
	for _, str := range os.Environ() {
		parts := strings.SplitN(str, "=", 2)
		ctx[parts[0]] = parts[1]
	}

	funcs := map[string]interface{}{}
	for _, f := range af.Functions {
		f := f
		for _, alias := range f.Aliases {
			alias := alias
			funcs[alias] = func(args ...interface{}) (interface{}, error) {
				if len(args) <= 1 {
					return f.ValidateRun(args)
				}
				return f.ValidateRun(append([]interface{}{args[len(args)-1]}, args[0:len(args)-1]...))
			}
		}
	}

	tmpl, err := template.New("main").Funcs(funcs).Parse(string(text))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, ctx)
	if err != nil {
		panic(err)
	}
}
