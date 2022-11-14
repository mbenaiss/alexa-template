package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}

	server := gin.Default()
	server.Any("/", dispatchIntent())

	err = server.Run(":" + cfg.HTTPPort)
	if err != nil {
		panic(fmt.Errorf("unable to start server: %w", err))
	}
}
