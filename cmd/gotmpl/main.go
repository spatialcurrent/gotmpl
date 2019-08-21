// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// gotmpl is a super simple command line program for rendering templates.
//
// Usage
//
// Use `gotmpl help` to see full help documentation.
//
//	gotmpl [k=v]... < template_file
//
// Examples
//
//	# convert .gitignore to JSON
//	cat .gitignore | gss -i csv --input-header path -o json
//
//	# extract version from CircleCI config
//	cat .circleci/config.yml | gss -i yaml -o json -c '#' | jq -r .version
//
//	# convert list of files to JSON Lines
//	find . -name '*.go' | gss -i csv --input-header path -o jsonl
package main

import (
	//"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

import (
	"github.com/spatialcurrent/gotmpl/pkg/gotmpl"
)

func initFlags(flag *pflag.FlagSet) {
	flag.StringP("delimiter", "d", "", "split stdin by delimiter with each element being treated as a separate template.")
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

			stdinBytes, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}

			templates := make([]string, 0)
			if delimiter, err := cmd.Flags().GetString("delimiter"); err == nil && len(delimiter) > 0 {
				templates = strings.Split(string(stdinBytes), delimiter)
			} else {
				templates = append(templates, string(stdinBytes))
			}

			ctx := gotmpl.LoadContext(args)
			funcs := gotmpl.InitFunctions()
			for _, text := range templates {
				tmpl, err := template.New("main").Funcs(funcs).Parse(text)
				if err != nil {
					return errors.Wrap(err, "error creating template")
				}
				err = tmpl.Execute(os.Stdout, ctx)
				if err != nil {
					return errors.Wrap(err, "error rendering template")
				}
			}

			return nil
		},
	}
	initFlags(cmd.Flags())

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
