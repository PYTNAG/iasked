package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver           string `mapstructure:"DB_DRIVER"`
	DBSource           string `mapstructure:"DB_SOURCE"`
	PasetoSymmetricKey string `mapstructure:"PASETO_SYMMETRIC_KEY"`
}

func LoadConfig(path string) (cfg Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
