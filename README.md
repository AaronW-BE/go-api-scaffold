# go-api-scaffold

**A modular Go API scaffold with Chi, Dig, Swagger, and graceful shutdown support.**

---

## Features

* **Chi Router**: Lightweight and modular HTTP routing
* **Dependency Injection**: Powered by [Uber Dig](https://github.com/uber-go/dig)
* **Automatic Handler Registration**: Scan and register handlers automatically
* **Swagger Documentation**: Auto-generate OpenAPI docs for your API
* **Graceful Shutdown**: Properly close server, database, and Redis connections
* **Configuration Management**: Supports config files and CLI overrides
* **Database & Redis Integration**: Easily inject DB and Redis clients into handlers

---

## Project Structure

```
go-api-scaffold/
├─ cmd/
│  └─ app/                # Main application entry point
├─ docs/                   # Swagger docs (generated)
├─ internal/
│  ├─ handler/             # API handlers
│  │  ├─ base.go
│  │  └─ user_handler.go
│  ├─ registry/            # Handler registry and code generation
│  ├─ router/              # Router builder
│  ├─ db/                  # Database and Redis clients
│  ├─ service/             # Business logic services
│  └─ util/                # Utility functions
├─ go.mod
├─ go.sum
└─ README.md
```

---

## Installation

```bash
git clone https://github.com/yourusername/go-api-scaffold.git
cd go-api-scaffold
go mod tidy
```

---

## Configuration

Default configuration file is located at `config/config.yaml`. You can override using command line flags:

```bash
go run ./cmd/app/main.go \
  -config=config/custom.yaml \
  -port=8081 \
  -dbhost=127.0.0.1 \
  -dbport=3306 \
  -dbuser=root \
  -dbpass=123456 \
  -dbname=test
```

---

## Running the Project

```bash
go run ./cmd/app/main.go
```

Server will start at the configured port (default `:8080`) and all handlers will be registered automatically.

### Graceful Shutdown

Press `Ctrl+C` to stop the server. The shutdown process will:

* Finish ongoing HTTP requests
* Close the database connection
* Close Redis connection
* Exit cleanly

---

## Swagger API Documentation

Generate docs using:

```bash
go generate ./internal/registry
swag init -g cmd/app/main.go --output ./docs
```

Then access Swagger UI (if integrated) via:

```
http://localhost:8080/swagger/index.html
```

---

## Adding a New Handler

1. Create a new handler file in `internal/handler`:

```go
package handler

import (
	"go-api-scaffold/internal/util"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Base *BaseHandler
}

func NewUserHandler(base *BaseHandler) Router {
	return &UserHandler{Base: base}
}

// GetUser @Summary Get user info
// @Param id path int true "用户ID"
// @Success 200 {string} string "ok"
// @Router /user/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if h.Base.GetDB() != nil {
		log.Println("database is ok")
	}

	util.Json(w, map[string]string{
		"id":   chi.URLParam(r, "id"),
		"name": r.URL.Query().Get("name"),
		"desc": "hello" + r.URL.Query().Get("name"),
	}, 0, "ok")
}

func (h *UserHandler) RegisterRoutes(r chi.Router) {
	r.Get("/user/{id}", h.GetUser)
}

```

2. Run `go generate ./generate/registry.go` to auto-register the handler.

---

## License

MIT License.

---

This scaffold provides a fully modular Go API framework that can be extended with new handlers, services, and middlewar
