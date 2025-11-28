# Architecture Documentation

This document describes the core architectural patterns used in `go-api-scaffold`.

## Core Components

### 1. Dependency Injection (Dig)
The project uses [uber-go/dig](https://github.com/uber-go/dig) for dependency injection. The container is initialized in `internal/app/bootstrap.go`.

- **Bootstrap**: The `Bootstrap()` function creates a new container and provides core dependencies (Config, DB, BaseHandler, Router).
- **Invocation**: The `main` function invokes the container to start the HTTP server.

### 2. Handler Auto-Registration
To avoid manually registering every new handler in the router, this project uses a code generation approach.

- **Registry**: `internal/registry/registry.go` maintains a list of handler constructors.
- **Code Generation**: `internal/registry/generate_registry.go` scans `internal/handler` for types ending in `Handler` and automatically generates the registration code.
- **Usage**:
    1. Create a new handler in `internal/handler`.
    2. Run `go generate ./internal/registry`.
    3. The handler is added to `registry.List()`.
    4. `internal/router/router.go` iterates over this list and registers routes.

### 3. Routing (Chi)
[Chi](https://github.com/go-chi/chi) is used for HTTP routing.

- **Router Builder**: `internal/router/router.go` builds the main router.
- **Handler Interface**: Each handler must implement `RegisterRoutes(r chi.Router)`.
- **Swagger**: Swagger UI is automatically served at `/swagger/*`.

### 4. Configuration (Viper)
[Viper](https://github.com/spf13/viper) handles configuration.

- **Source**: Defaults to `config/config.yaml`.
- **Overrides**: Can be overridden by environment variables (e.g., `SERVER_PORT`) or command-line flags.
- **Structure**: Defined in `internal/config/config.go`.
