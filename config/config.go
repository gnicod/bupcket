package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Tag struct{
	Bucket string `mapstructure:"bucket"`
}

type Config struct {
	Tags map[string]Tag `mapstructure:"tags"`
	Name string `string:"name"`
	Tag Tag `mapstructure:"tag"`
}

func (cfg *Config) GetTagConfig(tag string) (Tag, error) {
	if val, ok := cfg.Tags[tag]; ok {
		return val, nil
	}
	return Tag{}, fmt.Errorf("unknown config name")
}

func GetConfig() Config{
	var c Config
	viper.Unmarshal(&c)
	return c
}