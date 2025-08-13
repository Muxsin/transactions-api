FROM golang:1.24-alpine

WORKDIR /code

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

EXPOSE 8080

CMD go run cmd/transactions-api/main.go
