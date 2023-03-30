FROM golang:1.19-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerService ./cmd/api

RUN chmod +x /app/brokerService

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerService /app

CMD ["./app/brokerService"]