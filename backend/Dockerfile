FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./
# not running as root user, but as "appuser"
RUN go build -o backend-service .

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/backend-service /backend-service
#listen on port 8080 "internal" to the container, not exposed !!
EXPOSE 8080

CMD ["/backend-service"]