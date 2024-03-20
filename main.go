package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-playground/validator/v10"
	"github.com/mxngocqb/IoT-Project/cache"
	"github.com/mxngocqb/IoT-Project/config"
	"github.com/mxngocqb/IoT-Project/controller"
	"github.com/mxngocqb/IoT-Project/repository/driver"
	"github.com/mxngocqb/IoT-Project/repository/nav_record"
	vehicle "github.com/mxngocqb/IoT-Project/repository/vehicle"
	driverservice "github.com/mxngocqb/IoT-Project/service/driver"
	vehicleservice "github.com/mxngocqb/IoT-Project/service/vehicle"
	zerolog "github.com/rs/zerolog/log"
)

var es *elasticsearch.Client

// @title IoT Proecjt
// @version 1.0
// @description This is a sample server for the IoT Project API.
// @host 192.168.88.132:9090
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
	redisConfig := config.NewRedisConfig()
	redisClient, err := config.ConnectRedis(redisConfig)
	if err != nil {
		panic(err)
	}

	validate := validator.New()

	driverRepository := driver.NewDriverRepository(db)
	driverService := driverservice.NewDriverService(driverRepository, validate)
	driverCache := cache.NewDriverRedisCache(redisClient)
	driverController := controller.NewDriverController(driverService, driverCache)

	vehicleRepository := vehicle.NewVehicleRepository(db)
	vehicleService := vehicleservice.NewVehicleService(vehicleRepository, validate)
	vehicleCache := cache.NewVehicleRedisCache(redisClient)
	vechileController := controller.NewVehicleController(vehicleService, vehicleCache)

	nav_record.NewNavRecordRepository(db)

	controllers := &Controllers{
		DriverController:  driverController,
		VehicleController: vechileController,
	}

	routes := NewRouter(controllers)

	server := &http.Server{
		Addr:    ":9090",
		Handler: routes,
	}

	err = server.ListenAndServe()
	if err != nil {
		zerolog.Fatal().Err(err).Msg("Server Stopped")
	}

	
	defer redisClient.Close()

}
