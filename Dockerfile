# Use the official Golang image to create a binary for our application
FROM golang:1.17 as builder

WORKDIR /app

COPY . .

# Build the Go application
RUN go mod download && go build -o main .

# Use a minimalistic base image to run our application
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /app/main .

# Expose the port on which our application will run
EXPOSE 8000

# Command to run the application
CMD ["./main"]
