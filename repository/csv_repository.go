//go:generate mockgen -source csv_repository.go -destination mock/csv_repository_mock.go -package mock
package repository

import (
	"fmt"
	"strings"

	"github.com/Marcxz/academy-go-q32021/infraestructure"
)

type Reader interface {
	ReadCSVFile() ([]string, error)
}
type Storer interface {
	StoreAddressCSV(id int, a string, lat float64, lng float64) error
}

// csv - the interface for the csv repository
type Csv interface {
	Reader
	Storer
}

type cr struct {
	icsv infraestructure.Csv
}

// NewCsvRepository - func to create new csv repository used in usecase
func NewCsvRepository(i_csv infraestructure.Csv) Csv {
	return &cr{
		i_csv,
	}
}

// ReadCSVFile - func inteconnect repository with csv infraestructure to read csv files.
func (c *cr) ReadCSVFile() ([]string, error) {
	cl, err := c.icsv.ReadCSVFile()

	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	as := strings.Split(string(cl), "\n")

	return as[:len(as)-1], nil
}

// StoreAddressCSV - repository func to store an address in a csv file
func (c *cr) StoreAddressCSV(id int, a string, lat float64, lng float64) error {
	return c.icsv.StoreAddressCSV(id, a, lat, lng)
}
