# Stage 1: Build
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

WORKDIR /app/api
RUN GOOS=linux GOARCH=amd64 go build -o /submittal-tracker main.go

# Stage 2: Runtime
FROM alpine:3.21
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /submittal-tracker .
RUN chmod +x /app/submittal-tracker
EXPOSE 3000
CMD ["./submittal-tracker"]