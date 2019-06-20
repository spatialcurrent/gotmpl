const { render } = require('./../../dist/gotmpl.mod.min.js');

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
