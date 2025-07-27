#!/bin/bash

# Helper script to start the Fabric Docker stack

echo "Starting Fabric Docker stack..."
cd "$(dirname "$0")"
docker-compose up -d

echo "Fabric is now running!"
echo "Check logs with: docker-compose logs -f"
echo "Stop with: docker-compose down"