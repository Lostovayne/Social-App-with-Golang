package main

import (
	"log"

	"github.com/Elevate-Techworks/social/internal/db"
	"github.com/Elevate-Techworks/social/internal/env"
	"github.com/Elevate-Techworks/social/internal/store"
)

func main() {
	dbAddr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable")

	conn, err := db.New(dbAddr, 5, 5, "15m")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close()

	log.Println("Database connection pool established")

	// Truncate tables so the seed is idempotent.
	// Order matters: children before parents.
	tables := []string{"comments", "followers", "posts", "users"}
	for _, t := range tables {
		if _, err := conn.Exec("TRUNCATE TABLE " + t + " RESTART IDENTITY CASCADE"); err != nil {
			log.Fatalf("Failed to truncate %s: %v", t, err)
		}
	}
	log.Println("Tables truncated")

	storage := store.NewStorage(conn)

	if err := db.Seed(storage); err != nil {
		log.Fatalf("Seed failed: %v", err)
	}

	log.Println("Seed completed successfully")
}
