BINARY := bin/sqlfmt

.PHONY: build test fmt clean

build:
	go build -o $(BINARY) ./cmd/sqlfmt

test:
	go test ./...

fmt:
	go fmt ./...

clean:
	rm -rf bin