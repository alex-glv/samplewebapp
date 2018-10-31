FROM golang:1.11 as builder

WORKDIR /src

COPY . .
RUN go mod download