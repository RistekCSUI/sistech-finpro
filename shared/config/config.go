package config

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	ServerPort      string `mapstructure:"SERVER_PORT"`
	MongoUrl        string `mapstructure:"MONGO_URL"`
	MongoDatabase   string `mapstructure:"MONGO_DATABASE"`
	MaxRequestLimit int    `mapstructure:"MAX_REQUEST_LIMIT"`
	JWTSecret       string `mapstructure:"JWT_SECRET"`
	RedisHost       string `mapstructure:"REDIS_HOST"`
	RedisPort       string `mapstructure:"REDIS_PORT"`
}

func NewEnvConfig() (*EnvConfig, error) {
	var envConf EnvConfig
	viper.AddConfigPath("../../")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&envConf)
	if err != nil {
		return nil, err
	}

	return &envConf, nil
}
