all: build install

.PHONY: build
build:
	GO111MODULE=on go build ./...

.PHONY: install
install:
	GO111MODULE=on go install ./...