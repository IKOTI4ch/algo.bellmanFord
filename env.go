package main

import "github.com/caarlos0/env/v6"

type Config struct {
	PostgresURL   string `env:"PG_URL"`
	RedisGraphURL string `env:"REDIS_GRAPH_URL" envDefault:"redis://127.0.0.1:6379"`
}

func GetENV() (cfg Config) {
	cfg = Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return cfg
}
