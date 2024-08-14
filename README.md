
# Maritime Equipment Simulator

This is a golang project that simulates maritime equipment sensors and processes the generated data using RabbitMQ and PostgreSQL.


## Setup and Running

### Prerequisites

- Docker
- Docker Compose

### Configuration

Create a `.env` file in the root directory with the following variables:

```
# RabbitMQ Configuration
RABBITMQ_HOST=
RABBITMQ_PORT=
RABBITMQ_URL=

# PostgreSQL Configuration
POSTGRES_HOST=
POSTGRES_PORT=
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=

# Simulator Periods
PERIOD1=
PERIOD2=

# Simulator IDs
SIMULATOR_ID1=
SIMULATOR_ID2=

# PostgreSQL Connection String
POSTGRES_CONN_STR=
```

### Building and Running the Services

1. **Build the Docker images:**

   ```sh
   docker-compose build
   ```

2. **Run the Docker containers:**

   ```sh
   docker-compose --env-file ../.env up
   ```

### Stopping the Services

To stop the running services, use:

```sh
docker-compose down
```

### Querying the Database

You can query the PostgreSQL database to check the sensor data:

```sh
docker exec -it $(docker ps -qf "name=deployments-postgres-1") psql -U user -d maritime -c "SELECT * FROM sensor_data;"
```

This command will execute a SQL query inside the PostgreSQL container to retrieve all records from the `sensor_data` table.

## Explanation

This project consists of two main components: the simulator and the processor.

- **Simulator:** Generates random sensor data and sends it to RabbitMQ. Two instances of the simulator run in parallel, each with a different period for generating data.
- **Processor:** Consumes the sensor data from RabbitMQ, processes it, and stores it in a PostgreSQL database.

### Simulator

The simulator generates random sensor data values and assigns an alert level based on the value. If the value is above 90, the alert level is set to "high"; otherwise, it is set to "none". The data is then published to a RabbitMQ queue.

### Processor

The processor consumes messages from the RabbitMQ queue, processes the sensor data, and inserts it into the PostgreSQL database. Each record includes the sensor value, alert level, creation timestamp, and the simulator ID that generated the data.

This setup allows for simulating a real-time data processing pipeline using Docker containers, RabbitMQ, and PostgreSQL.
