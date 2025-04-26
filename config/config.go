package config

type Config struct {
	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	ServerAddress string
}

func NewConfig() *Config {
	return &Config{
		DBUser:        "root",
		DBPassword:    "",
		DBHost:        "127.0.0.1",
		DBPort:        "3306",
		DBName:        "localgems",
		ServerAddress: ":8080",
	}
}
