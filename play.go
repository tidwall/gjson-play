package main

import (
	"strings"

	"github.com/gopherjs/gopherjs/js"
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
)

const examplePath = "name.last"

const exampleJSON = `
{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}
`

var path, input, output *js.Object

var style = pretty.Style{
	False:  [2]string{`<span class="json-style false">`, `</span>`},
	True:   [2]string{`<span class="json-style true">`, `</span>`},
	String: [2]string{`<span class="json-style string">`, `</span>`},
	Number: [2]string{`<span class="json-style number">`, `</span>`},
	Null:   [2]string{`<span class="json-style null">`, `</span>`},
	Key:    [2]string{`<span class="json-style key">`, `</span>`},
}

func main() {
	js.Global.Set("jsonColor", func(json string) string {
		return string(pretty.Color([]byte(json), &style))
	})
	killTabs()
	path = js.Global.Get("document").Call("getElementById", "path")
	input = js.Global.Get("document").Call("getElementById", "json-input")
	output = js.Global.Get("document").Call("getElementById", "result-output")

	path.Set("value", strings.TrimSpace(examplePath))
	input.Set("value", strings.TrimSpace(exampleJSON))
	doGet()
	path.Set("onkeyup", func() {
		doGet()
	})
	input.Set("onkeyup", func() {
		doGet()
	})
}

func doGet() {
	res := gjson.Get(input.Get("value").String(), path.Get("value").String())
	output.Set("innerHTML", string(pretty.Color([]byte(res.Raw), &style)))
}

func killTabs() {
	textareas := js.Global.Get("document").Call("getElementsByTagName", "textarea")
	count := textareas.Get("length").Int()
	for i := 0; i < count; i++ {
		func(this *js.Object) {
			this.Set("onkeydown", func(e *js.Object) {
				if e.Get("keyCode").Int() == 9 || e.Get("which").Int() == 9 {
					e.Call("preventDefault")
					s := this.Get("selectionStart").Int()
					this.Set("value",
						this.Get("value").Call("substring", 0, this.Get("selectionStart")).String()+
							"  "+
							this.Get("value").Call("substring", this.Get("selectionEnd")).String())
					this.Set("selectionEnd", s+2)
				}
			})
		}(textareas.Index(i))
	}
}
