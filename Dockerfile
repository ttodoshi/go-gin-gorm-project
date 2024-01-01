FROM golang:1.21-alpine AS builder

WORKDIR /usr/local/src/app

COPY .env.docker ./
COPY go.mod go.sum ./
COPY internal internal
COPY pkg pkg
COPY cmd cmd

RUN go mod tidy