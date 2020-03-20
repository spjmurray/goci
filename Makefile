all:
	go build -o build/bin/fib ./cmd/fib

test:
	go test -race -cover -coverprofile=cover.out ./...
	go tool cover -html cover.out -o cover.html
