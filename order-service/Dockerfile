FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=1


RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

ENV ENV=doc

RUN mkdir -p /root/DBDirectory

COPY --from=builder /app/main .

COPY --from=builder /app/.env .

RUN apk add --no-cache sqlite-libs

EXPOSE 8082

CMD ["./main"]
