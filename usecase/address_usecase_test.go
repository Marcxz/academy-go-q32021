package usecase

import (
	"testing"

	"github.com/Marcxz/academy-go-q32021/repository/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReadCSVAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	aAddFound := []string{}

	defer ctrl.Finish()

	mcr := mock.NewMockCsv(ctrl)
	mgr := mock.NewMockGeo(ctrl)
	mdr := mock.NewMockGeoDB(ctrl)

	auc := NewAddressUseCase(mcr, mgr, mdr)

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

func TestGeocodeAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	add := "wizeline, guadalajara, méxico"
	defer ctrl.Finish()
	aAddFound := []string{}
	mcr := mock.NewMockCsv(ctrl)
	mgr := mock.NewMockGeo(ctrl)
	mdr := mock.NewMockGeoDB(ctrl)

	auc := NewAddressUseCase(mcr, mgr, mdr)
	mgr.EXPECT().GeocodeAddress(add)
	mcr.EXPECT().ReadCSVFile().Return(aAddFound, nil)

	am, err := auc.GeocodeAddress("wizeline, guadalajara, méxico")

	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if am == nil {
		t.Error("the address shouldn't be nil")
		t.Fail()
	}
}

func TestStoreGeocodeAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	add := "wizeline, guadalajara, mexico"

	defer ctrl.Finish()
	aAddFound := []string{}

	mgr := mock.NewMockGeo(ctrl)
	mcr := mock.NewMockCsv(ctrl)
	mdr := mock.NewMockGeoDB(ctrl)

	auc := NewAddressUseCase(mcr, mgr, mdr)

	mgr.EXPECT().GeocodeAddress(add)
	mcr.EXPECT().ReadCSVFile().Return(aAddFound, nil)
	var err error

	mcr.EXPECT().StoreAddressCSV(0, add, 0.0, 0.0).Return(err)

	if err != nil {
		t.Error(err.Error())
		t.Fatal()
	}
	a, _ := auc.StoreGeocodeAddress(add)

	if err != nil {
		t.Error(err.Error())
		t.Fatal()
	}

	if a == nil {
		t.Error("the address shouldn't be nil")
		t.Fatal()
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
