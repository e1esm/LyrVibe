FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . ./

WORKDIR /app/artist-service

RUN go mod tidy && go mod download

RUN go build -o /artist-service ./cmd/artist-service/main.go

FROM alpine

WORKDIR /

COPY --from=builder /app/artist-service/config.yml /config.yml
COPY --from=builder /artist-service /artist-service


CMD ["/artist-service"]