package infraestructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyReadCSVFile(t *testing.T) {
	_, err := ReadCSVFile("")
	assert.Equal(t, "open : The system cannot find the file specified.", err.Error())
}

func TestEmptyStoreAddressCSV(t *testing.T) {
	err := StoreAddressCSV("", 0, "", -1, -1)
	assert.Equal(t, "open : The system cannot find the file specified.", err.Error())
}
