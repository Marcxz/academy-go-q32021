package infraestructure 

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInvalidGeocodingAddress(t *testing.T) {
	gi := NewGeoInfraestructure()
	lat, lng, _ := gi.GeocodingAddress("asldkjfañsdlkjfasdlkfjsdalkfjasdñflkjasdflkasdjflsdakjfasdñlkfjasdflkj")
	assert.Equal(t, lat, -1.0)  
	assert.Equal(t, lng, -1.0) 
}