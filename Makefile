run:
	go run cmd/server/server.go

wasm:
	GOOS=js GOARCH=wasm go build -o assets/json.wasm cmd/wasm/main.go
