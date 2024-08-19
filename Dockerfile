# Builder stage
FROM golang:1.23rc2-alpine3.20 AS builder

WORKDIR /app

# Copy application source code
COPY . .

# Install curl for downloading additional binaries
# RUN apk add --no-cache curl

# Build the Go application
RUN go build -o main main.go

# Download and extract the migrate binary into a known directory
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar -xz -C /app/migrate --strip-components=1

# Run stage
FROM alpine:3.20

WORKDIR /app

# Copy application binary and migrate binary from builder stage
COPY --from=builder /app/main .
# COPY --from=builder /app/migrate /app/migrate

# Copy environment files, scripts, and migration files
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY migrate ./migrate
COPY db/migration ./migration

# Make sure the scripts and migrate binary are executable
RUN chmod +x start.sh wait-for.sh ./migrate

# Expose port
EXPOSE 8080

# Define entrypoint
ENTRYPOINT ["/app/start.sh"]
CMD ["/app/main"]
