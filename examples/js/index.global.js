// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

require('./../../dist/gotmpl.global.min.js');
const { render } = global.gotmpl;

const tmpl = "Sum: {{ .values | sum }}\nMax: {{ .values | max }}\nValues: {{ .values | join \", \" }}"
var ctx = {
  values: [1, 2, 3, 4]
}

console.log('Template:');
console.log(tmpl);
console.log();
// Destructure return value
var { str, err } = render(tmpl, ctx);
console.log('Output:');
console.log(str);
console.log();
console.log('Error:');
console.log(err);
console.log();
console.log("************************************");
console.log();
