package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-playground/validator/v10"
	"github.com/mxngocqb/IoT-Project/config"
	"github.com/mxngocqb/IoT-Project/controller"
	"github.com/mxngocqb/IoT-Project/repository/driver"
	// _ "github.com/mxngocqb/IoT-Project/docs"
	driverservice "github.com/mxngocqb/IoT-Project/service/driver"
	zerolog "github.com/rs/zerolog/log"
)

var es *elasticsearch.Client

// @title IoT Proecjt
// @version 1.0
// @description This is a sample server for the IoT Project API.
// @host 172.18.53.136:8080
// @BasePath /
// @schemes http https

func main() {
	zerolog.Info().Msg("Server Started!")
	var err error

	cert, _ := os.ReadFile("E:/elasticsearch-8.12.2-windows-x86_64/elasticsearch-8.12.2/config/certs/http_ca.crt")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "d5nAui77yHF9D-6Vyj1t",
		CACert:   cert,
	}

	for {
		es, err = elasticsearch.NewClient(cfg)
		if err != nil {
			zerolog.Error().Err(err).Msg("Error creating the client")
		} else {
			break
		}
		res, _ := es.Info()

		fmt.Println(res)
	}

	db := config.ConnectDatabase()
	validate := validator.New()

	driverRepository := driver.NewDriverRepository(db)
	driverService := driverservice.NewDriverService(driverRepository, validate)
	driverController := controller.NewDriverController(driverService)

	controllers := &Controllers{
		DriverController: driverController,
	}

	routes := NewRouter(controllers)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	
	err = server.ListenAndServe()
	if err != nil {
		zerolog.Fatal().Err(err).Msg("Server Stopped")
	}
	
}
