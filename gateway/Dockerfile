FROM golang:1.20-alpine as builder
WORKDIR /app

COPY . ./

WORKDIR /app/gateway

RUN go mod tidy && go mod download

RUN go build -o /gateway ./cmd/gateway/main.go

FROM alpine

WORKDIR /

COPY --from=builder /gateway /gateway
COPY --from=builder /app/gateway/config.yml /config.yml

CMD ["/gateway"]
