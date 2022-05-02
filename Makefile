.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot
	
start:
	go run cmd/bot/main.go

get-dependencies:
	go mod tidy
	go mod vendor