FROM golang:1.21 as builder

WORKDIR /app

COPY go.* ./
COPY ./.env ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./pkg ./pkg

RUN go build cmd/main.go

CMD ["./main"]