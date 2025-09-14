# syntax=docker/dockerfile:1

FROM golang:1.25-alpine AS builder

WORKDIR /src

# Install build dependencies
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /fabric ./cmd/fabric

FROM alpine:latest

LABEL org.opencontainers.image.description="A Docker image for running the Fabric CLI. See https://github.com/danielmiessler/Fabric/tree/main/scripts/docker for details."

RUN apk add --no-cache ca-certificates \
    && mkdir -p /root/.config/fabric

COPY --from=builder /fabric /usr/local/bin/fabric

EXPOSE 8080

ENTRYPOINT ["fabric"]
