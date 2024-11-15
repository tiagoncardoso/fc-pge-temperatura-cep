package config

import "github.com/spf13/viper"

type Conf struct {
	MyVariable string `mapstructure:"MY_VARIABLE"`
}

func LoadConfig() (*Conf, error) {
	viper.SetConfigName("app_config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var conf Conf
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
