package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/nkryuchkov/test-assignment-profitclicks/api"
	"github.com/nkryuchkov/test-assignment-profitclicks/config"
	"github.com/nkryuchkov/test-assignment-profitclicks/logger"
	"github.com/nkryuchkov/test-assignment-profitclicks/service"
	"github.com/nkryuchkov/test-assignment-profitclicks/storage"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	cfg, err := config.FromFile()
	if err != nil {
		log.Fatalf("Could not read config: %v. Exiting", err)
	}

	l := logger.New(cfg.Logger)
	defer func() {
		if err = l.Close(); err != nil {
			log.Printf("Could not close logger: %v", err)
		}
	}()

	database := storage.New(cfg.Storage, l)
	if err = database.Connect(); err != nil {
		l.Fatalf("Can't establish connection to database: %v", err)
		return
	}

	defer func() {
		if err = database.Close(); err != nil {
			l.Printf("Could not close database: %v", err)
		}
	}()

	s := service.New(l, database)

	apiServer := api.New(cfg.API, l, s)

	go func() {
		if err = apiServer.Start(); err != nil {
			l.Fatalf("Server error: %v", err)
		}
	}()

	<-quit
}
