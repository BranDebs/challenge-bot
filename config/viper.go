package config

import (
	"github.com/spf13/viper"
)

type config struct {
	v *viper.Viper
}

func New() Configer {
	v := viper.New()
	v.SetConfigFile("config.toml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		return nil
	}

	return &config{
		v: v,
	}
}

func (c config) UnmarshalKey(key string, val interface{}) error {
	return c.v.UnmarshalKey(key, val)
}

func (c config) GetString(key string) string {
	return c.v.GetString(key)
}
