package main

import (
	config "github.com/Walter-Alipio/gopportunities/Config"
	"github.com/Walter-Alipio/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}

// 3:25:58
