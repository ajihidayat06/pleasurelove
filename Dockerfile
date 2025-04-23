# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

# Install necessary dependencies
RUN apk add --no-cache git

# Install Delve for debugging
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Debugging: Periksa isi direktori sebelum build
RUN echo "Isi direktori /app:" && ls -la /app
RUN echo "Isi direktori /app/cmd:" && ls -la /app/cmd

# Build the Go application
RUN go build -o main ./cmd/main.go

# Stage 2: Runtime
FROM golang:1.24-alpine

# Install necessary dependencies for runtime
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy the built binary and Delve debugger from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /go/bin/dlv /usr/local/bin/dlv

# Copy migrations directory
COPY migrations/ /app/migrations/

# Expose the application and debugging ports
EXPOSE 8080
EXPOSE 40000

# Debugging: Jalankan aplikasi secara otomatis tanpa menunggu interaksi
CMD ["sh", "-c", "if [ \"$DEBUG_MODE\" = \"true\" ]; then echo 'Running in debug mode...'; dlv exec ./main --headless --listen=:40000 --api-version=2 --accept-multiclient --log; else echo 'Running in normal mode...'; ./main; fi"]