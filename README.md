# Affiliate Platform API

A robust, scalable backend service for an Affiliate Platform, built with Go using the **Modular Monolith** architecture.

## 🚀 Tech Stack

- **Language:** [Go](https://go.dev/) (1.21+)
- **Framework:** [Chi](https://github.com/go-chi/chi) (Lightweight, idiomatic Router)
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **Caching:** [Redis](https://redis.io/)
- **Infrastructure:** [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
- **Migration:** [Golang Migrate](https://github.com/golang-migrate/migrate)

## 🏗 Architecture: Modular Monolith

This project follows a **Modular Monolith** design pattern to ensure scalability and maintainability. It strikes a balance between the simplicity of a monolith and the decoupling of microservices.

### Key Concepts

- **Modules:** The application is divided into self-contained modules (e.g., `remote/modules/auth`, `user`, `product`, `affiliate`). Each module encapsulates its own domain logic, data access, and API handlers.
- **Shared Kernel:** Common utilities, DTOs, and infrastructure code reside in `internal/shared` or `pkg/`, accessible by all modules.
- **Dependency Injection:** Dependencies (like Database, Config) are explicitly passed to modules and services, making testing easier.
- **Clean Architecture:** Within each module, code is often organized into layers: `interfaces` (HTTP), `application` (Use Cases), `domain` (Business Logic), and `infrastructure` (Repositories).

### Directory Structure

```
├── cmd/api/            # Application entrypoint
├── configs/            # Configuration files
├── internal/
│   ├── config/         # Config loading logic
│   ├── infra/          # Infrastructure (DB, Cache, etc.)
│   ├── modules/        # Domain modules (The core logic)
│   │   ├── affiliate/
│   │   ├── auth/
│   │   ├── product/
│   │   └── user/
│   ├── server/         # HTTP Server setup & Router wiring
│   └── shared/         # Shared utilities
├── migrations/         # SQL Migration files
├── pkg/                # Public libraries
└── Makefile            # Command automation
```

## 🛠 Getting Started

### Prerequisites

- Go 1.21+
- Docker & Docker Compose
- Make

### Setup

1.  **Clone the repository**
    ```bash
    git clone <repository-url>
    cd golang-template
    ```

2.  **Environment Setup**
    Copy the example config (if available) or ensure `.env` is set up.
    _Currently, configuration defaults are handled in `internal/config/config.go`._

3.  **Start Infrastructure**
    Start PostgreSQL and Redis containers:
    ```bash
    make docker-up
    ```

4.  **Run Migrations**
    Apply database schemas:
    ```bash
    make migrate-up
    ```

5.  **Run the Application**
    ```bash
    make run
    ```
    The server will start at `http://localhost:8000`.

## 📜 Makefile Commands

The project uses a `Makefile` to simplify common tasks:

- `make run`: Run the application locally.
- `make build`: Compile the binary to `bin/api`.
- `make test`: Run unit tests with coverage.
- `make docker-up`: Start Docker containers (DB, Redis).
- `make docker-down`: Stop and remove containers.
- `make migrate-up`: Apply DB migrations.
- `make migrate-down`: Rollback DB migrations.
- `make lint`: Run golangci-lint.

## 🤝 Contributing

1.  Fork the repository.
2.  Create a feature branch (`git checkout -b feature/amazing-feature`).
3.  Commit your changes (`git commit -m 'Add amazing feature'`).
4.  Push to the branch (`git push origin feature/amazing-feature`).
5.  Open a Pull Request.
