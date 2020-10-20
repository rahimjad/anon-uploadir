package main

import (
	"log"

	"./db"
	"./routes"
)

func main() {
	// db.Rollback()
	_, err := db.Migrate()

	if err != nil {
		log.Panic(err)
	}

	routes.Run()
}
