# Stage build
FROM golang:1.24-alpine AS builder

# Set GOTOOLCHAIN untuk auto-download versi yang dibutuhkan
ENV GOTOOLCHAIN=auto

# Install tools yang diperlukan
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files dan download dependencies untuk layer caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy semua source code
COPY . .

# Build binary statis dengan optimasi
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -trimpath \
    -ldflags="-s -w -extldflags '-static'" \
    -a -installsuffix cgo \
    -o api-go .

# Final stage - distroless untuk keamanan maksimal
FROM gcr.io/distroless/static:nonroot

# Copy timezone data dari builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy CA certificates dari builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary dari builder stage
COPY --from=builder /app/api-go /

# Use non-root user untuk keamanan
USER nonroot:nonroot

# Expose port
EXPOSE 3000

# Set entrypoint
ENTRYPOINT ["/api-go"]
