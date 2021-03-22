#!/bin/bash

go get github.com/tidwall/gjson
go get github.com/tidwall/pretty
cp $(go env GOPATH)/src/github.com/tidwall/gjson/gjson.go .

# update some variable name stuff
sed -i '' 's/package gjson/package main/' gjson.go
sed -i '' 's/func stringBytes/func stringBytesDisabled/' gjson.go
sed -i '' 's/func bytesString/func bytesStringDisabled/' gjson.go
cat >> gjson.go << EOF

func stringBytes(s string) []byte {
	return []byte(s)
}

func bytesString(b []byte) string {
	return string(b)
}

EOF

