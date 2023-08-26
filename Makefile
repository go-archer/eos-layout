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
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server