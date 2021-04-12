wasm:
	GOOS=js GOARCH=wasm go build -o gjson.wasm cmd/gjson/main.go
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .

serve: wasm
	go run cmd/server/main.go
	