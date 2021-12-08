# syntax=docker/dockerfile:1

FROM golang:1.17.4-alpine3.15

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /registry-server

EXPOSE 8080

CMD [ "/registry-server" ]