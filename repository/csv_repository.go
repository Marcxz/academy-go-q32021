package repository

import (
	"fmt"
	"strings"

	"github.com/Marcxz/academy-go-q32021/infraestructure"
)

// csv - the interface for the csv repository
type csv interface {
	ReadCSVFile(f string) ([]string, error)
	StoreAddressCSV(f string, id int, a string, lat float64, lng float64) error
}

type cr struct{}

// NewCsvRepository - func to create new csv repository used in usecase
func NewCsvRepository() csv {
	return &cr{}
}

// ReadCSVFile - func inteconnect repository with csv infraestructure to read csv files.
func (*cr) ReadCSVFile(f string) ([]string, error) {
	cl, err := infraestructure.ReadCSVFile(f)

	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	as := strings.Split(string(cl), "\n")

	return as[:len(as)-1], nil
}

// StoreAddressCSV - repository func to store an address in a csv file
func (*cr) StoreAddressCSV(f string, id int, a string, lat float64, lng float64) error {
	return infraestructure.StoreAddressCSV(f, id, a, lat, lng)
}
