# Step 1: Build stage
FROM golang:1.23.5-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Move to the cmd/api directory where the main.go file is located
WORKDIR /app/cmd/api

# Build the Go application
RUN go build -o main .

# Step 2: Use a smaller base image to run go binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/cmd/api/main .
COPY --from=builder /app/.env .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the application
CMD ["./main"]