FROM golang:alpine

# Set the Current Working Directory inside the container
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go app
RUN go build -o ./goapp

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go build`
CMD ["./goapp"]