FROM golang:1.22-alpine AS builder

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod download
RUN go build -o broker-service ./cmd/api
RUN chmod +x broker-service

FROM alpine:latest
COPY --from=builder /app/broker-service /app/broker-service
WORKDIR /app
ENTRYPOINT ["./broker-service"]