package main

import (
	"log"

	"github.com/Elevate-Techworks/social/internal/db"
	"github.com/Elevate-Techworks/social/internal/env"
	"github.com/Elevate-Techworks/social/internal/store"
)

const version = "0.0.1"

func main() {

	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	// -- DB  Connection --
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Fatalf("database: %v", err)
	}

	defer db.Close()
	log.Println("Database connection pool established")

	store := store.NewStorage(db)

	app, err := newApplication(cfg, store)
	if err != nil {
		log.Fatalf("could not create application: %v", err)
	}

	mux := app.mount()
	if err := app.run(mux); err != nil {
		log.Fatal(err)
	}
}
