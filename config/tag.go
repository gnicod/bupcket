package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type TagConfig struct {
	Bucket string
}

func NewTagConfig(name string) (*TagConfig, error) {
	n := viper.GetStringMapString("tags." + name)
	if val, ok := n["bucket"]; ok {
		return &TagConfig{
			Bucket: val,
		}, nil
	}
	return &TagConfig{}, fmt.Errorf("Tag config without bucket")
}
