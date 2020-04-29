package utils

import (
	"github.com/jinzhu/configor"
)

var (
	Config = struct {
		DB struct {
			Host string `required:"true"`
		}
	}{}
)

func LoadConfig() {
	configor.Load(&Config, "config.yml")
}
