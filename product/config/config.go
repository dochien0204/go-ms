package config

import "github.com/spf13/viper"

type (
	PostgresConfig struct {
		Timeout  int
		DBname   string
		Username string
		Password string
		Host     string
		Port     string
	}
)

func LoadConfig() (config PostgresConfig, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
