package main

import (
	"log"

	"github.com/Marcxz/academy-go-q32021/conf"
	"github.com/Marcxz/academy-go-q32021/controller"
	"github.com/Marcxz/academy-go-q32021/infraestructure"
	"github.com/Marcxz/academy-go-q32021/repository"
	"github.com/Marcxz/academy-go-q32021/usecase"
)

var (
	cfg *conf.Config
)

func ConfigYML() {
	cfgPath, err := conf.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err = conf.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	ConfigYML()

	// infraestructure
	ir := infraestructure.NewMuxRouter()
	icsv := infraestructure.NewCsvInfraestructure(cfg)
	igeo := infraestructure.NewGeoInfraestructure(cfg)

	// repository
	rcsv := repository.NewCsvRepository(icsv)
	rgeo := repository.NewGeoRepository(igeo)

	// Usecase
	au := usecase.NewAddressUseCase(rcsv, rgeo)

	// controller
	ac := controller.NewAddressController(cfg, au)

	ir.Get("/address", ac.ReadCSVAddress)
	ir.Get("/geocodeAddress", ac.GeocodeAddress)
	ir.Get("/storeGeocodeAddress", ac.StoreGeocodeAddress)

	ir.Serve(cfg.Server)
}
