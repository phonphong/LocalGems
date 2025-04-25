package config

import "os"

type Config struct {
	DBPath        string
	ServerAddress string
}

func NewConfig() *Config {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./coffee.db"
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = ":8080"
	}

	return &Config{
		DBPath:        dbPath,
		ServerAddress: serverAddress,
	}
}
