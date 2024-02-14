package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/kalimoldayev02/kmf-task/app"
	"github.com/kalimoldayev02/kmf-task/pkg/config"
)

func main() {
	config := config.LoadConfig()

	app.Run(config)
}
