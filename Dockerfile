# syntax=docker/dockerfile:1
FROM golang:1.25-alpine

# Install Air for hot reload
RUN apk add --no-cache git curl
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Copy go mod first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the code
COPY . .

EXPOSE 8080

# Use Air for live reload
CMD ["air", "-c", ".air.toml"]
