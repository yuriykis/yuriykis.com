BINARY_NAME=app
build:
	@GOOS=js GOARCH=wasm go build -o public/$(BINARY_NAME).wasm ./cmd/render

run: build
	@go run main.go

.PHONY: build