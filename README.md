<!-- MARKDOWN ANCHOR TOC: START -->

<a name="table-of-contents"></a>

<!-- MARKDOWN ANCHOR TOC: END -->

<div align="center">

<a href="https://github.com/Elevate-Techworks/social">
<img src="https://capsule-render.vercel.app/api?type=waved&height=300&section=header&text=Social%20API&fontSize=60&fontAlignY=40&desc=Backend%20API%20for%20Social%20Platform&descAlignY=55&descSize=18&theme=transparent" width="100%" alt="Header"/>
</a>

<!-- BADGES -->

[![Go Version](https://img.shields.io/badge/Go_1.26.1-00ADD8?style=flat&logo=go&logoColor=white&labelColor=0b7d96)](https://go.dev/)
[![Chi Router](https://img.shields.io/badge/Chi_v5.2.5-FF6F61?style=flat&logo=lightning&labelColor=4a4a4a)](https://github.com/go-chi/chi)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL_17-336791?style=flat&logo=postgresql&logoColor=white&labelColor=2f4f4f)](https://www.postgresql.org/)
[![golang-migrate](https://img.shields.io/badge/golang--migrate-v4-4DB6AE?style=flat&logoColor=white&labelColor=0e5c4a)](https://github.com/golang-migrate/migrate)
[![air](https://img.shields.io/badge/air-v1.27.0-FF6F61?style=flat&logoColor=white&labelColor=d65d5d)](https://github.com/air-verse/air)
[![License](https://img.shields.io/badge/License-MIT-5e2d75?style=flat&labelColor=3d1d4a)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/Elevate-Techworks/social)](https://goreportcard.com/report/github.com/Elevate-Techworks/social)
[![Go Reference](https://pkg.go.dev/badge/github.com/Elevate-Techworks/social.svg)](https://pkg.go.dev/github.com/Elevate-Techworks/social)

<!-- TAGLINE -->

Modern RESTful API for social platforms, built with **Go**, **Chi router**, and **PostgreSQL**.

[Getting Started](#-getting-started) • [Documentation](#-documentation) • [API Reference](#-api-reference) • [Built With](#-built-with) • [Contributing](#-contributing)

</div>

---

## 📋 Table of Contents

- [📋 Table of Contents](#-table-of-contents)
- [💡 Overview](#-overview)
- [🚀 Getting Started](#-getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Quick Start](#quick-start)
- [🛠 Built With](#-built-with)
- [🏗 Architecture](#-architecture)
- [📂 Project Structure](#-project-structure)
- [⚙️ Configuration](#-configuration)
- [🐳 Docker](#-docker)
- [🗄 Database](#-database)
- [🔧 Commands Reference](#-commands-reference)
  - [Development](#development)
  - [Database](#database-1)
  - [Migrations](#migrations)
  - [Code Quality](#code-quality)
- [📡 API Reference](#-api-reference)
  - [Health Check](#health-check)
  - [Users](#users)
  - [Posts](#posts)
  - [Comments](#comments)
- [🧪 Testing](#-testing)
- [📜 Versioning](#-versioning)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)
- [🙏 Acknowledgments](#-acknowledgments)

---

## 💡 Overview

**Social API** is a high-performance, production-ready RESTful API designed to power social media platforms. Built with modern Go practices, it leverages the **Chi router** for blazing-fast HTTP routing and **PostgreSQL** for reliable data persistence.

### Key Features

- **RESTful Design** — Follows industry best practices with proper HTTP methods and status codes
- **Database Migrations** — Safe and reversible schema changes with golang-migrate
- **Hot Reloading** — Instant development feedback with air
- **Code Quality** — Built-in linting, formatting, and testing workflows
- **Environment-Based** — Configuration via environment variables (with `.envrc` support)

---

## 🚀 Getting Started

Follow these instructions to set up and run the project locally.

### Prerequisites

Ensure you have the following tools installed:

| Tool | Version | Purpose | Installation |
|------|---------|---------|--------------|
| <img src="https://simpleicons.org/icons/go.svg" width="16" height="16" /> **Go** | 1.26.1+ | Programming language | [go.dev/doc/install](https://go.dev/doc/install) |
| <img src="https://simpleicons.org/icons/docker.svg" width="16" height="16" /> **Docker** | Latest | Container runtime | [docs.docker.com/get-docker](https://docs.docker.com/get-docker/) |
| <img src="https://simpleicons.org/icons/gnubash.svg" width="16" height="16" /> **Just** | Latest | Command runner | [just.systems/man](https://just.systems/man/en/) |
| <img src="https://simpleicons.org/icons/postgresql.svg" width="16" height="16" /> **golang-migrate** | v4 | Database migrations | `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest` |
| <img src="https://simpleicons.org/icons/go.svg" width="16" height="16" /> **air** | Latest | Hot reload | `go install github.com/air-verse/air@latest` |

> **Tip:** Run `just tools` to automatically install `air` and `goimports`.

---

### Installation

```bash
# 1. Clone the repository
git clone https://github.com/Elevate-Techworks/social.git
cd social

# 2. Install development tools
just tools
```

---

### Quick Start

```bash
# 3. Start PostgreSQL with Docker
just db-up

# 4. Run database migrations
just migrate-up

# 5. Start development server (with hot reload)
just dev
```

The API will be available at `http://localhost:8081`

---

## 🛠 Built With

This project is built with the following technologies:

### Core

| Technology | Version | Description |
|-----------|---------|------------|
| <img src="https://img.icons8.com/color/48/go-logo.png" width="24" /> **Go** | 1.26.1 | Primary programming language |
| <img src="https://github.com/go-chi/chi/blob/master/_examples/logo.png" width="24"/> **Chi** | v5.2.5 | HTTP router |
| <img src="https://www.postgresql.org/media/img/about/press/elephant.png" width="24"/> **PostgreSQL** | 17 | Database |

### Tools

| Technology | Purpose |
|-----------|---------|
| **golang-migrate** | Database migration management |
| **air** | Live code reload for Go |
| **goimports** | Import organization |
| **pg_format** / **sqlfluff** | SQL formatting |
| **just** | Task automation |
| **docker-compose** | Container orchestration |

---

## 🏗 Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Social API Architecture               │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  ┌─────────────┐        ┌─────────────────────────────┐   │
│  │   Client    │───────▶│       Chi Router             │   │
│  │  (HTTP)     │        │    (api/v1/* handlers)      │   │
│  └─────────────┘        └──────────────┬────────────────┘   │
│                                      │                    │
│                                      ▼                    │
│                           ┌─────────────────────────────┐ │
│                           │      Business Logic          │ │
│                           │    (internal/store/*)       │ │
│                           └──────────────┬──────────────┘ │
│                                          │                │
│                                          ▼                │
│                           ┌─────────────────────────────┐ │
│                           │    PostgreSQL 17           │ │
│                           │    (persistent storage)    │ │
│                           └─────────────────────────────┘ │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## 📂 Project Structure

```
social/
├── cmd/                          # Executable applications
│   ├── api/                     # Main API application
│   │   ├── main.go              # Entry point
│   │   ├── api.go              # Router configuration
│   │   ├── errors.go           # Error handlers
│   │   ├── health.go          # Health endpoints
│   │   ├── json.go            # JSON utilities
│   │   └── posts.go           # Post handlers
│   │
│   ├── migrate/                # Database migrations
│   │   └── migrations/        # Migration SQL files
│   │       ├── 000001_create_users.up.sql
│   │       ├── 000001_create_users.down.sql
│   │       ├── 000002_posts_create.up.sql
│   │       ├── 000002_posts_create.down.sql
│   │       └── ...
│   │
│   ├── migrate/seed/           # Database seeding
│   │   └── main.go
│   │
│   └── seed/                  # Seed command
│       └── main.go
│
├── internal/                   # Private application code
│   ├── db/                    # Database connection
│   │   ├── db.go
│   │   └── seed.go
│   │
│   ├── env/                   # Environment variables
│   │   └── env.go
│   │
│   ├── handlers/              # HTTP handlers (placeholder)
│   ├── store/               # Data access layer
│   │   ├── users.go
│   │   ├── posts.go
│   │   ├── comments.go
│   │   └── storage.go
│   │   └── storage.go
│   │
│   └── ...
│
├── scripts/                    # Utility scripts
│   ├── db_init.sql
│   └── concurrency_test.go
│
├── bin/                       # Compiled binaries
│   └── main                  # Built binary
│
├── .air.toml                 # Air configuration
├── .envrc                    # Environment variables (direnv)
├── .gitignore
│
├── docker-compose.yml        # PostgreSQL container
├── endpoints.http           # HTTP test endpoints
├── go.mod                   # Go module definition
├── go.sum                   # Go checksums
├── justfile                  # Just commands
└── README.md               # This file
```

---

## ⚙️ Configuration

The application uses environment variables for configuration:

| Variable | Default | Description |
|----------|---------|-------------|
| `ADDR` | `:8081` | Server address (host:port) |
| `DB_ADDR` | `postgres://admin:adminpassword@localhost:5432/social?sslmode=disable` | Database connection string |
| `DB_MAX_OPEN_CONNS` | `30` | Maximum open database connections |
| `DB_MAX_IDLE_CONNS` | `30` | Maximum idle database connections |
| `DB_MAX_IDLE_TIME` | `15m` | Maximum idle connection time |
| `ENV` | `development` | Environment (development/production) |

> **Tip:** Use [direnv](https://direnv.net/) to automatically load `.envrc` variables.

---

## 🐳 Docker

### PostgreSQL Container

The project includes a pre-configured Docker Compose setup for PostgreSQL:

```yaml
# docker-compose.yml
services:
  db:
    image: postgres:17-alpine
    container_name: postgres_db
    environment:
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: adminpassword
      POSTGRES_DB: social
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"
```

### Managing the Database

```bash
# Start PostgreSQL
just db-up

# View logs
just db-logs

# Stop PostgreSQL
just db-down
```

---

## 🗄 Database

### Schema Overview

```sql
-- Users table
CREATE TABLE users (
    id          BIGSERIAL PRIMARY KEY,
    username    VARCHAR(255) UNIQUE NOT NULL,
    email       VARCHAR(255) UNIQUE NOT NULL,
    password    VARCHAR(255) NOT NULL,
    bio         TEXT,
    image_url   VARCHAR(500),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Posts table
CREATE TABLE posts (
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title       VARCHAR(255) NOT NULL,
    content     TEXT NOT NULL,
    tags        TEXT[],  -- Array of tags
    version     INTEGER DEFAULT 1,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Comments table
CREATE TABLE comments (
    id          BIGSERIAL PRIMARY KEY,
    post_id     BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    body        TEXT NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Migrations

Schema changes are managed with **golang-migrate**:

```bash
# Apply all pending migrations
just migrate-up

# Rollback the last migration
just migrate-down

# Create a new migration
just migrate-create add_comments_table
```

Migrations are located in `cmd/migrate/migrations/` following the naming convention:
- `000001_<name>.up.sql` — Apply migration
- `000001_<name>.down.sql` — Rollback migration

---

## 🔧 Commands Reference

All project commands are defined in the `justfile`. Run `just` or `just help` to see all available commands.

### Development

```bash
# Run with hot reload (recommended for development)
just dev

# Run without hot reload
just run

# Build the binary
just build

# Clean build artifacts
just clean
```

### Database

```bash
# Start PostgreSQL container
just db-up

# Stop PostgreSQL container
just db-down

# View PostgreSQL logs
just db-logs

# Seed the database with test data
just db-seed
```

### Migrations

```bash
# Apply all pending migrations
just migrate-up

# Rollback the last migration
just migrate-down

# Create a new migration
just migrate-create <migration_name>
```

### Code Quality

```bash
# Run full check (fmt + sqlfmt + vet + test)
just check

# Run tests
just test

# Run go vet
just vet

# Format Go code
just fmt-go

# Sort Go imports
just fmt-go-imports

# Format SQL files
just fmt-sql

# Install development tools (air, goimports)
just tools
```

---

## 📡 API Reference

> **Base URL:** `http://localhost:8081/v1`

### Health Check

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check endpoint |

**Example:**
```bash
curl -X GET http://localhost:8081/health
```

**Response:**
```json
{
  "status": "ok"
}
```

---

### Users

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/users` | List all users |
| `GET` | `/users/:id` | Get user by ID |
| `POST` | `/users` | Create new user |
| `PUT` | `/users/:id` | Update user |
| `DELETE` | `/users/:id` | Delete user |

---

### Posts

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/posts` | List all posts |
| `GET` | `/posts/:id` | Get post by ID |
| `POST` | `/posts` | Create new post |
| `PUT` | `/posts/:id` | Update post |
| `DELETE` | `/posts/:id` | Delete post |

---

### Comments

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/posts/:post_id/comments` | List comments for a post |
| `POST` | `/posts/:post_id/comments` | Create comment on a post |
| `DELETE` | `/comments/:id` | Delete comment |

---

> **Testing Tip:** Use the included `endpoints.http` file with your IDE's HTTP client (VS Code REST Client, IntelliJ HTTP Client) to test endpoints directly from your editor.

---

## 🧪 Testing

```bash
# Run all tests
just test

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

---

## 📜 Versioning

We use [SemVer](http://semver.org/) for versioning. For available versions, see the [tags on this repository](https://github.com/Elevate-Techworks/social/tags).

---

## ��� Contributing

Contributions are welcome! Please read our [contributing guidelines](CONTRIBUTING.md) first.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing-feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- [Chi Router](https://github.com/go-chi/chi) — Lightweight, idiomatic HTTP router
- [golang-migrate](https://github.com/golang-migrate/migrate) — Database migrations
- [air](https://github.com/air-verse/air) — Live reload for Go apps
- [just](https://just.systems) — Command runner
- [awesome-go](https://github.com/avelino/awesome-go) — Go resources

---

<div align="center">

**[↑ Back to Top](#table-of-contents)**

Built with ❤️ by **Elevate Techworks**

</div>