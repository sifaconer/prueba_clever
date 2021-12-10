FROM golang:1.17.3-alpine3.14

RUN apk update && apk add --virtual build-dependencies build-base gcc

RUN mkdir /app
ADD . /app

WORKDIR /app
RUN go mod download
RUN go build -o main cmd/main.go

EXPOSE 8182
CMD ["/app/main"]