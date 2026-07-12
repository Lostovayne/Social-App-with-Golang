# Social API

![Go](https://img.shields.io/badge/Go-1.26.1-00ADD8?logo=go&logoColor=white)
![Chi Router](https://img.shields.io/badge/Chi_v5-FF6F61?logo=lightning&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17-336791?logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-yellow)

RESTful API for a social platform. Built with **Go 1.26**, **Chi router**, and **PostgreSQL 17**.

## Architecture

```
HTTP Request
    │
    ▼
Chi Router ─── Middleware (logging, CORS, timeout)
    │
    ▼
Handlers (cmd/api/)
    │
    ▼
Store Layer (internal/store/) ─── SQL queries
    │
    ▼
PostgreSQL 17
```

The project follows a **layered architecture**:
- **Handlers** — HTTP layer: parse requests, validate input, return responses
- **Store** — Data access layer: SQL queries, connection pooling
- **DB** — Database connection management

## Quick Start

### Prerequisites

Install **make** (task runner), **direnv** (environment variables), and **migrate** (database migrations):

```bash
# Linux (ejemplo)
sudo apt install make direnv   # o tu gestor de paquetes
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

On Windows (with scoop):

```powershell
# Install scoop (Windows package manager)
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression

# Install make and migrate
scoop install make
scoop install migrate
```

Hook direnv into your shell — see [direnv docs](https://direnv.net/docs/hook.html).

### Setup

```bash
# Environment variables (direnv)
cp env.example .envrc
direnv allow

# Start PostgreSQL
make db-up

# Run migrations (creates all tables)
make migrate-up

# Seed test data (100 users, 50 posts, 100 comments)
make db-seed

# Start dev server (with hot reload)
make dev
```

API available at `http://localhost:8080/v1`

## API Endpoints

Base URL: `http://localhost:8080/v1`

### Health
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |

### Posts
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/posts/{postID}` | Get post by ID |
| POST | `/posts` | Create post |
| PATCH | `/posts/{postID}` | Update post |
| DELETE | `/posts/{postID}` | Delete post |

### Users
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users/{userID}` | Get user by ID |

### Comments
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/posts/{postID}/comments` | Create comment |

## Project Structure

```
├── cmd/
│   ├── api/              # Entry point, handlers, router
│   │   ├── main.go       # App bootstrap
│   │   ├── api.go        # Router & middleware setup
│   │   ├── posts.go      # Post handlers
│   │   ├── users.go      # User handlers
│   │   ├── health.go     # Health check
│   │   ├── json.go       # JSON encode/decode utilities
│   │   └── errors.go     # Error responses
│   ├── migrate/          # DB migrations
│   └── seed/             # Seed data
│
├── internal/
│   ├── db/               # DB connection pool
│   ├── env/              # Environment variables
│   └── store/            # Data access layer
│       ├── storage.go    # Storage interface
│       ├── users.go      # Users repository
│       ├── posts.go      # Posts repository
│       └── comments.go   # Comments repository
│
├── docker-compose.yml    # PostgreSQL container
├── Makefile              # Task runner
├── env.example           # Template for .envrc (direnv)
└── endpoints.http        # API test requests
```

## Configuration

Copy `env.example` to `.envrc` and run `direnv allow`. Variables are loaded by [direnv](https://direnv.net/) into your shell; `make` and Go commands inherit them automatically.

All config via environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `ADDR` | `:8081` | Server address |
| `DB_ADDR` | `postgres://admin:adminpassword@localhost:5432/social?sslmode=disable` | DB connection |
| `DB_MAX_OPEN_CONNS` | `30` | Max open connections |
| `DB_MAX_IDLE_CONNS` | `30` | Max idle connections |
| `DB_MAX_IDLE_TIME` | `15m` | Max idle time |
| `ENV` | `development` | Environment |

## Database Migrations

This project uses [golang-migrate](https://github.com/golang-migrate/migrate) for schema management. Migrations are versioned SQL files in `cmd/migrate/migrations/`.

### How It Works

```
cmd/migrate/migrations/
├── 000001_create_users.up.sql      # Creates users table
├── 000001_create_users.down.sql    # Drops users table
├── 000002_posts_create.up.sql      # Creates posts table
├── 000002_posts_create.down.sql    # Drops posts table
├── ...
└── 000008_add_indexes.up.sql       # Adds performance indexes
```

- **`.up.sql`** — applies the change (CREATE TABLE, ALTER, INSERT)
- **`.down.sql`** — reverts that exact change (DROP, ALTER back)
- Numbers (`000001`) define execution order
- `schema_migrations` table tracks applied versions

### Common Commands

```bash
make migrate-up           # Apply all pending migrations
make migrate-down         # Rollback last migration
make migrate-create NAME=name  # Create new migration pair (.up.sql + .down.sql)
```

### Creating a New Migration

```bash
make migrate-create NAME=add_notifications_table
```

This creates:
- `000009_add_notifications_table.up.sql`
- `000009_add_notifications_table.down.sql`

Edit the files with your SQL, then run `make migrate-up` to apply.

## Commands

```bash
# Development
make dev              # Hot reload with air
make run              # Run without hot reload
make build            # Compile binary
make clean            # Remove binaries

# Database
make db-up            # Start PostgreSQL container
make db-down          # Stop PostgreSQL container
make db-logs          # View DB logs
make db-seed          # Seed test data (100 users, 50 posts, 100 comments)

# Migrations
make migrate-up       # Apply all pending migrations
make migrate-down     # Rollback last migration
make migrate-create NAME=<name>  # Create new migration pair

# Quality
make test             # Run tests
make vet              # Go vet
make check            # fmt-go + fmt-sql + vet + test
make fmt-go           # Format Go code
make fmt-sql          # Format SQL files

# Tools
make tools            # Install air & goimports
make help             # List all commands
```

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Language | Go 1.26 |
| Router | Chi v5 |
| Database | PostgreSQL 17 |
| Migration | golang-migrate |
| Validator | go-playground/validator |
| Dev tools | air, make, direnv, Docker |

## License

MIT
