package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	mongodb "online_shop/internal/database/mongo"
	"online_shop/internal/handler"
	"online_shop/internal/repository"
	"online_shop/internal/server"
	service "online_shop/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error occurred while initiating config: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error occurred while loading env variables: %s", err)
	}

	db, err := mongodb.NewDBConnection(os.Getenv("DB_CONNECT"))
	if err != nil {
		logrus.Fatalf("error occurred while creating a new connection of database: %s", err)
	}
	defer db.Close()

	repos := repository.NewRepository(db.Client, os.Getenv("DB_NAME"))
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := server.NewServer(viper.GetString("port"), handlers.InitRoutes())
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
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
