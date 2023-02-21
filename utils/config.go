package utils

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DbDriver      string        `mapstructure:"DB_DRIVER"`
	DbSource      string        `mapstructure:"DB_SOURCE"`
	MigrationUrl  string        `mapstructure:"MIGRATION_URL"`
	ServerAddress string        `mapstructure:"SERVER_ADDRESS"`
	SymmetricKey  string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	TokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.Unmarshal(&config)
	return
}
