# Build stage
FROM golang:1.24.13-alpine3.22 AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/server/main.go

# Final stage
FROM alpine:3.22

# Change apk source to Aliyun
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# Install basic dependencies (optional, for debugging/timezone)
RUN apk --no-cache add tzdata

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./main"]
