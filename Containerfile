# ── Builder ──────────────────────────────────────────────────────────────────
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# Cache de módulos separado del código fuente
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /app/bin/api ./cmd/api

# ── Runner ────────────────────────────────────────────────────────────────────
FROM alpine:3.21 AS runner

# Certificados TLS y zona horaria
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/bin/api ./api

EXPOSE 8080

USER nobody

ENTRYPOINT ["./api"]
