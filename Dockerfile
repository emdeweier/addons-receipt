# Stage 1: build binary
FROM golang:1.21 AS builder

# Set working dir di dalam container build
WORKDIR /app

# Copy dependency terlebih dahulu agar cache efisien
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary dengan nama addons-receipt
RUN go build -o addons-receipt .

# Stage 2: image ringan untuk run
FROM debian:bookworm-slim

# Set working dir di container final
WORKDIR /app

# Copy binary dari stage build
COPY --from=builder /app/addons-receipt .

# Expose port aplikasi (ubah kalau app kamu pakai port lain)
EXPOSE 8080

# Command default
CMD ["./addons-receipt"]
