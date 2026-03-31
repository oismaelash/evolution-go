#!/bin/bash

# Output colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Initializing Evolution Go via Docker ===${NC}\n"

# 1. Update git submodules (fixes empty folder error in whatsmeow-lib)
echo -e "${YELLOW}[1/4] Preparing dependencies (git submodules)...${NC}"
git submodule update --init --recursive
if [ $? -ne 0 ]; then
    echo -e "${RED}Error initializing submodules. Please check if git is installed.${NC}"
    exit 1
fi
echo -e "${GREEN}Dependencies prepared successfully!${NC}\n"

# 2. Create .env file if it doesn't exist
echo -e "${YELLOW}[2/4] Configuring environment variables...${NC}"
if [ ! -f .env ]; then
    if [ -f .env.example ]; then
        cp .env.example .env
        
        # Automatically adjust for Docker Compose
        sed -i 's/localhost:5432/postgres:5432/g' .env
        sed -i 's/localhost:5672/rabbitmq:5672/g' .env
        sed -i 's/localhost:9000/minio:9000/g' .env
        
        echo -e "${GREEN}.env file created from .env.example and adjusted for Docker!${NC}\n"
    else
        echo -e "${RED}.env.example file not found!${NC}"
        exit 1
    fi
else
    echo -e "${GREEN}.env file already exists, keeping current configuration.${NC}\n"
fi

# 3. Build and start containers via Docker Compose
echo -e "${YELLOW}[3/4] Building image and starting containers (This may take a few minutes)...${NC}"
if command -v docker-compose &> /dev/null; then
    docker-compose up -d --build
elif docker compose version &> /dev/null; then
    docker compose up -d --build
else
    echo -e "${RED}Docker Compose not found. Please install Docker Compose first.${NC}"
    exit 1
fi

if [ $? -ne 0 ]; then
    echo -e "${RED}Error starting Docker containers.${NC}"
    exit 1
fi
echo -e "${GREEN}Containers started successfully!${NC}\n"

# 4. Finalization
echo -e "${BLUE}=== All set! ===${NC}"
echo -e "Evolution Go and its dependencies are running in the background."
echo -e "\nServices available:"
echo -e "- Evolution Go API: ${GREEN}http://localhost:8080${NC}"
echo -e "- Swagger Docs:     ${GREEN}http://localhost:8080/swagger/index.html${NC}"
echo -e "- Manager UI:       ${GREEN}http://localhost:8080/manager/login${NC}"
echo -e "- RabbitMQ Admin:   ${GREEN}http://localhost:15672${NC} (admin/admin)"
echo -e "- MinIO Console:    ${GREEN}http://localhost:9001${NC} (minioadmin/minioadmin)"
echo -e "\nTo view API logs, run:"
echo -e "${YELLOW}docker compose logs -f evolution-go${NC}"
