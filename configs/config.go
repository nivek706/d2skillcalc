package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	TxtDirPath string `mapstructure:"txtDirPath"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	viper.Unmarshal(&config)
	return config, err
}
