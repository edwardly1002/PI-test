FROM golang:1.17

WORKDIR /myapp

CMD ["make", "all"]