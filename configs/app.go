package configs

import "github.com/caarlos0/env/v10"

type App struct {
	Name    string `env:"APP_NAME,notEmpty" envDefault:"App" json:"name"`
	AppPort string `env:"APP_PORT" envDefault:"8000" json:"port"`
	Debug   string `env:"APP_DEBUG" envDefault:"true" json:"debug"`
}

func NewApp() (*App, error) {
	c := &App{}

	err := env.Parse(c)

	return c, err
}
