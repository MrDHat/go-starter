package config

import "github.com/ilyakaznacheev/cleanenv"

type config struct {
	ServiceName    string `env:"SERVICE_NAME" env-default:"go-starter"`
	ServiceVersion string `env:"SERVICE_VERSION" env-default:"0.0.1"`
	Env            string `env:"ENV" env-default:"dev"`
	APIPort        int    `env:"API_PORT" env-default:"3000"`
}

var Cfg config

func IsDevEnv() bool {
	return Cfg.Env == "dev"
}

func IsProdEnv() bool {
	return Cfg.Env == "prod" || Cfg.Env == "production" || Cfg.Env == "release"
}

func init() {
	err := cleanenv.ReadEnv(&Cfg)
	if err != nil {
		panic(err)
	}
}
