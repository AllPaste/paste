package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AllPaste/paste/config"
	"github.com/AllPaste/paste/pkg/log"
)

var stringFlag string

func init() {
	flag.StringVar(&stringFlag, "f", "./config/config.yaml", "config path")
}

func main() {
	flag.Parse()

	c := config.NewConfig(stringFlag)
	fmt.Printf("%#v\n", c)

	logger := log.New(os.Stderr, log.InfoLevel)
	defer log.Sync()

	app, cleanup, err := server(c, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	app.Run()
}
