# Menggunakan image golang sebagai base image
FROM golang:1.20 AS builder

# Set working directory di dalam container
WORKDIR /app

# Menyalin go.mod dan go.sum (untuk dependensi)
COPY go.mod go.sum ./

# Install dependensi Go
RUN go mod tidy

# Menyalin seluruh kode aplikasi ke dalam container
COPY . .

# Membangun aplikasi Go
RUN go build -o myapp .

# Stage kedua untuk image runtime
FROM gcr.io/distroless/base

# Menyalin hasil build dari stage pertama
COPY --from=builder /app/addons-receipt /addons-receipt

# Menentukan perintah untuk menjalankan aplikasi
CMD ["/addons-receipt"]
