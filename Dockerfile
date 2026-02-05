FROM golang:1.25-alpine3.22 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go

FROM alpine:3.22

# RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

ENTRYPOINT ["./app"]