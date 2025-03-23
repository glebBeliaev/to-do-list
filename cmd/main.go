package main

import (
	"github.com/glebbeliaev/to-do-list/internal/config"
	"github.com/glebbeliaev/to-do-list/internal/repository/memstorage"
	"github.com/glebbeliaev/to-do-list/internal/server"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	repo := memstorage.New()
	server := server.New(*cfg, repo)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
