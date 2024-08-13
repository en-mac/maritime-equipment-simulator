# maritime-equipment-simulator
# Maritime-Equipment-Simulator Project Documentation

## Overview:
The Maritime-Equipment-Simulator is designed to simulate and monitor maritime equipment for operational efficiency and maintenance strategies. This guide covers the setup and execution instructions for local development and deployment.

# Prerequisites:
## Before you begin, ensure you have the following installed:

* Docker
* Docker Compose
* Kubernetes CLI (kubectl)
* Minikube or Docker Desktop (for Kubernetes support)
* Python (if using Python components)
* Poetry (for managing Python dependencies)

# Installation:

## Clone the Repository

* Run git clone https://github.com/yourusername/Maritime-Equipment-Simulator.git
* Navigate to the project directory with cd Maritime-Equipment-Simulator

## Setup Python Environment (if applicable)

* If your project involves Python components, set up the Python environment using Poetry by running poetry install
Start Local Development Environment

* Use Docker Compose to start your local development environment, which includes any necessary services like RabbitMQ or Kafka, by running docker-compose up -d

## Running the Application:

### Local API and Simulator

* Navigate to the src directory.
* For Go components, run go run ./api/main.go
* For Python components, run poetry run python ./api/app.py

### Accessing the Application

* The API will be available at http://localhost:8080 (adjust the port as configured).
* Use API endpoints to interact with the simulator, e.g., start/stop simulation or fetch data.

### Using Docker

* Build the Docker image by running docker build -t maritime-equipment-simulator .
* Run the Docker container by running docker run -p 8080:8080 maritime-equipment-simulator

### Deploying to Kubernetes:

#### Start Minikube or Set up Kubernetes Cluster

Run minikube start

#### Deploy to Kubernetes

Apply the Kubernetes configurations to set up the necessary deployments and services by running kubectl apply -f deployments/kubernetes.yaml

#### Verify Deployment

Check the status of your Kubernetes pods to ensure everything is running smoothly by running kubectl get pods

#### Accessing the System:

After deployment, access the system via the Kubernetes load balancer or node port, depending on your configuration.

#### Documentation and Support:

Refer to the /docs directory for more detailed documentation on API usage, architecture, and other operational guides.

#### Troubleshooting:

If you encounter any issues during setup or deployment, ensure all prerequisites are correctly installed and configured. Check the Docker and Kubernetes logs for specific error messages.