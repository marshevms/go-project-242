

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

lint:
	golangci-lint run ./...

run: 
	./bin/hexlet-path-size

test:
	go test ./...