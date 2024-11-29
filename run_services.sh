#!/bin/bash

# Define colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No color

# Network configuration
NETWORK_NAME="my_app_network"

# MySQL configuration
MYSQL_CONTAINER_NAME="mysql_db"
MYSQL_ROOT_PASSWORD="mo1236"
MYSQL_DATABASE="Ecommerce"

# Backend configuration
BACKEND_IMAGE_NAME="backend-image"
BACKEND_CONTAINER_NAME="backend-container"

# Frontend configuration
FRONTEND_IMAGE_NAME="frontend-image"
FRONTEND_CONTAINER_NAME="frontend-container"

echo -e "${GREEN}Starting the setup process...${NC}"

# Step 0: Create a custom Docker network (if it doesn't already exist)
echo -e "${GREEN}Creating custom Docker network '${NETWORK_NAME}'...${NC}"
docker network inspect ${NETWORK_NAME} >/dev/null 2>&1 || \
docker network create ${NETWORK_NAME} || {
    echo -e "${RED}Failed to create Docker network.${NC}"
    exit 1
}

# Step 1: Stop and remove old MySQL container (if it exists)
echo -e "${GREEN}Stopping and removing old MySQL container (if any)...${NC}"
docker rm -f ${MYSQL_CONTAINER_NAME} 2>/dev/null || echo -e "${RED}No existing MySQL container to remove.${NC}"

# Step 2: Start MySQL
echo -e "${GREEN}Starting MySQL service...${NC}"
docker run -d \
  --name ${MYSQL_CONTAINER_NAME} \
  --network ${NETWORK_NAME} \
  -e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
  -e MYSQL_DATABASE=${MYSQL_DATABASE} \
  -p 3307:3306 \
  -v package_tracking_system_db_data:/var/lib/mysql \
  mysql:8.0 || {
    echo -e "${RED}Failed to start MySQL container.${NC}"
    exit 1
}

# Step 3: Stop and remove old backend container (if it exists)
echo -e "${GREEN}Stopping and removing old backend container (if any)...${NC}"
docker rm -f ${BACKEND_CONTAINER_NAME} 2>/dev/null || echo -e "${RED}No existing backend container to remove.${NC}"

# Step 4: Build and run the backend
echo -e "${GREEN}Building backend Docker image...${NC}"
docker build -t ${BACKEND_IMAGE_NAME} ./backend || {
    echo -e "${RED}Failed to build backend Docker image.${NC}"
    exit 1
}

echo -e "${GREEN}Starting backend service...${NC}"
docker run -d \
  --name ${BACKEND_CONTAINER_NAME} \
  --network ${NETWORK_NAME} \
  -p 8080:8080 \
  -e DB_HOST=${MYSQL_CONTAINER_NAME} \
  -e DB_PORT=3306 \
  -e DB_USER=root \
  -e DB_PASSWORD=${MYSQL_ROOT_PASSWORD} \
  -e DB_NAME=${MYSQL_DATABASE} \
  ${BACKEND_IMAGE_NAME} || {
    echo -e "${RED}Failed to start backend container.${NC}"
    exit 1
}

# Step 5: Stop and remove old frontend container (if it exists)
echo -e "${GREEN}Stopping and removing old frontend container (if any)...${NC}"
docker rm -f ${FRONTEND_CONTAINER_NAME} 2>/dev/null || echo -e "${RED}No existing frontend container to remove.${NC}"

# Step 6: Build and run the frontend
echo -e "${GREEN}Building frontend Docker image...${NC}"
docker build -t ${FRONTEND_IMAGE_NAME} ./frontend || {
    echo -e "${RED}Failed to build frontend Docker image.${NC}"
    exit 1
}

echo -e "${GREEN}Starting frontend service...${NC}"
docker run -d \
  --name ${FRONTEND_CONTAINER_NAME} \
  --network ${NETWORK_NAME} \
  -p 4200:4200 \
  -v $(pwd)/frontend:/app \
  -v /app/node_modules \
  ${FRONTEND_IMAGE_NAME} || {
    echo -e "${RED}Failed to start frontend container.${NC}"
    exit 1
}

# Step 7: Display running containers
echo -e "${GREEN}Checking running containers...${NC}"
docker ps

echo -e "${GREEN}Setup complete! Backend is running on port 8080, Frontend on port 4200, and MySQL on port 3307.${NC}"
