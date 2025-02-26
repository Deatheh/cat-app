package main

import (
	"log"
	"os"

	"github.com/Deatheh/cat-app"
	"github.com/Deatheh/cat-app/pkg/handler"
	"github.com/Deatheh/cat-app/pkg/repository"
	"github.com/Deatheh/cat-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Cat App API
// @version 1.0
// @description API Server for Cat-App API

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("errore initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("errore loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.ConfigDB{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initializing db: %s", err.Error())
	}

	minioDb, err := repository.MinioConnection(repository.ConfigMinio{
		Endpoint:        viper.GetString("minio.endpoint"),
		AccessKeyID:     viper.GetString("minio.accesskeyid"),
		SecretAccessKey: viper.GetString("secretaccesskey"),
	})
	if err != nil {
		logrus.Fatalf("failed to initializing minioDb: %s", err.Error())
	}
	repos := repository.NewRepository(db, minioDb)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(cat.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
