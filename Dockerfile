FROM golang:1.25.6-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server cmd/server/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
COPY .env .
EXPOSE 8080
CMD ["./server"]