# Start from the official Golang image to create a build artifact.
FROM golang:1.21.4 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod file from your project's sse-server directory
COPY sse-server/go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod file is not changed
RUN go mod download

# Copy the source from the sse-server directory to the Working Directory inside the container
COPY sse-server/ .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch for a smaller, final image
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
