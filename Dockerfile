FROM golang:1.17

RUN make build

CMD ["./bin/emailapp"]