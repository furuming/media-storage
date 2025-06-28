FROM golang:latest

WORKDIR /go/app

COPY . .
RUN cp .env.sample .env
