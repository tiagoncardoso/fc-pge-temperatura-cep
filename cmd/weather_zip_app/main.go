package main

import (
	"fmt"
	"github.com/tiagoncardoso/fc-pge-temperatura-cep/config"
	"github.com/tiagoncardoso/fc-pge-temperatura-cep/internal/application/usecase"
	"github.com/tiagoncardoso/fc-pge-temperatura-cep/internal/infra/web"
	"github.com/tiagoncardoso/fc-pge-temperatura-cep/internal/infra/web/webserver"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	findZip := usecase.NewRequestZipData(conf.ApiUrlZip)
	findWeather := usecase.NewRequestWeatherData(conf.ApiUrlWeather + "" + conf.ApiKeyWeather)

	webServer := webserver.NewWebServer(conf.WebServerPort)
	zipWeatherHandler := web.NewWeatherHandler(findZip, findWeather)

	webServer.AddHandler("/temperature/{cep}", "GET", zipWeatherHandler.GetWeatherByZip)

	fmt.Println("Starting web server on port", conf.WebServerPort)
	webServer.Start()
}
