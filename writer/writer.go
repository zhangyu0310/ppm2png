package writer

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func DrawPngPic(target string, pic *image.RGBA) error {
	tarFile, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	err = png.Encode(tarFile, pic)
	return err
}

func DrawJpgPic(target string, pic *image.RGBA) error {
	tarFile, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	err = jpeg.Encode(tarFile, pic, nil)
	return err
}
