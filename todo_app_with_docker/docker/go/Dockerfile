FROM golang:latest

WORKDIR /go/src/app/go

ADD ./go .

RUN go get -u https://github.com/gin-gonic/gin \
    go run main.go