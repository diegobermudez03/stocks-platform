FROM golang:1.24.2 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/app /app

ENTRYPOINT ["/app"]

EXPOSE 8000
