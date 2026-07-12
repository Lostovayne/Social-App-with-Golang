set dotenv-load := true

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
		@if ! command -v air &> /dev/null; then echo "air no esta instalado. Ejecuta: go install github.com/air-verse/air@latest"; exit 1; fi
		air -c .air.toml

# [Proyecto] Compila ./cmd/api en ./bin/api
build-api:
		go build -o ./bin/api ./cmd/api

# [Proyecto] Compila ./cmd/seed en ./bin/seed
build-seed:
		go build -o ./bin/seed ./cmd/seed

# [Proyecto] Compila ambos binarios en ./bin/
build: build-api build-seed

# [Proyecto] Elimina binarios generados
clean:
		rm -f ./bin/api ./bin/seed ./bin/main

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
		@if ! command -v goimports &> /dev/null; then echo "goimports no esta instalado. Ejecuta: go install golang.org/x/tools/cmd/goimports@latest"; exit 1; fi
		find ./cmd ./internal -name '*.go' -not -path './vendor/*' -not -path './bin/*' -not -path './tmp/*' -exec goimports -w {} +

# [SQL] Formatea SQL con pg_format o sqlfluff (instalados localmente)
fmt-sql:
		@if command -v pg_format &> /dev/null; then find ./cmd ./scripts -name '*.sql' -exec pg_format -i {} +; elif command -v sqlfluff &> /dev/null; then sqlfluff fix --dialect postgres ./cmd/migrate/migrations ./scripts; else echo "No se encontro formateador SQL local. Instala pg_format o sqlfluff."; exit 1; fi

# [DB] Levanta Postgres con Docker
db-up:
		@if command -v docker &> /dev/null; then docker compose up -d db; else echo "No se encontro docker compose."; exit 1; fi

# [DB] Baja contenedores
db-down:
		@if command -v docker &> /dev/null; then docker compose down; else echo "No se encontro docker compose."; exit 1; fi

# [DB] Sigue logs de Postgres
db-logs:
		@if command -v docker &> /dev/null; then docker compose logs -f db; else echo "No se encontro docker compose."; exit 1; fi

# [DB] Ejecuta el seed de la base de datos
db-seed:
		go run ./cmd/seed

# [Migraciones] Aplica migraciones
migrate-up:
		@if ! command -v migrate &> /dev/null; then echo "migrate no esta instalado. Ejecuta: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; exit 1; fi
		migrate -path=./cmd/migrate/migrations --database="${DB_ADDR:-postgres://admin:adminpassword@localhost:5432/social?sslmode=disable}" up

# [Migraciones] Revierte 1 migracion
migrate-down:
		@if ! command -v migrate &> /dev/null; then echo "migrate no esta instalado. Ejecuta: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; exit 1; fi
		migrate -path=./cmd/migrate/migrations --database="${DB_ADDR:-postgres://admin:adminpassword@localhost:5432/social?sslmode=disable}" down

# [Migraciones] Crea migracion: just migrate-create <nombre>
migrate-create migration_name:
		@if ! command -v migrate &> /dev/null; then echo "migrate no esta instalado. Ejecuta: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; exit 1; fi
		migrate create -seq -ext sql -dir ./cmd/migrate/migrations "{{migration_name}}"

# [Docs] Generacion de Documentacion
gen-docs:
	swag init -g ./api/main.go -d cmd,internal && swag fmt


# [Endpoints Http] Corre Kulala Http
kulala:
	kulala run endpoints.http

# [Tests Endpoints] Corre los tests de endpoints
#kulala-tests:
# kulala run --tests --env=production file.http

# [Herramientas] Instala air y goimports
tools:
		go install github.com/air-verse/air@latest
		go install golang.org/x/tools/cmd/goimports@latest
