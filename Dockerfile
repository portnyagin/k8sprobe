# syntax=docker/dockerfile:1
FROM golang:1.17-alpine

RUN apk add --no-cache git
RUN apk add --no-cache bash

RUN go get github.com/caarlos0/env/v6

WORKDIR /app
