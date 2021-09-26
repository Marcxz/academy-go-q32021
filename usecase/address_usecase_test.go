package usecase

import (
	"testing"

	"github.com/Marcxz/academy-go-q32021/conf"

	"github.com/stretchr/testify/assert"
)

func TestInvalidLineValidate(t *testing.T) {
	err := validate(0, "0|INVALID|-1")
	assert.Equal(t, "the line at the index 0 should be composed for 4 pipes", err.Error())
}
func TestInvalidIDValidate(t *testing.T) {
	err := validate(0, "invalidID|address|-1|-1")
	assert.Equal(t, "the id at the index 0 should be integer invalidID", err.Error())
}

func TestInvalidLatValidate(t *testing.T) {
	err := validate(0, "0|address|invalidLat|-1")
	assert.Equal(t, "the lat column at the index 0 should be float invalidLat", err.Error())
}

func TestInvalidLngValidate(t *testing.T) {
	err := validate(0, "0|address|-1|invalidLng")
	assert.Equal(t, "the lng column at the index 0 should be float invalidLng", err.Error())
}
func TestInvalidReadCSVAddress(t *testing.T) {
	au := NewAddressUseCase()
	_, err := au.ReadCSVAddress("")
	assert.Equal(t, "open : The system cannot find the file specified.", err.Error())

}
func TestInvalidGeocodeAddress(t *testing.T) {
	conf.ConfigInit()
	au := NewAddressUseCase()
	_, err := au.GeocodeAddress("lasdjfdsalkfj")
	assert.Equal(t, "the geocoding process can't be processed with the address specified", err.Error())
}

func TestInvalidStoreGeocodeAddress(t *testing.T) {
	conf.ConfigInit()
	au := NewAddressUseCase()
	_, err := au.StoreGeocodeAddress("lsadjflkasdñjf")
	assert.Equal(t, "the geocoding process can't be processed with the address specified", err.Error())
}
