// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gotmpl

import (
	"fmt"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/spatialcurrent/go-adaptive-functions/pkg/af"
	"github.com/spatialcurrent/go-simple-serializer/pkg/serializer"
)

// InitFunctions returns a map of functions for use with Go's templating engine.
func InitFunctions() map[string]interface{} {
	funcs := map[string]interface{}{
		"parse": func(args ...interface{}) (interface{}, error) {
			if len(args) != 2 {
				return nil, errors.New("invalid arguments for parse " + fmt.Sprint(args))
			}
			if text, ok := args[1].(string); ok {
				if format, ok := args[0].(string); ok {
					b, err := serializer.New(format).Limit(-1).Deserialize([]byte(text))
					if err != nil {
						return "", errors.Wrap(err, "error deserializing object with format "+format)
					}
					return b, nil
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
					return f.ValidateRun(args...)
				}
				return f.ValidateRun(append([]interface{}{args[len(args)-1]}, args[0:len(args)-1]...)...)
			}
		}
	}

	return funcs
}
