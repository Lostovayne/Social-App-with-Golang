.PHONY: help run dev build-api build-seed build clean test vet check fmt-go fmt-go-imports fmt-sql db-up db-down db-logs db-seed migrate-up migrate-down migrate-create gen-docs kulala kulala-list kulala-one tools

# kulala-cli: https://github.com/mistweaverco/kulala-cli
KULALA ?= $(shell command -v kulala 2>/dev/null || command -v kulala_cli 2>/dev/null)

# DB_ADDR viene de direnv (.envrc); fallback si no esta cargado
DB_ADDR ?= postgres://admin:adminpassword@localhost:5432/social?sslmode=disable
export DB_ADDR

.DEFAULT_GOAL := help

help: ## [Ayuda] Lista todos los comandos
	@grep -E '^[a-zA-Z0-9_-]+:.*##' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*## "}; {printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2}'

run: ## [Proyecto] Ejecuta la API sin hot reload
	go run ./cmd/api

dev: ## [Proyecto] Ejecuta la API con Air
	@if ! command -v air >/dev/null 2>&1; then echo "air no esta instalado. Ejecuta: go install github.com/air-verse/air@latest"; exit 1; fi
	air -c .air.toml

build-api: ## [Proyecto] Compila ./cmd/api en ./bin/api
	go build -o ./bin/api ./cmd/api

build-seed: ## [Proyecto] Compila ./cmd/seed en ./bin/seed
	go build -o ./bin/seed ./cmd/seed

build: build-api build-seed ## [Proyecto] Compila ambos binarios en ./bin/

clean: ## [Proyecto] Elimina binarios generados
	rm -f ./bin/api ./bin/seed ./bin/main

test: ## [Go] Corre tests
	go test ./...

vet: ## [Go] Corre go vet
	go vet ./...

check: fmt-go fmt-sql vet test ## [Go] Formato + SQL + vet + tests

fmt-go: ## [Go] Formatea Go por paquetes (cmd e internal)
	go fmt ./cmd/... ./internal/...

fmt-go-imports: ## [Go] Ordena imports con goimports
	@if ! command -v goimports >/dev/null 2>&1; then echo "goimports no esta instalado. Ejecuta: go install golang.org/x/tools/cmd/goimports@latest"; exit 1; fi
	find ./cmd ./internal -name '*.go' -not -path './vendor/*' -not -path './bin/*' -not -path './tmp/*' -exec goimports -w {} +

fmt-sql: ## [SQL] Formatea SQL con pg_format o sqlfluff
	@if command -v pg_format >/dev/null 2>&1; then find ./cmd ./scripts -name '*.sql' -exec pg_format -i {} +; \
	elif command -v sqlfluff >/dev/null 2>&1; then sqlfluff fix --dialect postgres ./cmd/migrate/migrations ./scripts; \
	else echo "No se encontro formateador SQL local. Instala pg_format o sqlfluff."; exit 1; fi

db-up: ## [DB] Levanta Postgres con Docker
	@if command -v docker >/dev/null 2>&1; then docker compose up -d db; else echo "No se encontro docker compose."; exit 1; fi

db-down: ## [DB] Baja contenedores
	@if command -v docker >/dev/null 2>&1; then docker compose down; else echo "No se encontro docker compose."; exit 1; fi

db-logs: ## [DB] Sigue logs de Postgres
	@if command -v docker >/dev/null 2>&1; then docker compose logs -f db; else echo "No se encontro docker compose."; exit 1; fi

db-seed: ## [DB] Ejecuta el seed de la base de datos
	go run ./cmd/seed

migrate-up: ## [Migraciones] Aplica migraciones
	@if ! command -v migrate >/dev/null 2>&1; then echo "migrate no esta instalado. Ejecuta: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; exit 1; fi
	migrate -path=./cmd/migrate/migrations --database="$(DB_ADDR)" up

migrate-down: ## [Migraciones] Revierte 1 migracion
	@if ! command -v migrate >/dev/null 2>&1; then echo "migrate no esta instalado. Ejecuta: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; exit 1; fi
	migrate -path=./cmd/migrate/migrations --database="$(DB_ADDR)" down

migrate-create: ## [Migraciones] Crear migracion (make migrate-create NAME=nombre)
	@if [ -z "$(NAME)" ]; then echo "Usage: make migrate-create NAME=nombre"; exit 1; fi
	@if ! command -v migrate >/dev/null 2>&1; then echo "migrate no esta instalado. Ejecuta: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; exit 1; fi
	migrate create -seq -ext sql -dir ./cmd/migrate/migrations "$(NAME)"

gen-docs: ## [Docs] Generacion de documentacion Swagger
	swag init -g ./cmd/api/main.go -d cmd,internal && swag fmt

kulala: ## [Endpoints Http] Ejecuta todos los requests (kulala-cli)
	@if [ -z "$(KULALA)" ]; then echo "Instala kulala-cli: https://github.com/mistweaverco/kulala-cli"; exit 1; fi
	$(KULALA) run endpoints.http

kulala-list: ## [Endpoints Http] Lista requests en endpoints.http
	@if [ -z "$(KULALA)" ]; then echo "Instala kulala-cli: https://github.com/mistweaverco/kulala-cli"; exit 1; fi
	$(KULALA) run endpoints.http --list

kulala-one: ## [Endpoints Http] Un request (make kulala-one NAME=health)
	@if [ -z "$(KULALA)" ]; then echo "Instala kulala-cli: https://github.com/mistweaverco/kulala-cli"; exit 1; fi
	@if [ -z "$(NAME)" ]; then echo "Usage: make kulala-one NAME=health"; exit 1; fi
	$(KULALA) run endpoints.http --name $(NAME)

tools: ## [Herramientas] Instala air y goimports
	go install github.com/air-verse/air@latest
	go install golang.org/x/tools/cmd/goimports@latest
