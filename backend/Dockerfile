FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./
RUN go build -o backend-service .

FROM alpine:3.22

WORKDIR /app

#[x] run as non-root user "newuser"
RUN adduser -D  newuser

COPY --from=builder /app/backend-service ./backend-service
#[x] listen on port 8080 "internal" to the container, not exposed !!
EXPOSE 8080

USER newuser

CMD ["./backend-service"]