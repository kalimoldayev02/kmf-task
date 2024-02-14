package main

import (
	"github.com/kalimoldayev02/kmf-task/app"
	"github.com/kalimoldayev02/kmf-task/pkg/config"
	_ "github.com/lib/pq"
)

func main() {
	config := config.GetInstance()

	app.Run(config)
}
