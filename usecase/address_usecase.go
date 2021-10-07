package usecase

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Marcxz/academy-go-q32021/models"
	"github.com/Marcxz/academy-go-q32021/repository"
)

// Reader - interface that makes the contract read a csv address file
type Reader interface {
	ReadCSVAddress() ([]models.Address, error)
}

// Geolocater - interface that makes the contract to geocode an address and return as a model address
type Geolocater interface {
	GeocodeAddress(string) (*models.Address, error)
}

// Storer - interface that makes the contract to geocode an address and append it in the csv file
type Storer interface {
	StoreGeocodeAddress(string) (*models.Address, error)
}

// Concurrencier - interface that makes the contract to read a csv file with concurrency
type Concurrencier interface {
	ReadAddressCSVConcurrency(string, int, int) (map[int]models.Address, error)
}

// Router - interface that makes the contract to generate the route from 2 address
type Router interface {
	GenerateRouterFrom2Address(string, string) (models.Route, error)
}

// Address - the interace for the address usecase
type Address interface {
	Reader
	Geolocater
	Storer
	Concurrencier
	Router
}

// auc - the struct that isolate the usecase layer func with the rest
type auc struct {
	Cr  repository.Csv
	Gr  repository.Geo
	Dbr repository.GeoDB
}

// NewAddressUseCase - func to create a new address usecase used to link with the controller, it accomplishes the Address interface requeriments
func NewAddressUseCase(rcsv repository.Csv, rgeo repository.Geo, rgeodb repository.GeoDB) Address {
	return &auc{
		rcsv,
		rgeo,
		rgeodb,
	}
}

// ReadCSVAddress - func to do the bussiness logic when you read all the address from a csv file, returns an model address array and an error if exist
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

// GeocodeAddress - func to geocode an address and return an address object or an error if exist
func (a *auc) GeocodeAddress(add string) (*models.Address, error) {
	ad, err := a.createGeocodeAddress(add)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

// StoreGeocodeAddress - func to geocode an address and push it in the csv file, returns the address model and an error if exist
func (a *auc) StoreGeocodeAddress(add string) (*models.Address, error) {

	ad, _ := a.createGeocodeAddress(add)
	err := a.Cr.StoreAddressCSV(ad.ID, ad.A, ad.P.Lat, ad.P.Lng)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

// validate - func to validate if an string address has the minimun requirements to be an address struc, returns an error if exist
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

// createGeocodeAddress - general func to geocode and generate an address, returns an address object or an error if exist
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

// ReadAddressCSVConcurrency - func to read a csv address file with concurrency workers and chanels, return a map with the address or an error if exist
func (a *auc) ReadAddressCSVConcurrency(t string, ias int, ipw int) (map[int]models.Address, error) {
	af, err := a.Cr.ReadCSVFile()
	if err != nil {
		return nil, err
	}
	if len(af) == 0 {
		return nil, fmt.Errorf("the file addresses is empty")
	}
	rem := float64(ipw)
	i := 0
	m := map[int]models.Address{}
	nw := 1.0
	var aj []chan string

	// Construct workers and channels
	if ias > 0 && ipw > 0 {
		if ias < len(af) {
			nw = float64(ias) / float64(ipw)
		} else {
			nw = float64(len(af)) / float64(ipw)
		}

		aj = make([]chan string, 0)
		for ij := 0; ij < int(nw); ij++ {
			j := make(chan string, ipw)
			aj = append(aj, j)
		}
		rem = (nw - float64(int(nw))) * float64(ipw)
		if rem > 0 {
			j := make(chan string, int(rem))
			aj = append(aj, j)
		}
	} else {
		aj = make([]chan string, 0)
		j := make(chan string)
		aj = append(aj, j)
	}

	// Sending data to channels
	i = 0
	for idx, j := range aj {
		//last
		if idx == len(aj)-1 && rem > 0 {
			at := af[i : i+int(rem)]
			go sendDataWorker(j, t, at)
			fmt.Printf("worker %d ended and read %d items \n", idx, len(at))
			i = i + int(rem)
		} else {
			at := af[i : i+ipw]
			go sendDataWorker(j, t, at)
			fmt.Printf("worker %d ended and read %d items \n", idx, len(at))
			i = i + ipw
		}
	}
	// Receiving data from channels
	for _, j := range aj {
		for lv := range j {
			al := strings.Split(lv, "|")
			id, err := strconv.Atoi(al[0])
			if err != nil {
				break
			}
			m[id], err = createAddress(id, lv)
			if err != nil {
				break
			}
		}
	}
	return m, nil
}

// sendDataWorker - factory func to select the address if it's ID is even, odd or all, the result is stored into the channel worker
func sendDataWorker(j chan string, t string, af []string) {
	for _, l := range af {
		al := strings.Split(l, "|")
		switch t {
		case "even":
			id, _ := strconv.Atoi(al[0])
			if id%2 == 0 {
				j <- l
			}
		case "odd":
			id, _ := strconv.Atoi(al[0])
			if id%2 != 0 {
				j <- l
			}
		default:
			j <- l
		}
	}
	close(j)
}

// createAddress - general func to create an address model from an id and address csv file, returns the moddel address or error if exist
func createAddress(id int, l string) (models.Address, error) {
	a := models.Address{}
	err := validate(id, l)

	if err != nil {
		return a, err
	}

	al := strings.Split(l, "|")

	lat, err := strconv.ParseFloat(al[2], 64)
	if err != nil {
		return a, err
	}

	lng, err := strconv.ParseFloat(al[2], 64)
	if err != nil {
		return a, err
	}

	a = models.Address{
		ID: id,
		A:  al[1],
		P: models.Point{
			Lat: lat,
			Lng: lng,
		},
	}

	return a, nil
}

// GenerateRouterFrom2Address - func to do the bussiness logic to generate 2 address model, a route from 1 point to another and send it to the controller, returns a model route or an error if exist
func (a *auc) GenerateRouterFrom2Address(from string, to string) (models.Route, error) {
	var r models.Route
	fa, err := a.createGeocodeAddress(from)
	if err != nil {
		return r, err
	}

	ta, err := a.createGeocodeAddress(to)
	if err != nil {
		return r, err
	}

	ar, err := a.Dbr.GenerateRoute(fa.P.Lat, fa.P.Lng, ta.P.Lat, ta.P.Lng)
	if err != nil {
		return r, err
	}
	r = models.Route{
		ID:   0,
		Name: fmt.Sprintf("FROM %s TO %s", from, to),
		From: *fa,
		To:   *ta,
		R:    ar,
	}
	return r, nil
}
