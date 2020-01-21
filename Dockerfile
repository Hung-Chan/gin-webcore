# Golang 1.13.1 Alpine
FROM golang:1.13.1-alpine3.10

WORKDIR /go/src/gin-webcore

COPY ./ .

RUN apk add bash && apk add gcc && apk add g++

# RUN go mod init go-core

RUN go build

CMD ["go", "run", "main.go"]
