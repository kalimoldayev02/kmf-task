package main

import (
	"log"

	"github.com/kalimoldayev02/kmf-task/internal/app"
	"github.com/kalimoldayev02/kmf-task/pkg/config"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.NewCoifig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	// run
	app.Run(config)
}
