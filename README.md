# Social API

![Go Version](https://img.shields.io/badge/Go-1.26.1-00ADD8?style=flat&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Ready-336791?style=flat&logo=postgresql)
![Chi Router](https://img.shields.io/badge/Router-Chi-blue)

A robust, performant RESTful API for a social platform built with Go, Chi router, and PostgreSQL.

## 🛠 Prerequisites

Before you begin, ensure you have the following installed:
- [Go](https://golang.org/doc/install) (v1.26.1 or higher)
- [Docker](https://docs.docker.com/get-docker/) & Docker Compose
- [Just](https://just.systems/man/en/) (A handy command runner)

*Optional but recommended tools:*
- `air` (for live reloading)
- `golang-migrate` (for managing database migrations)
- `goimports` (for formatting imports)

## 🚀 Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/Elevate-Techworks/social.git
cd social
```

### 2. Install development tools
We provide a convenient command to install necessary Go binaries (`air` and `goimports`):
```bash
just tools
```

### 3. Setup the Database
Start the PostgreSQL container using Docker Compose:
```bash
just db-up
```
*Note: To view database logs, you can run `just db-logs`. To stop the database, run `just db-down`.*

### 4. Run Migrations
Apply the database schema and structure:
```bash
just migrate-up
```

### 5. Run the Application
You can run the API in development mode with hot-reloading (via Air):
```bash
just dev
```
Alternatively, run it normally without hot-reload:
```bash
just run
```

The API should now be running. You can test the endpoints using the provided `endpoints.http` file with an HTTP client extension in your IDE (like REST Client for VSCode).

## 🗂 Project Structure

```text
.
├── cmd/
│   ├── api/                 # Main application entrypoint
│   └── migrate/migrations/  # SQL migration files
├── docs/                    # Project documentation
├── internal/                # Private application and business logic
├── scripts/                 # Utility scripts
├── docker-compose.yml       # Docker definitions for local infra (Postgres)
├── endpoints.http           # HTTP requests for quick local testing
├── go.mod / go.sum          # Go module dependencies
└── justfile                 # Task runner commands
```

## 📜 Available Commands

This project uses `just` to simplify common development tasks. Run `just` or `just help` to see all available commands.

### Development & Build
- `just dev` - Run the API with hot reload (`air`).
- `just run` - Run the API standardly.
- `just build` - Compile the binary into `./bin/main`.
- `just clean` - Remove generated binaries.

### Code Quality
- `just check` - Run formatting, go vet, and tests.
- `just test` - Run Go tests.
- `just vet` - Run `go vet`.
- `just fmt-go` - Format Go code.
- `just fmt-go-imports` - Sort and clean imports.
- `just fmt-sql` - Format SQL files using `pg_format` or `sqlfluff`.

### Database & Migrations
- `just db-up` / `db-down` / `db-logs` - Manage Postgres Docker container.
- `just migrate-up` - Apply pending migrations.
- `just migrate-down` - Revert the last applied migration.
- `just migrate-create <name>` - Create a new set of up/down SQL migration files.

## 🗄 Database Migrations

Migrations are managed via [golang-migrate](https://github.com/golang-migrate/migrate).

To create a new migration, run:
```bash
just migrate-create add_users_table
```
This will generate two new files (up and down) in `cmd/migrate/migrations`.

Ensure you have your environment variables set properly if you need to override the default local database address:
`DB_ADDR=postgres://user:password@localhost:5432/social?sslmode=disable`

## 🧪 Testing API Endpoints

Check out the `endpoints.http` file in the root directory. It contains raw HTTP requests that you can easily execute from your code editor (e.g., VS Code REST Client or IntelliJ HTTP Client) to test the active routes of the application.
