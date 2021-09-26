package infraestructure

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReadCSVFile - Read a CSV file with a filename specified
func ReadCSVFile(f string) ([]byte, error) {
	p := fmt.Sprintf("%s%s", os.Getenv("bp"), f)
	l, err := ioutil.ReadFile(p)

	if err != nil {
		return nil, err
	}

	return l, nil
}

func StoreAddressCSV(fn string, id int, a string, lat float64, lng float64) error {
	p := fmt.Sprintf("%s%s", os.Getenv("bp"), fn)

	f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	a = fmt.Sprintf("%d|%s|%f|%f\n", id, a, lat, lng)
	_, err = f.Write([]byte(a))
	if err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
