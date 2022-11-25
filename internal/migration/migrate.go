package main

import (
	"context"
	"log"

	"github.com/AllPaste/paste/config"
	"github.com/AllPaste/paste/internal/data/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	c := config.NewConfig("./config/config.yaml")

	client, err := ent.Open(c.DB.Driver, c.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	defer client.Close()
}
