package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config is the configuration for the application.
type Config struct {
	HTTPPort string `envconfig:"HTTP_PORT"`
	Secret   string `envconfig:"SECRET"`
}

// NewConfig return new instance of Config.
func NewConfig() (*Config, error) {
	c := Config{
		HTTPPort: "8080",
	}

	err := godotenv.Load()
	if err != nil {
		log.Println("unable to load .env file")
	}

	err = envconfig.Process("", &c)
	if err != nil {
		return nil, fmt.Errorf("unable to get envconfig %w", err)
	}

	return &c, nil
}
