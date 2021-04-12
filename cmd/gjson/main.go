package main

import (
	"syscall/js"

	"github.com/tidwall/gjson"
)

func main() {

	js.Global().Get("window").Set("gjsonGet", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			var json, path string
			if len(args) > 0 {
				json = args[0].String()
			}
			if len(args) > 1 {
				path = args[1].String()
			}
			return js.ValueOf(gjson.Get(json, path).Raw)
		},
	))

	select {}
}
