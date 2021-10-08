package main

import (
	"log"

	"github.com/Marcxz/academy-go-q32021/conf"
	"github.com/Marcxz/academy-go-q32021/controller"
	"github.com/Marcxz/academy-go-q32021/infraestructure"
	"github.com/Marcxz/academy-go-q32021/repository"
	"github.com/Marcxz/academy-go-q32021/usecase"

	"github.com/gorilla/mux"
)

var cfg *conf.Config

// ConfigYML - func to load the config yml file and store it in a config struct
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
	icsv := infraestructure.NewCsvInfraestructure(cfg)
	igeo := infraestructure.NewGeoInfraestructure(cfg)
	idb := infraestructure.NewGeoDB(cfg)
	r := mux.NewRouter()
	ir := infraestructure.NewRouterInfraestructure(r)

	// repository
	rcsv := repository.NewCsvRepository(icsv)
	rgeo := repository.NewGeoRepository(igeo)
	rdb := repository.NewGeoDBRepository(idb)
	// Usecase
	au := usecase.NewAddressUseCase(rcsv, rgeo, rdb)

	// controller
	ac := controller.NewAddressController(cfg, au)

	// address Handlers
	ir.Get("/address", ac.ReadCSVAddress)
	ir.Get("/geocodeAddress", ac.GeocodeAddress)
	ir.Get("/storeGeocodeAddress", ac.StoreGeocodeAddress)
	ir.Get("/readAddressConcurrency", ac.ReadAddressCSVConcurrency)
	ir.Get("/generateRouterFrom2Address", ac.GenerateRouterFrom2Address)
	ir.Serve(cfg.Server)
}
