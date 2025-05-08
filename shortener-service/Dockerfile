FROM golang:1.23.5-bookworm AS builder

WORKDIR /app

COPY go.mod  ./

RUN go mod download

COPY . .

RUN go build -o app

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/app .

ENTRYPOINT ["./app"]
