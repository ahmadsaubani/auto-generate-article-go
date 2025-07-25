# Stage build
FROM golang:1.24.4-alpine AS builder

# Install tools (git, ca-certificates)
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy mod files dan download dependencies dulu untuk caching
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary statis (no CGO)
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o api-go

# Final image ringan
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/api-go .

EXPOSE 8080

CMD ["./api-go"]
