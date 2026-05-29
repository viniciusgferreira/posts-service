# Build stage
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
COPY vendor/ vendor/

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /bin/posts-service ./cmd/...

# Runtime stage
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /bin/posts-service .

EXPOSE 8080

ENTRYPOINT ["./posts-service"]
