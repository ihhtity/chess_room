FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o chess-room-backend .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/chess-room-backend .
COPY --from=builder /app/config/config.yaml ./config/config.yaml

RUN mkdir -p /app/logs

EXPOSE 8080

CMD ["./chess-room-backend"]