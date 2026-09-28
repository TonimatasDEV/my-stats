FROM golang:1.27-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o my-stats

FROM alpine:latest
COPY --from=builder /app/my-stats /my-stats
EXPOSE 8088
ENV GIN_MODE=release
CMD ["/my-stats"]
