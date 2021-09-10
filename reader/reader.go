package reader

import (
	"io/ioutil"
	"os"
)

func ReadPPMFile(filePath string) (ppmContent []byte, err error) {
	ppm, err := os.Open(filePath)
	if err != nil {
		return
	}
	ppmContent, err = ioutil.ReadAll(ppm)
	if err != nil {
		return
	}
	return
}
