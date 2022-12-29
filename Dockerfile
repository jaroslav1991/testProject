FROM golang:1.19-alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git openssh

ENV APP_SRC_DIR=/go/src

RUN mkdir $APP_SRC_DIR/app
WORKDIR $APP_SRC_DIR/app

ADD . $APP_SRC_DIR/app

COPY . .

RUN go build -o service ./cmd/server

ENTRYPOINT "/go/src/app/service"
