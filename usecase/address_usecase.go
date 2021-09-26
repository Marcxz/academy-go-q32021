package usecase

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Marcxz/academy-go-q32021/models"
	"github.com/Marcxz/academy-go-q32021/repository"
)

// Address - the interace for the address usecase
type Address interface {
	ReadCSVAddress(string) ([]models.Address, error)
	GeocodeAddress(string) (*models.Address, error)
	StoreGeoCodeAddress(string) (*models.Address, error)
}

var (
	cr = repository.NewCsvRepository()
	gr = repository.NewGeoRepository()
	as = make([]models.Address, 0)
)

type auc struct{}

// NewAddressUseCase - func to create a new address usecase used to link with the controller
func NewAddressUseCase() Address {
	return &auc{}
}

// validate - func to validate if an string address has the minimun requirements to be an address struc
func validate(i int, l string) error {
	al := strings.Split(l, "|")
	if len(al) != 4 {
		return fmt.Errorf("the line at the index %d should be composed for 4 pipes", i)
	}

	_, err := strconv.Atoi(al[0])
	if err != nil {
		return fmt.Errorf("the id at the index %d should be integer %s", i, al[0])
	}

	_, err = strconv.ParseFloat(al[2], 64)
	if err != nil {
		return fmt.Errorf("the lat column at the index %d should be float %s", i, al[2])
	}

	_, err = strconv.ParseFloat(strings.Replace(al[3], "\r", "", 1), 64)
	if err != nil {
		return fmt.Errorf("the lng column at the index %d should be float %s", i, al[3])
	}

	return nil
}

/*
func (*auc) geoAddress(a string) (*models.Address, error) {
	ad := models.Address{
		A: a,
		P: models.Point{},
	}
	return &ad, errors.New("Winno")
}
*/

// ReadCSVAddress - func to do the bussiness logic when you read all the address from a csv file
func (*auc) ReadCSVAddress(f string) ([]models.Address, error) {
	as = make([]models.Address, 0)

	cl, err := cr.ReadCSVFile(f)

	if err != nil {
		return nil, err
	}

	for i, l := range cl {
		err = validate(i, l)
		if err != nil {
			return nil, err
		}
		al := strings.Split(l, "|")
		id, _ := strconv.Atoi(al[0])
		an := al[1]
		lat, _ := strconv.ParseFloat(al[2], 64)
		lng, _ := strconv.ParseFloat(strings.Replace(al[3], "\r", "", 1), 64)
		a := models.Address{
			Id: id,
			A:  an,
			P: models.Point{
				Lat: lat,
				Lng: lng,
			},
		}
		as = append(as, a)
	}

	return as, nil
}

func (*auc) GeocodeAddress(a string) (*models.Address, error) {
	lat, lng, err := gr.GeoCodeAddress(a)

	if err != nil {
		return nil, err
	}

	if lat == -1 || lng == -1 {
		return nil, errors.New("the geocoding process can't be processed with the address specified")
	}

	sa, err := cr.ReadCSVFile(os.Getenv("fn"))

	if err != nil {
		return nil, err
	}

	ad := &models.Address{
		Id: len(sa),
		A:  a,
		P: models.Point{
			Lat: lat,
			Lng: lng,
		},
	}

	return ad, nil
}

func (*auc) StoreGeoCodeAddress(a string) (*models.Address, error) {
	lat, lng, err := gr.GeoCodeAddress(a)

	if err != nil {
		return nil, err
	}

	if lat == -1 || lng == -1 {
		return nil, errors.New("the geocoding process can't be processed with the address specified")
	}

	sa, err := cr.ReadCSVFile(os.Getenv("fn"))

	if err != nil {
		return nil, err
	}

	ad := &models.Address{
		Id: len(sa),
		A:  a,
		P: models.Point{
			Lat: lat,
			Lng: lng,
		},
	}

	err = cr.StoreAddressCSV(os.Getenv("fn"), ad.Id, ad.A, ad.P.Lat, ad.P.Lng)
	if err != nil {
		return nil, err
	}

	return ad, nil
}
