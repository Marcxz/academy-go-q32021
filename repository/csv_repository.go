//go:generate mockgen -source csv_repository.go -destination mock/csv_repository_mock.go -package mock
package repository

import (
	"fmt"
	"strings"

	"github.com/Marcxz/academy-go-q32021/infraestructure"
)

// Reader - interface that makes the contract to read a csv address file
type Reader interface {
	ReadCSVFile() ([]string, error)
}

// Storer - interface that makes the contract to push the address into csv file
type Storer interface {
	StoreAddressCSV(id int, a string, lat float64, lng float64) error
}

// Csv - interface that concatenates all repository interfaces
type Csv interface {
	Reader
	Storer
}

// cr - struct that we use to isolate the repository layer from the rest
type cr struct {
	icsv infraestructure.Csv
}

// NewCsvRepository - repository constructor func to create a repository to link with the usecase layer which accomplish the Csv requirements
func NewCsvRepository(i_csv infraestructure.Csv) Csv {
	return &cr{
		i_csv,
	}
}

// ReadCSVFile - inteconnect repository with csv infraestructure to read csv files, returns an array string with the csv data and an error if exist
func (c *cr) ReadCSVFile() ([]string, error) {
	cl, err := c.icsv.ReadCSVFile()

	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	as := strings.Split(string(cl), "\n")

	return as[:len(as)-1], nil
}

// StoreAddressCSV - func to send to infraestructure the data to store an address, returns an error if exist
func (c *cr) StoreAddressCSV(id int, a string, lat float64, lng float64) error {
	return c.icsv.StoreAddressCSV(id, a, lat, lng)
}
