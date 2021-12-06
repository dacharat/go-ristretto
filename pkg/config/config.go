package config

import "github.com/kelseyhightower/envconfig"

var (
	Cfg Config
)

type Config struct {
	Redis Redis
}

type Redis struct {
	Host string `envconfig:"MY_ADDRESS" default:"redis:6379"`
}

func Process() {
	envconfig.MustProcess("", &Cfg)
}
