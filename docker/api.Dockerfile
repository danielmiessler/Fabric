FROM golang:1.25-bullseye AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/fabric ./cmd/fabric

FROM debian:bookworm-slim AS runtime

WORKDIR /app

ENV HOME=/home/fabric
RUN useradd -m -d "$HOME" fabric

COPY docker/api-entrypoint.sh /usr/local/bin/entrypoint.sh
COPY --from=builder /out/fabric /usr/local/bin/fabric

RUN chmod +x /usr/local/bin/entrypoint.sh

USER fabric

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["fabric", "--serve", "--address", ":8080"]
