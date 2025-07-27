# Docker Deployment

This directory contains Docker configuration files for running Fabric in containers.

## Files

- `Dockerfile` - Main Docker build configuration
- `docker-compose.yml` - Docker Compose stack configuration  
- `start-docker.sh` - Helper script to start the stack
- `README.md` - This documentation

## Quick Start

```bash
# Start the Docker stack
./start-docker.sh

# Or manually with docker-compose
docker-compose up -d

# View logs
docker-compose logs -f

# Stop the stack
docker-compose down
```

## Building

```bash
# Build the Docker image
docker build -t fabric .

# Or use docker-compose
docker-compose build
```

## Configuration

Make sure to configure your environment variables and API keys before running the Docker stack. See the main README.md for setup instructions.