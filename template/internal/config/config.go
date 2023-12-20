package config

import (
	"aocli/template/internal/folder"
	"path"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

var (
	ConfigPath  = path.Join(folder.FindRoot(), "/template/.config")
	SecretsPath = path.Join(folder.FindRoot(), "/template/.config.secrets")
)

type Config struct {
	Secrets struct {
		AocSession string `env:"SESSION"`
	} `envPrefix:"AOC_SECRETS_"`

	Public struct {
		CurrentYear string `env:"YEAR"`
		BenchFlags  string `env:"BENCH_FLAGS"`
	} `envPrefix:"AOC_"`
}

var C Config = Config{}

func init() {
	godotenv.Load(SecretsPath)
	godotenv.Load(ConfigPath)

	if err := env.Parse(&C); err != nil {
		panic(err)
	}
}
