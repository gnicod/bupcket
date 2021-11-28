package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Tag struct{
	Bucket string `mapstructure:"bucket"`
	MimeTypes []string `mapstructure:"mimetypes"`
}

func (t *Tag) AcceptMimeType(mimeType string) bool {
	if len(t.MimeTypes) == 0 {
		return true
	}
	for _, a := range t.MimeTypes {
        if a == mimeType {
            return true
        }
    }
    return false
}

type Config struct {
	Tags map[string]Tag `mapstructure:"tags"`
	Name string `string:"name"`
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