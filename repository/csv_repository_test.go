package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyReadCSVFile(t *testing.T) {
	cr := NewCsvRepository()
	_, err := cr.ReadCSVFile("")
	assert.Equal(t, "open : The system cannot find the file specified.", err.Error())
}

func TestInvalidReadCSVFile(t *testing.T) {
	cr := NewCsvRepository()
	_, err := cr.ReadCSVFile("invalid.csv")
	assert.Equal(t, "open invalid.csv: The system cannot find the file specified.", err.Error())
}

func TestEmptyStoreAddressCSV(t *testing.T) {
	cr := NewCsvRepository()
	err := cr.StoreAddressCSV("", -1, "", -1.0, 1-0)
	assert.Equal(t, "open : The system cannot find the file specified.", err.Error())
}