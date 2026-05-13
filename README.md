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

```bash
# Start PostgreSQL
just db-up

# Run migrations
just migrate-up

# Start dev server (with hot reload)
just dev
```

API available at `http://localhost:8081/v1`

## API Endpoints

Base URL: `http://localhost:8081/v1`

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
├── justfile              # Task runner
└── endpoints.http        # API test requests
```

## Configuration

All config via environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `ADDR` | `:8081` | Server address |
| `DB_ADDR` | `postgres://admin:adminpassword@localhost:5432/social?sslmode=disable` | DB connection |
| `DB_MAX_OPEN_CONNS` | `30` | Max open connections |
| `DB_MAX_IDLE_CONNS` | `30` | Max idle connections |
| `DB_MAX_IDLE_TIME` | `15m` | Max idle time |
| `ENV` | `development` | Environment |

## Commands

```bash
# Development
just dev          # Hot reload with air
just run          # Run without hot reload
just build        # Compile binary
just clean        # Remove binaries

# Database
just db-up        # Start PostgreSQL
just db-down      # Stop PostgreSQL
just db-logs      # View DB logs
just db-seed      # Seed test data

# Migrations
just migrate-up   # Apply migrations
just migrate-down # Rollback last migration
just migrate-create <name>  # New migration

# Quality
just test         # Run tests
just vet          # Go vet
just check        # fmt + vet + test
just fmt-go       # Format Go code
just fmt-sql      # Format SQL

# Tools
just tools        # Install air & goimports
```

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Language | Go 1.26 |
| Router | Chi v5 |
| Database | PostgreSQL 17 |
| Migration | golang-migrate |
| Validator | go-playground/validator |
| Dev tools | air, just, Docker |

## License

MIT
