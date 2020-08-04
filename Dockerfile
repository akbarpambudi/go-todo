# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Akbar Pambudi <akbar.pambudi@amarbank.co.id>"

# Set the Current Working Directory inside the container
WORKDIR /app

COPY ./dst/bin ./dst/bin

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./dst/bin/godig-practice"]