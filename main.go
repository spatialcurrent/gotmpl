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
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

import (
	"github.com/spatialcurrent/go-adaptive-functions/af"
	"github.com/spatialcurrent/go-simple-serializer/gss"
)

func initFunctions() map[string]interface{} {
	funcs := map[string]interface{}{
		"parse": func(args ...interface{}) (interface{}, error) {
			if len(args) != 2 {
				return nil, errors.New("invalid arguments for parse " + fmt.Sprint(args))
			}
			if text, ok := args[1].(string); ok {
				if f, ok := args[0].(string); ok {
					t, err := gss.GetType([]byte(text), f)
					if err != nil {
						return "", errors.Wrap(err, "error creating new object for format "+f)
					}
					return gss.DeserializeString(
						text,
						f,
						gss.NoHeader,
						gss.NoComment,
						true,
						0,
						gss.NoLimit,
						t,
						false,
						false)
				}
			}
			return nil, errors.New("invalid arguments for parse " + fmt.Sprint(args))
		},
	}

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

	return funcs
}

func main() {
	cmd := &cobra.Command{
		Use:                   "gotmpl [k=v]... < template_file",
		DisableFlagsInUseLine: true,
		Short:                 "gotmpl",
		Long:                  `gotmpl is a super simple command line program for rendering templates that uses environment variables and command line arguments as its context variables.  The template is read from stdin.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			fi, err := os.Stdin.Stat()
			if err != nil {
				return err
			}

			if fi.Mode()&os.ModeNamedPipe == 0 && !fi.Mode().IsRegular() {
				return cmd.Usage()
			}

			text, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}

			ctx := map[string]string{}

			// load context from environment variables
			for _, str := range os.Environ() {
				parts := strings.SplitN(str, "=", 2)
				ctx[parts[0]] = parts[1]
			}

			// load context from command line arguments
			for _, str := range args {
				parts := strings.SplitN(str, "=", 2)
				ctx[parts[0]] = parts[1]
			}

			tmpl, err := template.New("main").Funcs(initFunctions()).Parse(string(text))
			if err != nil {
				return err
			}

			err = tmpl.Execute(os.Stdout, ctx)
			if err != nil {
				return err
			}

			return nil
		},
	}

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
