FROM golang:1.13 AS builder

WORKDIR /app
COPY http2https.go .
RUN CGO_ENABLED=0 go build -o app .

FROM alpine:latest
COPY --from=builder /app /
CMD ["/app"]
