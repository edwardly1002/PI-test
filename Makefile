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
	export TEMPLATE_FILE="asset/email_template.json" && \
	export CUSTOMERS_FILE="asset/customers.csv" && \
	export OUTPUT_FILE="asset/output_emails.json" && \
	export ERRORS_FILE="asset/errors.csv" && \
	go run ./cmd/emailapp/.

ENV_FILE = "env.list"