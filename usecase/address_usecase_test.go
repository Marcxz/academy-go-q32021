package usecase

import (
	"testing"

	"github.com/Marcxz/academy-go-q32021/models"
	"github.com/Marcxz/academy-go-q32021/repository/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestReadCSVAddress - func to test ReadCSVAddress with mocks
func TestReadCSVAddress(t *testing.T) {
	testCases := []struct {
		name           string
		expectedLength int
		response       []models.Address
		hasError       bool
		error          error
		filter         string
	}{
		{
			name:           "Csv address file with 1 row",
			expectedLength: 1,
			response:       nil,
			hasError:       false,
			error:          nil,
			filter:         "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			aAddFound := []string{}

			defer ctrl.Finish()
			aAddFound = append(aAddFound, "1|wizeline, zapopan, jalisco, mexico|20.6443271|-103.4163436")

			mcr := mock.NewMockCsv(ctrl)
			mgr := mock.NewMockGeo(ctrl)
			mdr := mock.NewMockGeoDB(ctrl)

			auc := NewAddressUseCase(mcr, mgr, mdr)
			mcr.EXPECT().ReadCSVFile().Return(aAddFound, nil)
			aadd, err := auc.ReadCSVAddress()
			assert.EqualValues(t, tc.expectedLength, len(aadd))
			if tc.hasError {
				assert.EqualError(t, err, tc.error.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}

}

// TestGeocodeAddress - func to test GeocodeAddress with mocks

func TestGeocodeAddress(t *testing.T) {
	testCases := []struct {
		name           string
		expectedLength int
		response       []models.Address
		hasError       bool
		error          error
		filter         string
	}{
		{
			name:           "Geocode Wizeline address",
			expectedLength: 1,
			response:       nil,
			hasError:       false,
			error:          nil,
			filter:         "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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

			_, err := auc.GeocodeAddress("wizeline, guadalajara, méxico")
			if tc.hasError {
				assert.EqualError(t, err, tc.error.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}

}

// func TestStoreGeocodeAddress - func to test StoreGeocodeAddress with mocks
func TestStoreGeocodeAddress(t *testing.T) {
	testCases := []struct {
		name           string
		expectedLength int
		response       []models.Address
		hasError       bool
		error          error
		filter         string
	}{
		{
			name:           "Geocode Wizeline address",
			expectedLength: 1,
			response:       nil,
			hasError:       false,
			error:          nil,
			filter:         "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
			_, err = auc.StoreGeocodeAddress(add)
			if tc.hasError {
				assert.EqualError(t, err, tc.error.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}

}

// TestInvalidLineValidate - func to test an invalid address with inval
func TestInvalidLineValidate(t *testing.T) {
	err := validate(0, "0|INVALID|-1")
	assert.Equal(t, "the line at the index 0 should be composed for 4 pipes", err.Error())
}

// TestInvalidIDValidate - func to test an invalid id

func TestInvalidIDValidate(t *testing.T) {
	err := validate(0, "invalidID|address|-1|-1")
	assert.Equal(t, "the id at the index 0 should be integer invalidID", err.Error())
}

// TestInvalidLatValidate - func to test an invalid lat coordinate

func TestInvalidLatValidate(t *testing.T) {
	err := validate(0, "0|address|invalidLat|-1")
	assert.Equal(t, "the lat column at the index 0 should be float invalidLat", err.Error())
}

// TestInvalidLatValidate - func to test an invalid lng coordinate

func TestInvalidLngValidate(t *testing.T) {
	err := validate(0, "0|address|-1|invalidLng")
	assert.Equal(t, "the lng column at the index 0 should be float invalidLng", err.Error())
}
