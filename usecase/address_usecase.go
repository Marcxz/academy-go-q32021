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
	StoreGeocodeAddress(string) (*models.Address, error)
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
		id, err := strconv.Atoi(al[0])
		if err != nil {
			return nil, err
		}
		an := al[1]
		lat, err := strconv.ParseFloat(al[2], 64)
		if err != nil {
			return nil, err
		}
		lng, err := strconv.ParseFloat(strings.Replace(al[3], "\r", "", 1), 64)
		if err != nil {
			return nil, err
		}
		a := models.Address{
			ID: id,
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

//GeocodeAddress - Bussiness logic to validate if an address can be geocoded
func (*auc) GeocodeAddress(a string) (*models.Address, error) {
	lat, lng, err := gr.GeocodeAddress(a)

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
		ID: len(sa),
		A:  a,
		P: models.Point{
			Lat: lat,
			Lng: lng,
		},
	}

	return ad, nil
}

//StoreGeocodeAddress - Bussiness logic to validate if an address can be geocoded and stored
func (*auc) StoreGeocodeAddress(a string) (*models.Address, error) {
	lat, lng, err := gr.GeocodeAddress(a)

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
		ID: len(sa),
		A:  a,
		P: models.Point{
			Lat: lat,
			Lng: lng,
		},
	}

	err = cr.StoreAddressCSV(os.Getenv("fn"), ad.ID, ad.A, ad.P.Lat, ad.P.Lng)
	if err != nil {
		return nil, err
	}

	return ad, nil
}
