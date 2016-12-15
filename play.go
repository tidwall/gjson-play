package main

import (
	"strings"

	"github.com/gopherjs/gopherjs/js"
	"github.com/tidwall/gjson"
)

const examplePath = "name.last"

const exampleJSON = `
{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44},
    {"first": "Roger", "last": "Craig", "age": 68},
    {"first": "Jane", "last": "Murphy", "age": 47}
  ]
}
`

var path, input, output *js.Object

func main() {
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
}

func doGet() {
	res := gjson.Get(input.Get("value").String(), path.Get("value").String())
	output.Set("innerHTML", res.String())
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
