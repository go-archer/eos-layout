.PHONY: init
init:
	go install github.com/google/wire/cmd/wire@latest

.PHONY: wire
wire:
	cd ./cmd/server && wire

.PHONY: run
run:
	cd ./cmd/server && go run . -c ./config.toml

.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server

.PHONY: linux
linux:
	env CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/server ./cmd/server