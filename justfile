set dotenv-load := true
set shell := ["bash", "-euo", "pipefail", "-c"]

default:
		@just --list

# [Ayuda] Lista todos los comandos con su categoria
help:
		@just --list

# [Proyecto] Ejecuta la API sin hot reload
run:
		go run ./cmd/api

# [Proyecto] Ejecuta la API con Air
dev:
		@if ! command -v air >/dev/null 2>&1; then \
			echo "air no esta instalado. Ejecuta: go install github.com/air-verse/air@latest"; \
			exit 1; \
		fi
		air -c .air.toml

# [Proyecto] Compila ./cmd/api en ./bin/main
build:
		go build -o ./bin/main ./cmd/api

# [Proyecto] Elimina binarios generados
clean:
		rm -rf ./bin/main

# [Go] Corre tests
test:
		go test ./...

# [Go] Corre go vet
vet:
		go vet ./...

# [Go] Formato + SQL + vet + tests
check: fmt-go fmt-sql vet test

# [Go] Formatea Go por paquetes (cmd e internal)
fmt-go:
		go fmt ./cmd/... ./internal/...

# [Go] Ordena imports con goimports
fmt-go-imports:
		@if ! command -v goimports >/dev/null 2>&1; then \
			echo "goimports no esta instalado. Ejecuta: go install golang.org/x/tools/cmd/goimports@latest"; \
			exit 1; \
		fi
		@find . -type f -name '*.go' \
			-not -path './vendor/*' \
			-not -path './bin/*' \
			-not -path './tmp/*' \
			-print0 | xargs -0 goimports -w

# [SQL] Formatea SQL con pg_format o sqlfluff (instalados localmente)
fmt-sql:
		@if command -v pg_format >/dev/null 2>&1; then \
			while IFS= read -r -d '' file; do pg_format -i "$file"; done < <(find ./cmd ./scripts -type f -name '*.sql' -print0); \
		elif command -v sqlfluff >/dev/null 2>&1; then \
			sqlfluff fix --dialect postgres ./cmd/migrate/migrations ./scripts; \
		else \
			echo "No se encontro formateador SQL local. Instala pg_format o sqlfluff."; \
			exit 1; \
		fi

# [DB] Levanta Postgres con Docker
db-up:
		@if docker compose version >/dev/null 2>&1; then \
			docker compose up -d db; \
		elif command -v docker-compose >/dev/null 2>&1; then \
			docker-compose up -d db; \
		else \
			echo "No se encontro docker compose."; \
			exit 1; \
		fi

# [DB] Baja contenedores
db-down:
		@if docker compose version >/dev/null 2>&1; then \
			docker compose down; \
		elif command -v docker-compose >/dev/null 2>&1; then \
			docker-compose down; \
		else \
			echo "No se encontro docker compose."; \
			exit 1; \
		fi

# [DB] Sigue logs de Postgres
db-logs:
		@if docker compose version >/dev/null 2>&1; then \
			docker compose logs -f db; \
		elif command -v docker-compose >/dev/null 2>&1; then \
			docker-compose logs -f db; \
		else \
			echo "No se encontro docker compose."; \
			exit 1; \
		fi

# [Migraciones] Aplica migraciones
migrate-up:
		@if ! command -v migrate >/dev/null 2>&1; then \
			echo "migrate no esta instalado. Ejecuta: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
			exit 1; \
		fi
		migrate -path=./cmd/migrate/migrations --database="${DB_ADDR:-postgres://admin:adminpassword@localhost:5432/social?sslmode=disable}" up

# [Migraciones] Revierte 1 migracion
migrate-down:
		@if ! command -v migrate >/dev/null 2>&1; then \
			echo "migrate no esta instalado. Ejecuta: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
			exit 1; \
		fi
		migrate -path=./cmd/migrate/migrations --database="${DB_ADDR:-postgres://admin:adminpassword@localhost:5432/social?sslmode=disable}" down 

# [Migraciones] Crea migracion: just migrate-create <nombre>
migrate-create migration_name:
		@if ! command -v migrate >/dev/null 2>&1; then \
			echo "migrate no esta instalado. Ejecuta: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
			exit 1; \
		fi
		migrate create -seq -ext sql -dir ./cmd/migrate/migrations "{{migration_name}}"

# [Herramientas] Instala air y goimports
tools:
		go install github.com/air-verse/air@latest
		go install golang.org/x/tools/cmd/goimports@latest
