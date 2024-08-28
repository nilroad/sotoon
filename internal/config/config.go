package config

type Config struct {
	Debug       bool
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
	Port     int
	Username string
	Password string
	Database string
	Tz       string
}
