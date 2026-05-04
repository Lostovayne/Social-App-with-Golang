package main

import (
	"log"

	"github.com/Elevate-Techworks/social/internal/db"
	"github.com/Elevate-Techworks/social/internal/env"
	"github.com/Elevate-Techworks/social/internal/store"
)



func main() {
   addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable")

   conn, err := db.New(addr,3,3,"15m")
   if err != nil {
   log.Fatal(err)
   }

   defer conn.Close()

   store := store.NewStorage(conn)
   db.Seed(store)
}
