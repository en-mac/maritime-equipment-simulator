
# Maritime Equipment Simulator

This project simulates maritime equipment sensor data and processes the data using a Go application.

## Prerequisites

- Docker
- Docker Compose

## Setup

1. Clone the repository:

   ```
   git clone https://github.com/en-mac/maritime-equipment-simulator.git
   cd maritime-equipment-simulator
   ```

2. Create a `.env` file in the `deployments` directory with the following environment variables:

   ```
   # RabbitMQ Configuration
   RABBITMQ_HOST=
   RABBITMQ_PORT=
   POSTGRES_CONN_STR=
   RABBITMQ_URL=
   ```

   Replace the empty values with your configuration.

3. Build and start the services using Docker Compose:

   ```
   cd deployments
   docker-compose --env-file .env up --build
   ```

4. To stop the services, press `CTRL+C` and then run:

   ```
   docker-compose down
   ```

## Project Structure

- `cmd/processor/main.go`: Entry point for the processor service.
- `cmd/simulator/main.go`: Entry point for the simulator service.
- `internal/processor/processor.go`: Contains the logic for the processor service.
- `internal/simulator/simulator.go`: Contains the logic for the simulator service.
- `deployments/`: Contains Docker-related files and configurations.
  - `Dockerfile`: Dockerfile to build the Go application.
  - `docker-compose.yml`: Docker Compose configuration.
  - `entrypoint.sh`: Entrypoint script for the Docker containers.
  - `init/`: Contains initialization SQL scripts for PostgreSQL.
- `go.mod` and `go.sum`: Go module dependencies.

## Notes

- Ensure that the `.env` file is not committed to version control by adding it to `.gitignore`.
- The services will automatically retry connecting to RabbitMQ and PostgreSQL until they are available.
