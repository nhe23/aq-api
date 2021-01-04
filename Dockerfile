FROM golang:latest


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build ./cmd/server

ENV HOST=0.0.0.0

# host.docker.internal points to local machine. Change if db is deployed somewhere else.
ENV mongodb=mongodb://host.docker.internal:27018

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./server"]