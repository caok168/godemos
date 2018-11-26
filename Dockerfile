FROM golang:1.10

WORKDIR /go/src/godemos
COPY . .

RUN go get -d -v ./...
