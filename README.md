# Social API

![Go](https://img.shields.io/badge/Go-1.26.1-00ADD8?logo=go&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-Cloud-232F3E?logo=amazonaws&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Container-2496ED?logo=docker&logoColor=white)

## Overview
Social API is a RESTful backend for a social platform built with Go, Chi, and PostgreSQL.

## Requirements
- Go 1.26.1+
- Docker (Postgres)
- just (command runner)
- golang-migrate (migrations)
- air (hot reload, optional)

## Quick Start
```bash
just db-up
just migrate-up
just dev
```

## Commands
```bash
just help
just run
just dev
just build
just clean
just test
just vet
just check
just fmt-go
just fmt-go-imports
just fmt-sql
just db-up
just db-down
just db-logs
just db-seed
just migrate-up
just migrate-down
just migrate-create <name>
just tools
```

## Configuration
Environment variables:
- `ADDR` (default `:8081`)
- `DB_ADDR` (default `postgres://admin:adminpassword@localhost:5432/social?sslmode=disable`)
- `DB_MAX_OPEN_CONNS` (default `30`)
- `DB_MAX_IDLE_CONNS` (default `30`)
- `DB_MAX_IDLE_TIME` (default `15m`)
- `ENV` (default `development`)

## Project Structure
- `cmd/api/` — API entrypoint, router, handlers
- `cmd/migrate/migrations/` — SQL migrations
- `cmd/seed/` — seed command
- `internal/db/` — DB connection and seeding
- `internal/env/` — env handling
- `internal/store/` — data access layer
- `scripts/` — utility scripts
- `docker-compose.yml` — Postgres container

## API Base URL
`http://localhost:8081/v1`

## License
MIT. See `LICENSE`.
