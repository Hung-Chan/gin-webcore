# Golang 1.13.1 Alpine
FROM golang:1.13.1-alpine3.10

WORKDIR /go/src/gin-webcore

COPY ./ .

RUN ls

RUN apk add bash

# RUN go mod init go-core

RUN go build

CMD ["go", "run", "./main.go"]
