
# Maritime Equipment Simulator Project Documentation

## Overview
The Maritime Equipment Simulator is designed to simulate and monitor maritime equipment to enhance operational efficiency and maintenance strategies. This guide covers the setup and execution instructions for local development and deployment.

## Prerequisites
Before you begin, ensure you have the following installed:
- Docker
- Docker Compose
- Kubernetes CLI (kubectl)
- Minikube or Docker Desktop for Kubernetes support
- Poetry (for managing Python dependencies)

## Installation

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/Maritime-Equipment-Simulator.git
cd Maritime-Equipment-Simulator
```

### 2. Setup Python Environment (if applicable)
If your project involves Python components, set up the Python environment using Poetry:
```bash
poetry install
```

### 3. Start Local Development Environment
Use Docker Compose to start your local development environment:
```bash
docker-compose up -d
```

## Running the Application

### Local API and Simulator
Navigate to the `src` directory:
- For Go components:
  ```bash
  go run ./api/main.go
  ```

### Accessing the Application
The API will be available at `http://localhost:8080`. Use API endpoints to interact with the simulator, e.g., start/stop simulation or fetch data.

### Using Docker
Build the Docker image and run the container:
```bash
docker build -t maritime-equipment-simulator .
docker run -p 8080:8080 maritime-equipment-simulator
```

## Deploying to Kubernetes

### Start Minikube or Set up Kubernetes Cluster
```bash
minikube start
```

### Deploy to Kubernetes
Apply the Kubernetes configurations:
```bash
kubectl apply -f deployments/kubernetes.yaml
```

### Verify Deployment
Check the status of your Kubernetes pods:
```bash
kubectl get pods
```

## Accessing the System
After deployment, access the system via the Kubernetes load balancer or node port, depending on your configuration.

## Documentation and Support
Refer to the `/docs` directory for more detailed documentation on API usage, architecture, and other operational guides.

## Troubleshooting
If you encounter any issues during setup or deployment, ensure all prerequisites are correctly installed and configured. Check the Docker and Kubernetes logs for specific error messages.
