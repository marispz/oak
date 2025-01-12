
# Nakama RPC Plugin Example

This project demonstrates how to implement custom RPCs for the Nakama server using Go. It includes Docker configurations to build and run the necessary services and infrastructure for testing.

## Table of Contents
- [Overview](#overview)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Setup Instructions](#setup-instructions)
- [Running the Application](#running-the-application)
- [Testing](#testing)
- [Docker Setup](#docker-setup)
- [License](#license)

## Overview

This project provides a basic example of extending Nakama server functionalities with custom RPCs implemented in Go. The `rpc` directory contains the Go files for the custom RPCs. It includes Docker support for easy setup and management of the necessary infrastructure, including the Nakama server and a PostgreSQL database.

### Features:
- Implement custom RPCs (`account_metadata_update`, `game_configuration_read`, etc.)
- Docker Compose setup for running Nakama server and PostgreSQL
- Go-based unit tests with coverage tracking
- Build and run the plugin with Nakama server in Docker containers

## Project Structure

```
.
├── Dockerfile                # Dockerfile for building the Nakama plugin
├── Makefile                  # Makefile for managing common commands (build, test, etc.)
├── README.md                 # Project documentation
├── common                    # Common shared code for the project
│   ├── constants.go
│   ├── errors.go
│   ├── status_enum.go
│   └── types.go
├── coverage.out              # Test coverage output file
├── docker-compose.yml        # Docker Compose file for running the service
├── go.mod                    # Go modules file
├── go.sum                    # Go checksum file
├── hook                      # Contains hook-related logic (e.g., user initialization)
│   └── user_initializer.go
├── local.yml                 # Nakama configuration file
├── main.go                   # Entry point for the plugin
├── mocks                     # Mock files for testing
│   ├── Logger.go
│   └── NakamaModule.go
└── rpc                       # Custom RPC implementations and tests
    ├── account_metadata_update.go
    ├── account_metadata_update_test.go
    ├── config
    │   └── game_config.json
    ├── game_configuration_read.go
    ├── game_configuration_read_test.go
    ├── s2s_read_stats.go
    └── s2s_read_stats_test.go
```

## Prerequisites

To run this project locally, you need to have the following installed:
- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/dl/) (Go 1.23.3)

## Setup Instructions

1. Clone the repository:

    ```bash
    git clone https://github.com/marispz/oak.git
    cd oak
    ```

2. Install the Go dependencies:

    ```bash
    go mod tidy
    ```

3. Build the Docker images and start the services:

    If you have `make` installed (Unix-like systems):
    ```bash
    make start
    ```

    If you don't have `make`, you can run the equivalent Docker Compose command directly:

    On Windows (in CMD or PowerShell):
    ```bash
    docker-compose -f docker-compose.yml up -d --build
    ```

    This will start the Nakama server and PostgreSQL in Docker containers.

4. Verify the setup:

    If using `make`:
    ```bash
    make status
    ```

    If not using `make`, run the Docker Compose command directly:

    On Windows (in CMD or PowerShell):
    ```bash
    docker-compose -f docker-compose.yml ps
    ```

## Running the Application

Once the services are up, you can interact with the Nakama server and test the custom RPCs. The server will be available at `http://localhost:7350`.

To view logs for the Nakama server:

If using `make`:
```bash
make logs
```

If not using `make`, run the Docker command directly:

On Windows (in CMD or PowerShell):
```bash
docker logs oak-nakama
```

## Testing

To run tests for the custom RPCs, use the following command:

If using `make`:
```bash
make test
```

If not using `make`, run the Go test command directly:

On Windows (in CMD or PowerShell):
```bash
go test -v -coverprofile=coverage.out -cover ./...
```

This will execute the Go tests and generate a coverage report in `coverage.out`.

To view the test coverage:

If using `make`:
```bash
make coverage
```

If not using `make`, run the Go tool command directly:

On Windows (in CMD or PowerShell):
```bash
go tool cover -func=coverage.out
```

## Docker Setup

### Building the Docker Image

To build the Docker image for the Nakama plugin, run:

If using `make`:
```bash
make build
```

If not using `make`, run the Docker Compose command directly:

On Windows (in CMD or PowerShell):
```bash
docker-compose -f docker-compose.yml build --no-cache
```

### Stopping the Containers

To stop the Docker containers:

If using `make`:
```bash
make stop
```

If not using `make`, run the Docker Compose command directly:

On Windows (in CMD or PowerShell):
```bash
docker-compose -f docker-compose.yml down
```

### Cleaning up Docker Resources

To remove the Docker containers and images:

If using `make`:
```bash
make clean
```

If not using `make`, run the following Docker commands directly:

On Windows (in CMD or PowerShell):
```bash
docker rm -f oak-nakama || true
docker rmi -f oak-nakama-image:v1 || true
docker rmi -f oak-nakama-image:latest || true
```

### Invoking Custom RPCs
