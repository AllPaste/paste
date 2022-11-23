package driver

import (
	"github.com/AllPaste/paste/config"
	"github.com/AllPaste/paste/internal/data/ent"
	"log"
)

func NewDriver(c config.Config) (*ent.Client, func(), error) {
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
