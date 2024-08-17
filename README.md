

# TODO - Parameterize quantity of simulators. Each simulator should have its own period, id, and IP address. The simulator should be able to generate data for multiple sensors.

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

```sh
docker-compose down
docker-compose --env-file ../.env up --build
```

### Querying the Database

You can query the PostgreSQL database to check the sensor data:

```sh
docker exec -it $(docker ps -qf "name=deployments-postgres-1") psql -U user -d maritime -c "SELECT * FROM salinity_data;"
```

This command will execute a SQL query inside the PostgreSQL container to retrieve all records from the `salinity_data` table.

## Explanation

This project consists of two main components: the simulator and the processor.

- **Simulator:** Generates random sensor data and sends it to RabbitMQ. Two instances of the simulator run in parallel, each with a different period for generating data.
- **Processor:** Consumes the sensor data from RabbitMQ, processes it, and stores it in a PostgreSQL database.

### Simulator

The simulator generates random salinity values between 30 and 40. Alert levels are set as follows:

* low if the value is below 33.
* normal if the value is between 33 and 37.
* high if the value is above 37.

This data is then published to a RabbitMQ queue. Two instances of the simulator run in parallel, each with a different period for generating data.

### Processor

The processor consumes messages from the RabbitMQ queue, processes the salinity data, and stores it in a PostgreSQL database. Each record includes the salinity value, alert level, creation timestamp, and the simulator ID that generated the data.

This setup allows for simulating a real-time data processing pipeline using Docker containers, RabbitMQ, and PostgreSQL.
