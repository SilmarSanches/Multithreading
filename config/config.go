package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	URLViaCep    string `mapstructure:"URL_VIACEP"`
	URLBrasilAPI string `mapstructure:"URL_BRASILAPI"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(path + "/.env")
	viper.AutomaticEnv()

	_ = viper.BindEnv("URL_VIACEP")
	_ = viper.BindEnv("URL_BRASILAPI")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	_ = os.Setenv("URL_VIACEP", cfg.URLViaCep)
	_ = os.Setenv("URL_BRASILAPI", cfg.URLBrasilAPI)

	return &cfg, nil
}
