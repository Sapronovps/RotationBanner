FROM golang:1.24-alpine AS builder

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:3.18

RUN apk add --no-cache libc6-compat postgresql-client

COPY --from=builder /go/bin/goose /usr/local/bin/goose