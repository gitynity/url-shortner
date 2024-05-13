# Use Ubuntu as base image
FROM ubuntu:latest AS builder

# Install necessary packages for Golang
RUN apt-get update && apt-get install -y \
    wget \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# Download and install Golang
RUN wget https://golang.org/dl/go1.22.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz && \
    rm go1.22.3.linux-amd64.tar.gz

# Set the PATH environment variable
ENV PATH="/usr/local/go/bin:${PATH}"

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Debugging stage: Inspect directory structure
RUN ls -la
RUN pwd

# Build the Go application
RUN go build -o /app/bin/url-shortner ./cmd/api

# Start a new stage from scratch
FROM ubuntu:latest

# Set the working directory to /app
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/bin/url-shortner .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./url-shortner"]

