package main

import (
	"context"
	"log"

	"github.com/Sudo-Victor-Victory/surgiledger/internal/api"
	"github.com/Sudo-Victor-Victory/surgiledger/internal/database"
)

func main() {
	ctx := context.Background()

	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := api.SetupRouter(db)

	log.Println("API running on http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
