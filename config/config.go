package config

import (
	"github.com/joho/godotenv"
)

var configs map[string]string

func Config(key string) string {
	return configs[key]
}

// Config func to get env value
func LoadConfig(filename string) {
	// load .env file
	c, err := godotenv.Read(filename)
	if err != nil {
		panic("Error loading config file")
	}
	configs = c
}
