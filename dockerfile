FROM golang:1.12

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT=3001
