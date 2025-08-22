# Stage 1: build aplikasi
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go.mod dan go.sum dulu supaya cache dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary
RUN go build -o app .

# Stage 2: image final yang ringan
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary hasil build
COPY --from=builder /app/app .

# Expose port aplikasi (ubah sesuai port aplikasi Go kamu)
EXPOSE 8080

# Jalankan aplikasi
CMD ["./app"]
