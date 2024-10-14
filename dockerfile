FROM golang:1.22.2-alpine AS builder

# Set destination for COPY
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN go build -o auth ./main/main.go

EXPOSE 2020

# Run
CMD ["./auth"]