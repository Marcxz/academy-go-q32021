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
	icsv := infraestructure.NewCsvInfraestructure(cfg)
	igeo := infraestructure.NewGeoInfraestructure(cfg)

	// repository
	rcsv := repository.NewCsvRepository(icsv)
	rgeo := repository.NewGeoRepository(igeo)

	// Usecase
	au := usecase.NewAddressUseCase(rcsv, rgeo)

	// controller
	ac := controller.NewAddressController(cfg, au)
	r := mux.NewRouter()

	ir := infraestructure.NewRouterInfraestructure(r, ac)
	ir.ConfigureHandlers()
	ir.Serve(cfg.Server)
}
