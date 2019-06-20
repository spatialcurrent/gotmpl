const { render } = global.gotmpl;

function log(str) {
  console.log(str.replace(/\n/g, "\\n").replace(/\t/g, "\\t").replace(/"/g, "\\\""));
}

describe('gotmpl', () => {

  it('render single context variable', () => {
    var { str, err } = render("{{ .foo }}", {foo: "bar"});
    expect(err).toBeNull();
    expect(str).toEqual("bar");
  });

  it('split variable and join', () => {
    var { str, err } = render("{{ .foo | split \",\" | join \":\" }}", {foo: "a,b,c"});
    expect(err).toBeNull();
    expect(str).toEqual("a:b:c");
  });

  it('sum array', () => {
    var { str, err } = render("{{ .values | sum }}", {values: [1,2,3,4]});
    expect(err).toBeNull();
    expect(str).toEqual("10");
  });

  it('max variables', () => {
    var { str, err } = render("{{ .values | max }} {{ .unit }}", {values: [1,2,3,4], unit: "seconds"});
    expect(err).toBeNull();
    expect(str).toEqual("4 seconds");
  });

  it('sum, max, and values', () => {
    var { str, err } = render("Sum: {{ .values | sum }}\nMax: {{ .values | max }}\nValues: {{ .values | join \", \" }}", {values: [1,2,3,4], unit: "seconds"});
    expect(err).toBeNull();
    expect(str).toEqual("Sum: 10\nMax: 4\nValues: 1, 2, 3, 4");
  });

});
