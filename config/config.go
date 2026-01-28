package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Env struct {
	APP_HOST string `mapstructure:"APP_HOST"`
	APP_PORT string `mapstructure:"APP_PORT"`
	DB_URL   string `mapstructure:"DB_URL"`
}

func LoadConfig() *Env {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}

	viper.SetDefault("APP_HOST", "http://localhost")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("DB_URL", "postgresql://postgres:postgres@localhost:5432/kasir")

	var config Env
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
