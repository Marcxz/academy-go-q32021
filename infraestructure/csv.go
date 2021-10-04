package infraestructure

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Marcxz/academy-go-q32021/conf"
)

type Reader interface {
	ReadCSVFile() ([]byte, error)
}

type Storer interface {
	StoreAddressCSV(id int, a string, lat float64, lng float64) error
}

type Csv interface {
	Reader
	Storer
}

type ic struct {
	con *conf.Config
}

func NewCsvInfraestructure(cfg *conf.Config) Csv {
	return &ic{
		cfg,
	}
}

// ReadCSVFile - Read a CSV file with a filename specified
func (c *ic) ReadCSVFile() ([]byte, error) {
	p := fmt.Sprintf("%s%s", c.con.Base_path, c.con.Filename)
	l, err := ioutil.ReadFile(p)

	if err != nil {
		return nil, err
	}

	return l, nil
}

func (c *ic) StoreAddressCSV(id int, a string, lat float64, lng float64) error {
	p := fmt.Sprintf("%s%s", c.con.Base_path, c.con.Filename)

	f, err := os.OpenFile(p, os.O_APPEND, 0644)
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
