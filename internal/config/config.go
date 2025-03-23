package config

import (
	"cmp"
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Host  string
	Port  int
	Debug bool
}

func ReadConfig() (*Config, error) {
	var cfg Config
	flag.StringVar(&cfg.Host, "host", "localhost", "server host")
	flag.IntVar(&cfg.Port, "port", 8080, "server port")
	flag.BoolVar(&cfg.Debug, "debug", false, "debug mode")
	flag.Parse()
	if cfg.Host == "localhost" {
		cfg.Host = cmp.Or(os.Getenv("HOST"), cfg.Host)
	}
	if cfg.Port == 8080 {
		devPort := strconv.Itoa(cfg.Port)
		envPort := cmp.Or(os.Getenv("PORT"), devPort)
		port, err := strconv.Atoi(envPort)
		if err != nil {
			return nil, err
		}
		cfg.Port = port
	}

	return &cfg, nil
}
