# Use Go image as the base image
FROM golang:1.22-alpine

# Set working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download dependencies and build the binary
RUN go mod download && go build -o app-load2

# Expose port for the application
EXPOSE 8000

# Run the application
CMD ["./app-load2"]
