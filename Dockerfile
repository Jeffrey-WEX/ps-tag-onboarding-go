# Use a single base image that includes both Go build tools and runtime environment
FROM golang:1.22.0-bullseye AS builder

WORKDIR /app
# Copy go.mod and go.sum into /app, which is specified by WORKDIR
COPY go.mod .
COPY go.sum .
# Download dependencies in go.mod
RUN go mod download

# Copy everything in root of project (internal, test, cmd, etc.) into /app (WORKDIR)
COPY . .

# Show what content we have cloned inside of the image app for debugging.
RUN echo "Files inside of docker image:"
RUN ls

WORKDIR /app/cmd/user-api

# Create the main executable inside /app/cmd
RUN go build -o main .

# Set environment variables
ENV ENVIRONMENT=dockerfile_local_container
ENV DATABASE_URI=mongodb://localhost:27017/

# Expose the application port
EXPOSE 8080

# Set the entry point to the main executable, so when we run the container, it executes the executable.
ENTRYPOINT ["/app/cmd/user-api/main"]