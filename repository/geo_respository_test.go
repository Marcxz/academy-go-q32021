package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidGeocodeAddress(t *testing.T) {
	gr := NewGeoRepository()
	lat, lng, _ := gr.GeocodeAddress("ñlsadjfñlkasdjflksdajflkasdjfñlaksdjfadsñlkjf")
	assert.Equal(t, lat, -1.0)
	assert.Equal(t, lng, -1.0)
}
