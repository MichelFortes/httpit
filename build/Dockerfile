# Build Stage
FROM golang:1.22.5 AS builder

WORKDIR /app

# Layer Caching: Copy modules manifests first
COPY go.mod ./
# COPY go.sum ./ # Uncomment when go.sum exists
RUN go mod download

# Copy source code
COPY . .

# Build the binary
# -ldflags="-w -s" reduces binary size
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o httpit ./cmd/httpit/main.go

# Final Stage
FROM alpine:latest

WORKDIR /app

# Install curl for debugging (optional, can be removed for lighter image)
RUN apk --no-cache add curl

# Copy binary from builder
COPY --from=builder /app/httpit .

ENTRYPOINT ["./httpit"]
