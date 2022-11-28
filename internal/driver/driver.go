package driver

import (
	"log"

	"github.com/AllPaste/paste/config"
	"github.com/AllPaste/paste/internal/data/ent"
	"github.com/google/wire"

	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet is driver providers.
var ProviderSet = wire.NewSet(NewDriver)

func NewDriver(c *config.Config) (*ent.Client, func(), error) {
	client, err := ent.Open(c.DB.Driver, c.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	return client, func() {
		if err := client.Close(); err != nil {
			log.Fatal(err)
		}
	}, nil
}
