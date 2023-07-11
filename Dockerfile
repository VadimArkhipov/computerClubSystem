FROM golang:1.19

WORKDIR /app

COPY ./src .
COPY . ./tests

RUN go build -o task
