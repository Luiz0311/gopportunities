package main

import (
	"github.com/Luiz0311/gopportunities/config"
	"github.com/Luiz0311/gopportunities/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
