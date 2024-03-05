package main

import (
	"fmt"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	zerolog "github.com/rs/zerolog/log"
)

var es *elasticsearch.Client

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
}
