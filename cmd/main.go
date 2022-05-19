package main

import (
	"context"
	"log"

	"github.com/ElinaUrbanovich/menu-app-items/pkg/items"
	"github.com/jackc/pgx/v4"
)

func main() {
	database_url := "postgres://postgres:myownsummer12@localhost:5432/menu"
	var itemsServer *items.ItemServiceServer = items.NewCategoryServer()
	conn, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	defer conn.Close(context.Background())
	itemsServer.Conn = conn
	if err := itemsServer.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
