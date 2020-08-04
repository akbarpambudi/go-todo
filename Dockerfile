# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Akbar Pambudi <akbar.pambudi@amarbank.co.id>"

WORKDIR /app

COPY ./build/bin ./build/bin

EXPOSE 8080

CMD ["./buld/bin/todo"]