FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o expense-api ./cmd/api

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/expense-api .

COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./expense-api"]
