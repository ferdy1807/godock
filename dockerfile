# Gunakan image resmi Golang 1.24.2 sebagai base image
FROM golang:1.24.2-alpine as builder

# Set direktori kerja di dalam container
WORKDIR /app

# Copy file go mod dan sum
COPY go.mod go.sum ./

# Install dependensi
RUN go mod tidy

# Copy semua kode ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Gunakan image alpine sebagai base untuk hasil build
FROM alpine:latest

WORKDIR /root/

# Copy binary yang sudah dibangun dari stage builder
COPY --from=builder /app/main .

# Expose port 3000
EXPOSE 3000

# Perintah untuk menjalankan aplikasi
CMD ["./main"]