# Stage 1: build aplikasi
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build binary dengan nama addons-receipt
RUN go build -o addons-receipt .

# Stage 2: image final
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary
COPY --from=builder /app/addons-receipt .

EXPOSE 8080

# Jalankan binary
CMD ["./addons-receipt"]
