Maritime Equipment Simulator
Description
This project simulates maritime equipment and processes sensor data. The system consists of simulators generating data and a processor storing this data into a PostgreSQL database.

Setup
Prerequisites
Docker
Docker Compose
Environment Variables
Create a .env file in the root of your project.

Build and Run
To build and run the Docker containers, execute:

css
Copy code
docker-compose up --build
Stopping the Containers
To stop the running containers, execute:

Copy code
docker-compose down
Check Database Records
To check the records in the PostgreSQL database, execute:

bash
Copy code
docker exec -it $(docker ps -qf "name=postgres") psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -c "SELECT * FROM sensor_data;"
This command connects to the PostgreSQL container and executes a SQL query to display all records in the sensor_data table. The docker exec command runs a command inside a running container. The $(docker ps -qf "name=postgres") part retrieves the container ID of the running PostgreSQL container. The psql command connects to the database and executes the given SQL query.