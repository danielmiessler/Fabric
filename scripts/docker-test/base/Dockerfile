FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/fabric ./cmd/fabric
COPY ./internal ./internal
RUN go build -o fabric ./cmd/fabric

FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Create a test user
RUN adduser -D -s /bin/sh testuser

# Switch to test user
USER testuser
WORKDIR /home/testuser

# Set environment variables for the test user
ENV HOME=/home/testuser
ENV USER=testuser

COPY --from=builder /app/fabric .

# Create fabric config directory and empty .env file
RUN mkdir -p .config/fabric && touch .config/fabric/.env

ENTRYPOINT ["./fabric"]
