# Stage 1: Builder
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main


# Stage 2: Runner
FROM alpine

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080
