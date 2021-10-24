package infraestructure

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Marcxz/academy-go-q32021/conf"
)

// reader - interface that makes the contract to read a cvs file
type reader interface {
	ReadCSVFile() ([]byte, error)
}

// storer - interface that makes the contract to push an address to a csv file
type storer interface {
	StoreAddressCSV(id int, a string, lat float64, lng float64) error
}

// Csv - interface for the csv infraestructure
type Csv interface {
	reader
	storer
}

// ic - the struct that we use to isolate infraestructure with repository
type ic struct {
	con *conf.Config
}

// NewCsvInfraestructure - csv infraestructure constructor, returns csv interface
func NewCsvInfraestructure(cfg *conf.Config) Csv {
	return &ic{
		cfg,
	}
}

// ReadCSVFile - read a csv addresses file, returns a byte arrays of all the addresses and an error if exist
func (c *ic) ReadCSVFile() ([]byte, error) {
	p := fmt.Sprintf("%s%s", c.con.BasePath, c.con.Filename)
	l, err := ioutil.ReadFile(p)

	if err != nil {
		return nil, err
	}

	return l, nil
}

// StoreAddressCSV - func to push an address in a csv file, returns error if exist
func (c *ic) StoreAddressCSV(id int, a string, lat float64, lng float64) error {
	p := fmt.Sprintf("%s%s", c.con.BasePath, c.con.Filename)

	f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	a = fmt.Sprintf("%d|%s|%f|%f\n", id, a, lat, lng)
	_, err = f.Write([]byte(a))
	if err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
