FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o auth-service

# Expose the application port
EXPOSE 8081

# Start the application
CMD ["./auth-service"]