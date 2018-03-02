.PHONY: test build

default: build

test:
	ginkgo -r

build: 
	go build -ldflags="-s -w"