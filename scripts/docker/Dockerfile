# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS builder

WORKDIR /src

# Install build dependencies
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /fabric ./cmd/fabric

FROM alpine:latest

RUN apk add --no-cache ca-certificates \
    && mkdir -p /root/.config/fabric

COPY --from=builder /fabric /usr/local/bin/fabric

EXPOSE 8080

ENTRYPOINT ["fabric"]
