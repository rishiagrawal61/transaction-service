# Stage 1: Build the binary
FROM golang:1.24-alpine AS builder

# Install Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# IMPORTANT: Disable optimizations (-N) and inlining (-l)
RUN CGO_ENABLED=0 GOOS=linux go build -gcflags "all=-N -l" -o main .

# Stage 2: Final lightweight image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /go/bin/dlv /usr/local/bin/dlv
EXPOSE 8080 2345

CMD ["/app/main"]