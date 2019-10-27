# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Pundlik Sarafdar <sarafdarpundlik@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Build Args
ARG LOG_DIR=/app/logs

# Create Log Directory
RUN mkdir -p ${LOG_DIR}

# Environment Variables
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log 

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

#Set env for arm64
# Build the Go app
RUN env GOOS=linux GOARCH=arm GOARM=7 go build -o main ./cmd

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]