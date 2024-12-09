# Directory containing the Docker Compose configuration
DOCKER_DIR :=db

# Default target: Bring up the Docker Compose services
dkup:
	docker-compose -f $(DOCKER_DIR)/docker-compose.yml up -d

# Stop and remove all services
dkdown:
	docker-compose -f $(DOCKER_DIR)/docker-compose.yml down

# Restart the services
restart:
	make dkdown
	make dkup

# View logs from the services
dklogs:
	docker-compose -f $(DOCKER_DIR)/docker-compose.yml logs -f

# Run a specific service command (e.g., make exec service=web cmd=bash)
dkexec:
	docker-compose -f $(DOCKER_DIR)/docker-compose.yml exec $(service) $(cmd)

# Build or rebuild services
dkbuild:
	docker-compose -f $(DOCKER_DIR)/docker-compose.yml build

# Remove unused Docker resources
dkprune:
	docker system prune -f

# Docker enter database 
dkdb:
   docker exec -it db-postgres-1 psql -U postgres



# Run paca stream
stream:
	go run cmd/paca-stream/stream.go


