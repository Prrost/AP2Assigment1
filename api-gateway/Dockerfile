FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

ENV ENV=doc

COPY --from=builder /app/main .

COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]
