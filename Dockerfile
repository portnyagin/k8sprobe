# syntax=docker/dockerfile:1
FROM golang:1.17-alpine

RUN apk add --no-cache git
RUN apk add --no-cache bash
RUN apk add --no-cache curl

RUN go get github.com/caarlos0/env/v6

WORKDIR /app/src
RUN git clone https://github.com/portnyagin/k8sprobe
WORKDIR /app/src/k8sprobe
RUN go mod download
RUN go get github.com/portnyagin/k8sprobe/internal/app
RUN go build -o /app/service ./cmd/probe/main.go
WORKDIR /app
RUN rm -r ./src
#EXPOSE 8080
#CMD [ "./service" ]