FROM golang:1.16

WORKDIR /app

# Install netcat
RUN apt-get update && apt-get install -y netcat

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Build the application
RUN go build -o /simulator ./cmd/simulator/main.go
RUN go build -o /processor ./cmd/processor/main.go

# Copy the entrypoint script
COPY deployments/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Default command, will be overridden in docker-compose.yml
CMD ["/simulator"]
