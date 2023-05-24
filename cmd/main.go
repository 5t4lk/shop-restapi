package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"online_shop/internal/db/mongo"
	"online_shop/internal/repository"
	"online_shop/internal/server"
	"online_shop/internal/services"
	v1 "online_shop/internal/transporthttp/v1"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error occurred while initiating config: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error occurred while loading env variables: %s", err)
	}

	db, err := mongo.NewDBConnection(os.Getenv("DB_CONNECT"))
	if err != nil {
		logrus.Fatalf("error occurred while creating a new connection of database: %s", err)
	}
	defer db.Close()

	repo := repository.New(db.Client, os.Getenv("DB_NAME"))
	service := services.New(repo)
	handler := v1.NewHandler(service)

	srv := server.NewServer(viper.GetString("port"), handler.Init())
	go func() {
		if err := srv.Run(); err != nil {
			if err == http.ErrServerClosed {
				return
			}
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Println("Online Shop Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Online Shop Shut Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred while shutting down the server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
