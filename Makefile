
local-run:
	go run cmd/main.go -D -C config.json

build-server:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/chatgpt ./cmd/main.go
