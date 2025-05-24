.PHONY: build run test

build:
	go build -o main main.go unpack.go

run: build
	./main

test:
	go test -v