# Build stage
FROM golang:1.23-bookworm AS builder

WORKDIR /app

# Copy dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/go-buks .

# Runtime stage - minimal image with CA certificates
FROM gcr.io/distroless/static-debian12:nonroot

# Copy binary from builder
COPY --from=builder /app/go-buks /go-buks

# Expose application port
EXPOSE 8080

# Run as non-root user (nonroot user is built into distroless)
USER nonroot:nonroot

# Run the application
ENTRYPOINT ["/go-buks"]
