FROM golang:1.13-alpine AS builder

WORKDIR /go/src/github.com/murilosrg/go-pay-me
ENV CC=gcc

COPY . .

RUN apk add --no-cache gcc musl-dev \
    && go build ./cmd/payme && mv payme /go/bin

FROM alpine:3.6

COPY --from=builder /go/bin/payme /usr/local/bin
COPY --from=builder /go/src/github.com/murilosrg/go-pay-me /payme
COPY configuration.example.yaml /etc/payme/configuration.yaml

WORKDIR /payme

CMD "payme" "-init"