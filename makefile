.PHONY: run
BINARY_NAME=myapp
build:
	go build -o $(BINARY_NAME) ./cmd/web
run:build
	go run ./cmd/web
git:
	git init