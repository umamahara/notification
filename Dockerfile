# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Step 1: Build Stage for Golang Application
FROM golang:latest


# Add Maintainer Info
LABEL maintainer="Uma Mahara <uma.mahara7@gmail.com>"

# Set the working directory in the container
WORKDIR /application

# Copy go modules files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o main

CMD ["./main"]
