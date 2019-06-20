// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// gotmpl.mod.js is the package for gotmpl that is built as a JavaScript module.
// In modern JavaScript, the module can be imported using destructuring assignment.
// The functions are defined in the Exports variable in the gotmpljs package.
//
// Usage
//	// Below is a simple set of examples of how to use this package in a JavaScript application.
//
//	// load functions into current scope
//	const { render } = require('./dist/gotmpl.global.min.js);
//
//	// Render a template to a string.
//	// Returns an object, which can be destructured to the rendered string and error as a string.
//	// If there is no error, then err will be null.
//	var { str, err } = render(tmpl, context);
//
// References
//	- https://godoc.org/pkg/github.com/spatialcurrent/gotmpl/pkg/gotmpljs/
//	- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment
package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/spatialcurrent/gotmpl/pkg/gotmpljs"
)

func main() {
	jsModuleExports := js.Module.Get("exports")
	for name, value := range gotmpljs.Exports {
		jsModuleExports.Set(name, value)
	}
}
