package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.AllowEmptyEnv(true)

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, fmt.Errorf("reading config: %w", err)
		}
	}

	return &Config{
		Debug:    loadBool("DEBUG"),
		LogLevel: loadString("LOG_LEVEL"),
		HTTPServer: HTTPServer{
			Host: loadString("HTTP_HOST"),
			Port: loadInt("HTTP_PORT"),
		},
		MYSQLConfig: MYSQLConfig{
			Name:     loadString("MYSQL_DATABASE_NAME"),
			Host:     loadString("MYSQL_DATABASE_HOST"),
			Port:     loadInt("MYSQL_DATABASE_PORT"),
			Username: loadString("MYSQL_DATABASE_USERNAME"),
			Password: loadString("MYSQL_DATABASE_PASSWORD"),
			Database: loadString("MYSQL_DATABASE_NAME"),
			Tz:       loadString("TZ"),
		},
	}, nil
}
