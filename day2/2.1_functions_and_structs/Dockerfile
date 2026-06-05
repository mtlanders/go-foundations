# First attempt at creating a Docker image and running a container with a Go app

# Pull the Golang 1.26.4 image and name the stage "builder"
# Alpine is a lightweight Linux distro, good for containers
FROM golang:1.26.4-alpine AS builder

# Set the working dir inside the container as /app
WORKDIR /app

# Copies main.go from local machine into working dir in container
COPY go.mod main.go .

# Kick off the build to generate /app/main
RUN go build -o main .

# Pull the clean, minimal Alpine image for new stage
FROM alpine:latest

# Set the working dir inside the container as /app
WORKDIR /app

# Copy
COPY --from=builder /app/main .

# Now run the command from the native shell
CMD ["./main"]