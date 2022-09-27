package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App      App
	Database Database
	Kafka    Kafka
}

func Setup() (config Config, err error) {
	if err = envconfig.Process("", &config); err != nil {
		return
	}

	return
}
