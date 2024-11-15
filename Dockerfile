# Dockerfile
FROM golang:1.18-alpine

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 8080

# Run the Go application
CMD ["./main"]
