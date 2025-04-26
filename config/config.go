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
		DBUser:        "root",      // user MySQL của XAMPP
		DBPassword:    "",          // mật khẩu XAMPP, mặc định là rỗng
		DBHost:        "127.0.0.1", // localhost
		DBPort:        "3306",      // port MySQL
		DBName:        "localgems", // tên database bạn đã tạo
		ServerAddress: ":8080",     // port server Go
	}
}
