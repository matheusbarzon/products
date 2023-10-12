FROM golang:1.21 as builder

WORKDIR /app

COPY ./.env ./
COPY ./src .
COPY ./sql ./sql

RUN go mod download

RUN go build cmd/main.go

CMD ["./main"]