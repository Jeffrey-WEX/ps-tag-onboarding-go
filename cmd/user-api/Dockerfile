# This section builds docker Image #
ARG BUILDER_IMAGE_REPOSITORY=""
FROM ${BUILDER_IMAGE_REPOSITORY}golang:1.22 as builder

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




# This section builds Docker Container #
FROM debian:bookworm-20240110-slim

ENV ENVIRONMENT=dockerfile_local_container
ENV DATABASE_URI=mongodb://localhost:27017/

RUN mkdir /app
WORKDIR /app
RUN mkdir /cmd
WORKDIR /cmd
RUN mkdir /user-api

EXPOSE 8080

# Copy main exe from builder (image) into container
COPY --from=builder /app/cmd/user-api/main /cmd/user-api/main
# Set the entry point to the main executable, so when we run the container, it executes the executable.
ENTRYPOINT ["/cmd/user-api/main"]