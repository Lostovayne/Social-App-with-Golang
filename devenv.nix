{ pkgs, ... }:

let
  # ── Database configuration ──────────────────────────────────────────
  # Safe defaults — override in devenv.local.nix (gitignored) for real creds
  # See devenv.local.nix.example for the expected override structure
  dbUser = "admin";
  dbPass = "changeme";
  dbName = "social";
  dbHost = "127.0.0.1";
  dbPort = 5433; # devenv offset — avoids conflicts with system postgres

  dbAddr = "postgres://${dbUser}:${dbPass}@${dbHost}:${toString dbPort}/${dbName}?sslmode=disable";

  migrationsPath = "./cmd/migrate/migrations";

  # ── Custom packages ─────────────────────────────────────────────────
  # go-migrate from nixpkgs fails due to snowflake driver; we build only postgres support.
  migrate = pkgs.writeShellApplication {
    name = "migrate";
    runtimeInputs = [ pkgs.go ];
    text = ''
      migrate_bin="$(go env GOPATH)/bin/migrate"
      if [ ! -x "$migrate_bin" ]; then
      echo "Installing go-migrate..."
      if ! go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.19.1; then
      echo "ERROR: Failed to install go-migrate. Check your Go installation and network."
      exit 1
      fi
      fi
      exec "$migrate_bin" "$@"
    '';
  };

  # kulala-cli: HTTP client for .http files (https://github.com/mistweaverco/kulala-cli)
  kulala = pkgs.writeShellApplication {
    name = "kulala";
    runtimeInputs = [ pkgs.bun ];
    text = ''
      kulala_dir="$HOME/.cache/kulala-cli"
      if [ ! -f "$kulala_dir/node_modules/.bin/kulala" ]; then
      echo "Installing kulala-cli..."
      mkdir -p "$kulala_dir"
      cd "$kulala_dir"
      if ! bun init -y 2>/dev/null; then
      echo "ERROR: Failed to initialize bun project in $kulala_dir"
      exit 1
      fi
      if ! bun add @mistweaverco/kulala-cli; then
      echo "ERROR: Failed to install kulala-cli. Check your network connection."
      exit 1
      fi
      fi
      exec "$kulala_dir/node_modules/.bin/kulala" "$@"
    '';
  };
in
{
  # ── Languages ───────────────────────────────────────────────────────
  # https://devenv.sh/languages/
  languages.go.enable = true;

  # ── Packages ────────────────────────────────────────────────────────
  # https://devenv.sh/packages/
  packages = with pkgs; [
    gnumake
    git
    bun
    migrate
    kulala
    golangci-lint
    go-tools # goimports, godoc, etc.
    pgformatter
    sqlfluff
    air
  ];

  # ── Environment variables ───────────────────────────────────────────
  # https://devenv.sh/basics/
  env = {
    ADDR = ":8080";
    ENV = "development";
    EXTERNAL_URL = "localhost:8080";
    DB_MAX_OPEN_CONNS = "30";
    DB_MAX_IDLE_CONNS = "30";
    DB_MAX_IDLE_TIME = "15m";
    DB_ADDR = dbAddr;
  };

  # ── Services ────────────────────────────────────────────────────────
  # https://devenv.sh/services/
  # Run with: devenv up (all) | devenv up postgres (just db)
  services.postgres = {
    enable = true;
    package = pkgs.postgresql_17;
    listen_addresses = dbHost;
    port = dbPort;
    initialDatabases = [
      {
        name = dbName;
        user = dbUser;
        pass = dbPass;
      }
    ];
  };

  # ── Processes ───────────────────────────────────────────────────────
  # https://devenv.sh/processes/
  # Run with: devenv up (starts postgres + dev)
  processes.dev.exec = "air -c .air.toml";

  # ── Tasks ───────────────────────────────────────────────────────────
  # https://devenv.sh/tasks/
  # Run with: devenv tasks run <task>
  tasks = {
    # ── Setup ─────────────────────────────────────────────────────────
    "app:setup".exec = ''
      migrate -path=${migrationsPath} --database="$DB_ADDR" up
      go run ./cmd/seed
    '';

    # ── Database ──────────────────────────────────────────────────────
    "app:migrate-up".exec = ''
      migrate -path=${migrationsPath} --database="$DB_ADDR" up
    '';

    "app:migrate-down".exec = ''
      migrate -path=${migrationsPath} --database="$DB_ADDR" down 1
    '';

    "app:migrate-new".exec = ''
      if [ -z "$1" ]; then
        echo "Usage: devenv tasks run app:migrate-new -- <name>"
        exit 1
      fi
      migrate create -seq -ext sql -dir ${migrationsPath} "$1"
    '';

    "app:seed".exec = "go run ./cmd/seed";

    # ── Run ───────────────────────────────────────────────────────────
    "app:run".exec = "go run ./cmd/api";

    "app:dev".exec = "air -c .air.toml";

    # ── Build ─────────────────────────────────────────────────────────
    "app:build-api".exec = ''
      mkdir -p bin
      go build -o ./bin/api ./cmd/api
    '';

    "app:build-seed".exec = ''
      mkdir -p bin
      go build -o ./bin/seed ./cmd/seed
    '';

    "app:build".exec = ''
      mkdir -p bin
      go build -o ./bin/api ./cmd/api
      go build -o ./bin/seed ./cmd/seed
    '';

    "app:clean".exec = "rm -f ./bin/api ./bin/seed ./bin/main";

    # ── Test & Lint ───────────────────────────────────────────────────
    "app:test".exec = "go test ./...";

    "app:vet".exec = "go vet ./...";

    "app:check".exec = ''
      go fmt ./cmd/... ./internal/...
      go vet ./...
      go test ./...
    '';

    # ── Format ────────────────────────────────────────────────────────
    "app:fmt".exec = "go fmt ./cmd/... ./internal/...";

    "app:fmt-imports".exec = ''
      find ./cmd ./internal -name '*.go' \
        -not -path './vendor/*' \
        -not -path './bin/*' \
        -not -path './tmp/*' \
        -exec goimports -w {} +
    '';

    "app:fmt-sql".exec = ''
      if command -v pg_format >/dev/null 2>&1; then
        find ./cmd ./scripts -name '*.sql' -exec pg_format -i {} +
      elif command -v sqlfluff >/dev/null 2>&1; then
        sqlfluff fix --dialect postgres ./cmd/migrate/migrations ./scripts
      else
        echo "No SQL formatter found. Install pg_format or sqlfluff."
        exit 1
      fi
    '';

    # ── Docs ──────────────────────────────────────────────────────────
    "app:docs".exec = ''
      swag init -g ./cmd/api/main.go -d cmd,internal && swag fmt
    '';

    # ── HTTP Client (kulala) ──────────────────────────────────────────
    "app:http".exec = "kulala run endpoints.http";

    "app:http-run".exec = ''
      if [ -z "$1" ]; then
        echo "Usage: devenv tasks run app:http-run -- <name>"
        exit 1
      fi
      kulala run endpoints.http --name "$1"
    '';
  };

  # ── Git hooks ───────────────────────────────────────────────────────
  # https://devenv.sh/git-hooks/
  git-hooks.hooks = {
    gofmt.enable = true;
    nixfmt.enable = true;
    shellcheck.enable = true;
  };

  # ── Shell entry ─────────────────────────────────────────────────────
  enterShell = ''
    echo ""
    echo "  \033[36m┌─────────────────────────────────────────────────┐\033[0m"
    echo "  \033[36m│\033[0m  Social API — devenv                          \033[36m│\033[0m"
    echo "  \033[36m└─────────────────────────────────────────────────┘\033[0m"
    echo ""
    echo "  Go:     $(go version | awk '{print $3}')"
    echo "  DB:     $DB_ADDR"
    echo ""
    echo "  \033[33mQuick start:\033[0m"
    echo "    devenv up                    Postgres + API (hot reload)"
    echo "    devenv up postgres           Database only"
    echo ""
    echo "  \033[33mTasks:\033[0m"
    echo "    devenv tasks run app:setup        Migrate + seed"
    echo "    devenv tasks run app:dev          Hot reload"
    echo "    devenv tasks run app:test         Run tests"
    echo "    devenv tasks run app:check        Fmt + vet + tests"
    echo "    devenv tasks run app:http          Run all endpoints"
    echo "    devenv tasks list                 All tasks"
    echo ""
  '';

  # ── Tests ───────────────────────────────────────────────────────────
  # https://devenv.sh/tests/
  # Run with: devenv test
  enterTest = ''
    go version
    go vet ./...
    go test ./...
  '';
}
