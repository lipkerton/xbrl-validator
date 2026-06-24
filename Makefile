.PHONY: deps tools

deps:
	go mod tidy
tools:
	go install github.com/pressly/goose/v3/cmd/goose@latest