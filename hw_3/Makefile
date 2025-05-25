.PHONY: build run test

build:
	go build -o main main.go text_analysis.go

run: build
	./main

test:
	go test -v