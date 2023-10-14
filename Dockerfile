FROM golang:1.21 as builder

WORKDIR /app

COPY go.* ./
COPY ./.env ./
COPY ./src ./src
COPY ./sql ./sql

RUN go mod download

RUN go build src/cmd/main.go

CMD ["./main"]