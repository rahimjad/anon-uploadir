package main

import (
	"log"

	"./api"
	"./db"
)

func main() {
	// db.Rollback()
	_, err := db.Migrate()

	if err != nil {
		log.Panic(err)
	}

	api.Run()
}
