# Use official golang image as builder
FROM golang:1.24.2-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o fabric ./cmd/fabric

# Use scratch as final base image
FROM alpine:latest

# Copy the binary from builder
COPY --from=builder /app/fabric /fabric

# Copy patterns directory
COPY patterns /patterns

# Ensure clean config directory and copy ENV file
RUN rm -rf /root/.config/fabric && \
    mkdir -p /root/.config/fabric
COPY ENV /root/.config/fabric/.env

# Add debug commands
RUN ls -la /root/.config/fabric/

# Expose port 8080
EXPOSE 8080

# Run the binary with debug output
ENTRYPOINT ["/fabric"]
CMD ["--serve"] 
