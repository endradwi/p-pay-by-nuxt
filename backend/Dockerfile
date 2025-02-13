# Use multi-stage builds to optimize the image size
FROM golang AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/backend .

ENTRYPOINT ["./backend"]
