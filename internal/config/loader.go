package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func loadString(envName string) string {
	validate(envName)

	return viper.GetString(envName)
}

func loadInt(envName string) int {
	validate(envName)

	return viper.GetInt(envName)
}

func loadBool(envName string) bool {
	validate(envName)

	return viper.GetBool(envName)
}
func validate(envName string) {
	exists := viper.IsSet(envName)
	if !exists {
		panic(fmt.Sprintf("environment variable [%s] does not exist", envName))
	}
}
