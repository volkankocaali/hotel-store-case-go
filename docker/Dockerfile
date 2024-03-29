# Builder stage
FROM golang:1.22-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server-linux ./cmd/api

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add tzdata
RUN cp /usr/share/zoneinfo/Europe/Istanbul /etc/localtime
RUN echo "Europe/Istanbul" > /etc/timezone
RUN apk del tzdata
WORKDIR /root/
COPY ./.env.dist ./.env
COPY --from=builder /app/server-linux ./cmd/api/server
EXPOSE 3000
CMD ["./cmd/api/server"]