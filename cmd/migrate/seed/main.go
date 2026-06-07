// NOTE: This seed binary duplicates functionality with cmd/seed/main.go.
// Consider consolidating into a single seed entry point.
package main

import (
	"log"

	"github.com/Elevate-Techworks/social/internal/db"
	"github.com/Elevate-Techworks/social/internal/env"
	"github.com/Elevate-Techworks/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable")

	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatalf("connecting to database: %v", err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	err = db.Seed(store)
	if err != nil {
		log.Fatalf("seeding database: %v", err)
	}
}
