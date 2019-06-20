// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// gotmpl.global.js is the package for gotmpl that adds the render function to the global scope under the "gotmpl" variable.
//
// In Node, depending on where require is called and the build system used, the functions may need to be required at the top of each module file.
// In a web browser, gss can be made available to the entire web page.
// The functions are defined in the Exports variable in the gotmpljs package.
//
// Usage
//	// Below is a simple set of examples of how to use this package in a JavaScript application.
//
//	// load functions into global scope
//	// require('./dist/gotmpl.global.min.js);
//
//	// Render a template to a string.
//	// Returns an object, which can be destructured to the rendered string and error as a string.
//	// If there is no error, then err will be null.
//	var { str, err } = gotmpl.render(tmpl, context);
//
// References
//	- https://godoc.org/pkg/github.com/spatialcurrent/gotmpl/pkg/gotmpljs/
//	- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment
//	- https://nodejs.org/api/globals.html#globals_global_objects
//	- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects
package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/spatialcurrent/gotmpl/pkg/gotmpljs"
)

func main() {
	js.Global.Set("gotmpl", gotmpljs.Exports)
}
