# Code Cage API

This is a code cage API.

## Building and Running 🚀

A Dockerfile is included to build a Docker image of the application.

```bash
# Build Docker Image for production
docker build -t code-cage-api .

# Run Docker Container for development (hot reloading)
docker build -t code-cage-api -f Dockerfile.dev .

# Copy env file from .env.example
cp .env.example .env

# Run Docker Container for production
docker run -p 8080:8080 --env-file .env --name code-cage-api code-cage-api
```

## DB Setup 📦

### PostgreSQL

- Install PostgreSQL locally using Docker

```bash
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres --name postgres postgres
```

## Dependencies 📦

Go Modules are used for dependency management, which are listed in go.mod and go.sum.

## Contributing 🤝

We welcome contributions! Please see CONTRIBUTING.md for details on how to contribute.

## License 📄

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.
