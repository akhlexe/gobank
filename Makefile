# Main build target
build: clean
	go build -o bin/gobank

run: build
	@./bin/gobank

test:
	@go test -v ./...

# Clean artifacts
clean:
	rm -f bin/myapi