#!/bin/bash

# Load environment variables from mongo_credentials.txt
while IFS='=' read -r key value; do
    if [[ -n "$key" && -n "$value" ]]; then
        export "$key=$value"
    fi
done < "/c/Users/Rekanto/Desktop/employee-service/mongo_credentials.txt"

# Start services with Docker Compose
docker-compose up -d

# Wait for the web service to start
echo "Waiting for the web service to start..."
sleep 5

# Run Newman with Docker, mounting the Docker socket and files
docker run --network host \
  --name employee-service \
  -v /c/Users/Rekanto/Desktop/employee-service/collection/:/etc/postman \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -t postman/newman:latest run -r cli,json --reporter-json-export /etc/postman/newman-report.json /etc/postman/collection.json

# Shut down services with Docker Compose
docker-compose down
