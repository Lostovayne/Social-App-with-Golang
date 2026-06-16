set dotenv-load := true
set shell := ["powershell", "-Command"]

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
		@if (-not (Get-Command air -ErrorAction SilentlyContinue)) { echo "air no esta instalado. Ejecuta: go install github.com/air-verse/air@latest"; exit 1 }
		air -c .air.toml

# [Proyecto] Compila ./cmd/api en ./bin/main
build:
		go build -o ./bin/main.exe ./cmd/api

# [Proyecto] Elimina binarios generados
clean:
		Remove-Item -Recurse -Force ./bin/main.exe -ErrorAction SilentlyContinue

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
		@if (-not (Get-Command goimports -ErrorAction SilentlyContinue)) { echo "goimports no esta instalado. Ejecuta: go install golang.org/x/tools/cmd/goimports@latest"; exit 1 }
		Get-ChildItem -Recurse -Filter *.go -Exclude vendor,bin,tmp | ForEach-Object { goimports -w $_.FullName }

# [SQL] Formatea SQL con pg_format o sqlfluff (instalados localmente)
fmt-sql:
		@if (Get-Command pg_format -ErrorAction SilentlyContinue) { Get-ChildItem -Recurse -Path ./cmd,./scripts -Filter *.sql | ForEach-Object { pg_format -i $_.FullName } }
		elseif (Get-Command sqlfluff -ErrorAction SilentlyContinue) { sqlfluff fix --dialect postgres ./cmd/migrate/migrations ./scripts }
		else { echo "No se encontro formateador SQL local. Instala pg_format o sqlfluff."; exit 1 }

# [DB] Levanta Postgres con Docker
db-up:
		@if (Get-Command docker -ErrorAction SilentlyContinue) { docker compose up -d db }
		else { echo "No se encontro docker compose."; exit 1 }

# [DB] Baja contenedores
db-down:
		@if (Get-Command docker -ErrorAction SilentlyContinue) { docker compose down }
		else { echo "No se encontro docker compose."; exit 1 }

# [DB] Sigue logs de Postgres
db-logs:
		@if (Get-Command docker -ErrorAction SilentlyContinue) { docker compose logs -f db }
		else { echo "No se encontro docker compose."; exit 1 }

# [DB] Ejecuta el seed de la base de datos
db-seed:
		go run ./cmd/seed

# [Migraciones] Aplica migraciones
migrate-up:
		@if (-not (Get-Command migrate -ErrorAction SilentlyContinue)) { echo "migrate no esta instalado. Ejecuta: scoop install migrate"; exit 1 }
		$dbAddr = if ($env:DB_ADDR) { $env:DB_ADDR } else { "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable" }
		migrate -path=./cmd/migrate/migrations --database="$dbAddr" up

# [Migraciones] Revierte 1 migracion
migrate-down:
		@if (-not (Get-Command migrate -ErrorAction SilentlyContinue)) { echo "migrate no esta instalado. Ejecuta: scoop install migrate"; exit 1 }
		$dbAddr = if ($env:DB_ADDR) { $env:DB_ADDR } else { "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable" }
		migrate -path=./cmd/migrate/migrations --database="$dbAddr" down

# [Migraciones] Crea migracion: just migrate-create <nombre>
migrate-create migration_name:
		@if (-not (Get-Command migrate -ErrorAction SilentlyContinue)) { echo "migrate no esta instalado. Ejecuta: scoop install migrate"; exit 1 }
		migrate create -seq -ext sql -dir ./cmd/migrate/migrations "{{migration_name}}"

# [Herramientas] Instala air y goimports
tools:
		go install github.com/air-verse/air@latest
		go install golang.org/x/tools/cmd/goimports@latest
