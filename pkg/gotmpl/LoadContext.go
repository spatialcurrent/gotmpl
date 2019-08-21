// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gotmpl

import (
	"os"
	"strings"
)

// LoadContext returns a context for use with Go's templating engine.
// LoadContext parses the context from environment variables and positional arguments.
func LoadContext(args []string) map[string]string {
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

	return ctx
}
