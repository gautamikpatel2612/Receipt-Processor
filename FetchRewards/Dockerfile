# Use the official Go image as a base
FROM golang:1.24

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o receipt-server

# Expose port 8080 for the app
EXPOSE 8080

# Start the server
CMD ["./receipt-server"]