
wasm-build:
	GOOS=js GOARCH=wasm go build -o  ./json.wasm

wasm-exec:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .