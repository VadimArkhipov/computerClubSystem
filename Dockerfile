FROM golang:1.19

WORKDIR /app

COPY ./src .
COPY ./tests ./tests

RUN go build -o task
