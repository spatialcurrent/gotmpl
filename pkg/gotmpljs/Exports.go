// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gotmpljs

import (
	"bytes"
	"text/template"
)

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jsbuiltin"
	"github.com/pkg/errors"
)

import (
	"github.com/spatialcurrent/gotmpl/pkg/gotmpl"
)

func isArray(object *js.Object) bool {
	return js.Global.Get("Array").Call("isArray", object).Bool()
}

func toArray(object *js.Object) []interface{} {
	arr := make([]interface{}, 0, object.Length())
	for i := 0; i < object.Length(); i++ {
		arr = append(arr, parseObject(object.Index(i)))
	}
	return arr
}

func parseObject(in *js.Object) interface{} {
	if isArray(in) {
		return toArray(in)
	}
	switch jsbuiltin.TypeOf(in) {
	case jsbuiltin.TypeObject:
		out := map[string]interface{}{}
		for _, key := range js.Keys(in) {
			out[key] = parseObject(in.Get(key))
		}
		return out
	}
	return in.Interface()
}

var Exports = map[string]interface{}{
	"render": func(tmpl string, ctx *js.Object) map[string]interface{} {
		funcs := gotmpl.InitFunctions()
		t, err := template.New("main").Funcs(funcs).Parse(tmpl)
		if err != nil {
			return map[string]interface{}{"str": nil, "err": errors.Wrap(err, "error creating template").Error()}
		}
		buf := new(bytes.Buffer)
		err = t.Execute(buf, parseObject(ctx))
		if err != nil {
			return map[string]interface{}{"str": nil, "err": errors.Wrap(err, "error rendering template").Error()}
		}
		return map[string]interface{}{"str": buf.String(), "err": nil}
	},
}
