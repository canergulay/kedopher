FROM golang:1.21.0-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN go mod download

RUN go build -o main

FROM alpine

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]
