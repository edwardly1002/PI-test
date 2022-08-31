build:
	- mkdir bin
	go build -o ./bin ./cmd/emailapp

clean:
	rm -rf bin

test:
	go test -v ./...

example:
	@export TEMPLATE_FILE="asset/email_template.json" CUSTOMERS_FILE="asset/customers.csv" \
	OUTPUT_FILE="asset/output_emails.json" ERRORS_FILE="asset/errors.csv" && go run ./cmd/emailapp/.