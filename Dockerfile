FROM golang:1.17

WORKDIR /myapp

RUN make build

CMD ["./bin/emailapp"]