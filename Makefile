build:
	- mkdir bin
	go build -o ./bin ./cmd/emailapp

clean:
	rm -rf bin