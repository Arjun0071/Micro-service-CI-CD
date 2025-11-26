# ---------- Stage 1: Build ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy whole microservice source code
COPY . .

# Install C compiler and musl-dev for CGO (needed by go-sqlite3)
RUN apk add --no-cache gcc musl-dev

# Enable CGO
ENV CGO_ENABLED=1

# Build binary
RUN GOOS=linux go build -o main .

# ---------- Stage 2: Run ----------
FROM alpine:latest

# Security: create non-root user
RUN adduser -D appuser

# Set working directory for app
WORKDIR /app

# Make /app writable
RUN chown -R appuser:appuser /app

# Copy compiled binary
COPY --from=builder /app/main .

# Expose service port
EXPOSE 8084

# Create /data for database mount
RUN mkdir /data && chown -R appuser:appuser /data

# Switch to non-root user
USER appuser

# Run the app
CMD ["./main"]

