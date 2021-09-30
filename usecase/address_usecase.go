package usecase

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Marcxz/academy-go-q32021/models"
	"github.com/Marcxz/academy-go-q32021/repository"
)

type Reader interface {
	ReadCSVAddress() ([]models.Address, error)
}
type Geolocater interface {
	GeocodeAddress(string) (*models.Address, error)
}

type Storer interface {
	StoreGeocodeAddress(string) (*models.Address, error)
}

// Address - the interace for the address usecase
type Address interface {
	Reader
	Geolocater
	Storer
}

type auc struct {
	Cr repository.Csv
	Gr repository.Geo
}

// NewAddressUseCase - func to create a new address usecase used to link with the controller
func NewAddressUseCase(rcsv repository.Csv, rgeo repository.Geo) Address {
	return &auc{
		rcsv,
		rgeo,
	}
}

// ReadCSVAddress - func to do the bussiness logic when you read all the address from a csv file
func (a *auc) ReadCSVAddress() ([]models.Address, error) {
	as := make([]models.Address, 0)

	cl, err := a.Cr.ReadCSVFile()

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
func (a *auc) GeocodeAddress(add string) (*models.Address, error) {
	ad, err := a.createGeocodeAddress(add)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

//StoreGeocodeAddress - Bussiness logic to validate if an address can be geocoded and stored
func (a *auc) StoreGeocodeAddress(add string) (*models.Address, error) {

	ad, _ := a.createGeocodeAddress(add)
	err := a.Cr.StoreAddressCSV(ad.ID, ad.A, ad.P.Lat, ad.P.Lng)
	if err != nil {
		return nil, err
	}

	return ad, nil
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

func (a *auc) createGeocodeAddress(add string) (*models.Address, error) {
	lat, lng, err := a.Gr.GeocodeAddress(add)

	if err != nil {
		return nil, err
	}

	if lat == -1 || lng == -1 {
		return nil, errors.New("the geocoding process can't be processed with the address specified")
	}

	sa, err := a.Cr.ReadCSVFile()

	if err != nil {
		return nil, err
	}

	ad := &models.Address{
		ID: len(sa),
		A:  add,
		P: models.Point{
			Lat: lat,
			Lng: lng,
		},
	}
	return ad, nil
}
