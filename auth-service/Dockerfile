FROM golang:1.20-alpine as builder
WORKDIR /app

COPY . ./

WORKDIR /app/auth-service

RUN go mod tidy && go mod download

RUN go build -o /auther ./cmd/auth-service/main.go

FROM alpine

WORKDIR /

COPY --from=builder /auther /auther
COPY --from=builder /app/auth-service/config.yml /config.yml
COPY --from=builder /app/auth-service/postgres.env /postgres.env
COPY --from=builder /app/auth-service/redis.env /redis.env

CMD ["/auther"]