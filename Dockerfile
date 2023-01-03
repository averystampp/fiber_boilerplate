# syntax=docker/dockerfile:1
FROM golang:1.18-alpine
RUN mkdir /app
WORKDIR /app
ADD . /app

RUN go mod download

RUN apk add build-base
RUN go build -o /main
CMD [ "/app/main" ]