# Use multi-stage builds to optimize the image size
FROM golang AS builder

WORKDIR /app

# Copy only necessary files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend main.go

# Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/backend .

# Set any required environment variables here
# ENV SOME_ENV_VAR=value

# Run the application
ENTRYPOINT ["./backend"]
