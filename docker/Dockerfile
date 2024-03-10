# Builder stage
FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Linux için derleme
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server-linux ./cmd/api
# Windows için derleme
RUN CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -o server-windows.exe ./cmd/api

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
COPY --from=builder /app/server-windows.exe ./cmd/api/server.exe
EXPOSE 3000
CMD ["./cmd/api/server"]