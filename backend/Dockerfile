# Use the official Golang image as the base image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Install dependencies like git (if needed) for module dependencies
RUN apk add --no-cache git bash

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Copy the wait-for-it script into the container
COPY wait-for-it.sh /wait-for-it.sh

# Make the wait-for-it script executable
RUN chmod +x /wait-for-it.sh

# Build the Go application
RUN go build -o main .

# Expose the port your backend will run on
EXPOSE 8080

# Wait for MySQL service to be ready, then start the backend
CMD ["/wait-for-it.sh", "db:3306", "--", "./main"]
