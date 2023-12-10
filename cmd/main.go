package main

import (
	"github.com/dmytrodemianchuk/bank-transactions/internal/config"
	"github.com/dmytrodemianchuk/bank-transactions/internal/repositories"
	"github.com/dmytrodemianchuk/bank-transactions/internal/services"
	"github.com/dmytrodemianchuk/bank-transactions/internal/transport"
	"github.com/dmytrodemianchuk/bank-transactions/pkg/database"
	"github.com/dmytrodemianchuk/bank-transactions/pkg/server"
	"github.com/dmytrodemianchuk/bank-transactions/pkg/signaler"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
)

// @title bank-transactions API
// @version 1.0

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}

func main() {
	cfg, err := config.Init()
	if err != nil {
		logrus.Fatalln(err)
	}

	db, err := database.NewPostgresConnection(cfg.Postgres)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer db.Close()

	repo := repositories.New(db)
	service := services.New(repo.RepoBank, repo.RepoRemote)
	handler := transport.NewHandler(service.ServicesBank, service.ServicesRemote)

	srv := server.New(cfg.HTTP.Port, handler.Init(cfg))

	go func() {
		if err = srv.Run(); err != nil {
			logrus.Errorf("error start server: %s", err.Error())
		}
	}()

	logrus.Println("rest-api started")

	signaler.Wait()

	if err = srv.Stop(); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
