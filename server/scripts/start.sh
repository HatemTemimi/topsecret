#!/bin/bash

# Load environment variables from config/.env file
if [ -f ../config/.env.dev ]; then
  export $(grep -v '^#' ../config/.env.dev | xargs)
else
  echo "config/.env file not found"
  exit 1
fi

# Check if required environment variables are set
if [[ -z "$DATABASE_PORT" || -z "$DATABASE_USER" || -z "$DATABASE_PASSWORD" ]]; then
  echo "Required environment variables (DATABASE_PORT, DATABASE_USER, DATABASE_PASSWORD) are missing."
  exit 1
fi

# Kill any process currently using the database port
if lsof -i tcp:"$DATABASE_PORT" &>/dev/null; then
  echo "Port $DATABASE_PORT is in use. Terminating processes..."
  lsof -t -i tcp:"$DATABASE_PORT" | xargs kill -9
  echo "Processes terminated on port $DATABASE_PORT."
fi

# Check for existing Docker container with the name mongodb_container
container_id=$(docker ps -q --filter "ancestor=mongodb/mongodb-community-server")

if [ -n "$container_id" ]; then
  echo "Stopping existing Docker container with ID $container_id..."
  docker stop "$container_id" >/dev/null
  
  # Wait until the container has fully stopped
  while [ "$(docker inspect -f '{{.State.Running}}' "$container_id")" == "true" ]; do
    echo "Waiting for container $container_id to stop..."
    sleep 3
  done
  echo "Container $container_id stopped."

  # Remove the container after it has fully stopped
  echo "Removing Docker container with ID $container_id..."
  docker rm "$container_id"
  echo "Container $container_id removed."
fi


# Run MongoDB Docker container with environment variables
docker run -d -p "$DATABASE_PORT":27017 \
  -e MONGO_INITDB_ROOT_USERNAME="$DATABASE_USER" \
  -e MONGO_INITDB_ROOT_PASSWORD="$DATABASE_PASSWORD" \
  mongodb/mongodb-community-server

# Check if MongoDB Docker container started successfully
if [ $? -ne 0 ]; then
  echo "Failed to start MongoDB Docker container."
  exit 1
fi

# Wait for MongoDB to initialize
echo "Waiting for MongoDB to initialize..."
sleep 5

# Switch to cli directory and run main.go
cd ../cli || { echo "Directory cli/ not found"; exit 1; }

echo "Running main.go..."
go run main.go
