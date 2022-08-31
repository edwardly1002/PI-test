all:
	make build
	./bin/emailapp

build:
	- mkdir bin
	go build -o ./bin ./cmd/emailapp

clean:
	rm -rf bin

test:
	go test -v ./...

example:
	@export $(cat {{ENV_FILE}}) && go run ./cmd/emailapp/.

ENV_FILE = "env.list"