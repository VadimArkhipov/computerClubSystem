FROM golang:1.19

WORKDIR /app

COPY ./src .

RUN go build -o task
