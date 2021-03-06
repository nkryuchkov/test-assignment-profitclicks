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
	quit := make(chan os.Signal, 1)
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
		l.Errorf("Can't establish connection to database: %v", err)
		return
	}

	defer func() {
		if err = database.Close(); err != nil {
			l.Errorf("Could not close database: %v", err)
		}
	}()

	if err = database.CreateSchemaIfNotExists(); err != nil {
		l.Errorf("Could not create database schema: %v", err)
		return
	}

	s := service.New(l, database)

	apiServer := api.New(cfg.API, l, s)

	go func() {
		if err = apiServer.Start(); err != nil {
			l.Warnf("API serving finished: %v", err)
		}
	}()

	defer func() {
		if err = apiServer.Shutdown(); err != nil {
			l.Errorf("Could not shut down API server: %v", err)
		}
	}()

	<-quit
}
