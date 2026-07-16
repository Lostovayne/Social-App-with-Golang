# Social API

![Go](https://img.shields.io/badge/Go-1.26.1-00ADD8?logo=go&logoColor=white)
![Chi Router](https://img.shields.io/badge/Chi_v5-FF6F61?logo=lightning&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17-336791?logo=postgresql&logoColor=white)
![Devenv](https://img.shields.io/badge/devenv-1.0-02471E?logo=nixos&logoColor=white)
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

## Prerequisites

- [Nix](https://nixos.org/download.html)
- [devenv](https://devenv.sh/getting-started/)
- [direnv](https://direnv.net/docs/hook.html) (optional, for auto-activation)

## Quick Start

```bash
# 1. Clone and enter the project
git clone <repo-url> && cd social-app

# 2. Activate the environment (first time)
devenv shell

# 3. Configure local credentials (first time)
cp devenv.local.nix.example devenv.local.nix
# Edit devenv.local.nix with your real DB credentials

# 4. Start Postgres + API with hot reload
devenv up

# 5. In another terminal: run migrations + seed data
devenv tasks run app:setup
```

API available at `http://localhost:8080/v1`

## Commands

All commands run through devenv tasks. Run `devenv tasks list` to see everything.

### Setup & Run

```bash
devenv up                       # Postgres + API (hot reload)
devenv up postgres              # Database only
devenv shell                    # Enter dev shell

devenv tasks run app:setup      # Migrate + seed (first time)
devenv tasks run app:dev        # Hot reload (standalone)
devenv tasks run app:run        # Run without hot reload
```

### Database

```bash
devenv tasks run app:migrate-up     # Apply all pending migrations
devenv tasks run app:migrate-down   # Rollback last migration
devenv tasks run app:migrate-new -- <name>  # Create new migration pair
devenv tasks run app:seed           # Seed test data
```

### Build

```bash
devenv tasks run app:build          # Build all binaries (api + seed)
devenv tasks run app:build-api      # Build API binary only
devenv tasks run app:build-seed     # Build seed binary only
devenv tasks run app:clean          # Remove binaries
```

### Test & Lint

```bash
devenv tasks run app:test           # Run tests
devenv tasks run app:vet            # Go vet
devenv tasks run app:check          # fmt + vet + tests
devenv tasks run app:fmt            # Format Go code
devenv tasks run app:fmt-imports    # Format imports (goimports)
devenv tasks run app:fmt-sql        # Format SQL files
```

### Docs & HTTP Client

```bash
devenv tasks run app:docs           # Regenerate Swagger docs
devenv tasks run app:http           # Run all endpoints (kulala-cli)
devenv tasks run app:http-run -- <name>  # Run a single named request
```

### Swagger UI

Once the server is running, open: `http://localhost:8080/v1/swagger/index.html`

## API Endpoints

Base URL: `http://localhost:8080/v1`

### Health

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |

### Posts

| Method | Endpoint | Description |
| -------- | ---------- | ------------- |
| POST | `/posts` | Create a post |
| GET | `/posts/{postID}` | Get post by ID (with comments) |
| PATCH | `/posts/{postID}` | Update post (title/content) |
| DELETE | `/posts/{postID}` | Delete a post |

### Users

| Method | Endpoint | Description |
| -------- | ---------- | ------------- |
| GET | `/users/{userID}` | Get user profile |
| PUT | `/users/{userID}/follow` | Follow a user |
| PUT | `/users/{userID}/unfollow` | Unfollow a user |

### Feed

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users/feed` | Get authenticated user's feed |

Query params: `limit`, `offset`, `sort` (asc/desc), `tags` (comma-separated), `search`

## Project Structure

```
├── cmd/
│   ├── api/                # Entry point, handlers, router
│   │   ├── main.go         # App bootstrap, Swagger config
│   │   ├── api.go          # Router & middleware setup
│   │   ├── posts.go        # Post handlers
│   │   ├── users.go        # User handlers
│   │   ├── feed.go         # Feed handler
│   │   ├── health.go       # Health check
│   │   ├── json.go         # JSON encode/decode utilities
│   │   └── errors.go       # Error responses
│   ├── migrate/            # DB migrations
│   │   └── migrations/     # Versioned SQL files
│   └── seed/               # Seed data
│
├── internal/
│   ├── db/                 # DB connection pool
│   ├── env/                # Environment variables
│   └── store/              # Data access layer
│       ├── storage.go      # Storage interface
│       ├── users.go        # Users repository
│       ├── posts.go        # Posts repository
│       ├── comments.go     # Comments repository
│       ├── followers.go    # Followers repository
│       └── pagination.go   # Feed query parsing
│
├── docs/                   # Generated Swagger docs
├── scripts/                # Utility scripts
├── devenv.nix              # Devenv config (env, services, tasks)
├── devenv.local.nix        # Local overrides (gitignored)
├── .air.toml               # Hot reload config
├── .envrc                  # direnv activation
├── endpoints.http          # API test requests (kulala-cli)
└── go.mod / go.sum
```

## Configuration

Environment variables are defined in `devenv.nix` with safe defaults. Override them in `devenv.local.nix` (gitignored).

| Variable | Default | Description |
| ---------- | --------- | ------------- |
| `ADDR` | `:8080` | Server address |
| `EXTERNAL_URL` | `localhost:8080` | Public URL (Swagger host) |
| `DB_ADDR` | `postgres://admin:changeme@127.0.0.1:5433/social?sslmode=disable` | DB connection |
| `DB_MAX_OPEN_CONNS` | `30` | Max open connections |
| `DB_MAX_IDLE_CONNS` | `30` | Max idle connections |
| `DB_MAX_IDLE_TIME` | `15m` | Max idle time |
| `ENV` | `development` | Environment |

## Tech Stack

| Layer | Technology |
| ------- | ----------- |
| Language | Go 1.26 |
| Router | Chi v5 |
| Database | PostgreSQL 17 |
| Migration | golang-migrate |
| Validator | go-playground/validator |
| Dev tools | devenv, air, kulala-cli, swag |

## License

MIT
