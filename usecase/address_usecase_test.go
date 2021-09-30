package usecase
/*
import (
	"testing"

	"github.com/Marcxz/academy-go-q32021/models"
	"github.com/Marcxz/academy-go-q32021/repoitory/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReadCSVAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	aAddFound := []models.Address{
		models.Address{
			ID: 0,
			A:  "centro,guadalajara,jalisco",
			P: models.Point{
				Lat: 20.6866131,
				Lng: -103.3507872,
			},
		},
	}

	defer ctrl.Finish()

	mcr := mock.NewMockCsv(ctrl)
	mgr := mock.NewMockGeo(ctrl)

	auc := NewAddressUseCase(mcr, mgr)

	mcr.EXPECT().ReadCSVFile().Return(aAddFound, nil)
	aadd, err := auc.ReadCSVAddress()

	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if aadd == nil {
		t.Error("the array of addresses shouldn't be nil")
		t.Fail()
	}
}
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
*/