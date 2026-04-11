# syntax=docker/dockerfile:1

ARG GO_VERSION=1.25.9
ARG ALPINE_VERSION=3.21

FROM alpine:${ALPINE_VERSION} AS builder

ARG GO_VERSION
ARG TARGETARCH
ENV GOROOT=/usr/local/go
ENV PATH=${GOROOT}/bin:${PATH}

WORKDIR /src

# Install build dependencies and the pinned Go toolchain.
RUN apk add --no-cache ca-certificates curl git tar \
    && case "${TARGETARCH:-$(apk --print-arch)}" in \
        amd64|x86_64) go_arch='amd64' ;; \
        arm64|aarch64) go_arch='arm64' ;; \
        *) echo "Unsupported TARGETARCH: ${TARGETARCH:-$(apk --print-arch)}" >&2; exit 1 ;; \
    esac \
    && curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-${go_arch}.tar.gz" -o /tmp/go.tgz \
    && rm -rf "${GOROOT}" \
    && tar -C /usr/local -xzf /tmp/go.tgz \
    && rm -f /tmp/go.tgz

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /fabric ./cmd/fabric

FROM alpine:${ALPINE_VERSION}

LABEL org.opencontainers.image.description="A Docker image for running the Fabric CLI. See https://github.com/danielmiessler/Fabric/tree/main/scripts/docker for details."

RUN apk add --no-cache ca-certificates python3 py3-pip \
    && pip3 install --no-cache-dir --break-system-packages 'setuptools>=78.1.1' \
    && pip3 install --no-cache-dir --break-system-packages 'yt-dlp>=2026.02.21' \
    && mkdir -p /root/.config/fabric \
    && apk upgrade --no-cache

COPY --from=builder /fabric /usr/local/bin/fabric

EXPOSE 8080

# Run as non-root user
RUN adduser -D appuser
USER appuser

ENTRYPOINT ["fabric"]
