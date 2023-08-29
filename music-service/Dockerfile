FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . ./

WORKDIR /app/music-service

RUN go mod tidy && go mod download

RUN go build -o /music-service ./cmd/music-service/main.go

FROM alpine

WORKDIR /

COPY --from=builder /app/music-service/config.yml /config.yml
COPY --from=builder /music-service /music-service
COPY --from=builder /app/music-service/internal/repository/migrations/ /db/migrations

CMD ["/music-service"]