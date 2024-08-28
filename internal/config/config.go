package config

type Config struct {
	Debug       bool
	Tz          string
	LogLevel    string
	HTTPServer  HTTPServer
	MYSQLConfig MYSQLConfig
}

type HTTPServer struct {
	Host string
	Port int
}

type MYSQLConfig struct {
	Name     string
	Host     string
	Username string
	Password string
	Database string
	Port     int
}
