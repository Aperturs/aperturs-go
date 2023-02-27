# Start with the base Golang image
FROM golang:latest

WORKDIR /app

# Copy the go.mod and go.sum files into the container
COPY go.mod go.sum ./

# Download and install the Go modules
RUN go mod download

# Build the application
COPY . .
RUN go build -o main .

# Expose port 8080 for the application
EXPOSE 8080

# Start the application
CMD ["./main"]
