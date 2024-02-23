package main

import (
	"log"
	"p2p/cmd/app"
	"p2p/config"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("can't get config", err)
	}

	err = app.Run(cfg)
	if err != nil {
		log.Fatal("can't run app")
	}

}
