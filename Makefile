.PHONY: test build clean run stop restart logs

# Variables
IMAGE_NAME = oak-nakama-image
CONTAINER_NAME = oak-nakama
COVERAGE_FILE = coverage.out
TAG = v1

# Run the tests
test:
	go test -v -coverprofile=$(COVERAGE_FILE) -cover ./...

# Generate the coverage summary
coverage: test
	go tool cover -func=$(COVERAGE_FILE)

# Build the Docker image for service to execute
build:
	docker compose build --no-cache

# Run the Docker containers for the service
start: build
	docker compose -f docker-compose.yml up -d

status:
	docker compose -f docker-compose.yml ps

# Lists the available make commands
list:
	@grep '^[^#[:space:]].*:' Makefile

# Stop the Docker container
stop:
	docker compose -f docker-compose.yml down

# Check the logs of the Docker container
logs:
	docker logs $(CONTAINER_NAME)

# Clean up Docker images and containers
clean:
	docker rm -f $(CONTAINER_NAME) || true
	docker rmi -f $(IMAGE_NAME):$(TAG) || true
	docker rmi -f $(IMAGE_NAME):latest || true
