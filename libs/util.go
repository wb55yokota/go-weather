package libs

import (
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func Cp932(str string) (string, error) {
	r, _, err := transform.String(japanese.ShiftJIS.NewEncoder(), str)
	if err != nil {
		panic(err)
	}
	return r, err
}

func SaveStringToFile(path string, str string) (bool, error) {
	file, err := os.Create(path)
	if err != nil {
		return false, err
	}
	defer file.Close()
	cp932str, err := Cp932(str)
	if err != nil {
		return false, err
	}
	bytes := []byte(cp932str)
	_, err = file.Write(bytes)
	if err != nil {
		return false, err
	}
	return true, nil
}
