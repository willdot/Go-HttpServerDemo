FROM golang:1.12

ADD . /go/src/app
WORKDIR /go/src/app

ENV PORT=3001
