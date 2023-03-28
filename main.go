package main

import (
	"cosmart/model"
	"cosmart/repository"
	"cosmart/routing/openlibrary"
	"cosmart/transport/https"
	"cosmart/usecase"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func init() {
	viper.SetConfigType(`json`)
	viper.AddConfigPath(`./`)
	viper.SetConfigName(`config`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Fatal error config file: %s\n", err)
	}
}

func main() {
	e := echo.New()

	schedules := []model.Schedule{}
	repo := repository.NewRepository(&schedules)

	client := http.Client{}

	usecaseDeps := usecase.Dependencies{
		Repository:         repo,
		OpenLibraryService: openlibrary.New(client),
		ContextTimeout:     time.Duration(viper.GetInt(`context.timeout`)) * time.Second,
	}

	usecase := usecase.NewUsecase(usecaseDeps)
	server := https.NewServer(e, usecase)

	server.Initialize()
	server.Run()
}
