package converter

import (
	"errors"
	"image"
	"image/color"
	"strconv"
)

func readLine(src []byte, pos int) (line []byte, end bool, newPos int) {
	newPos = pos
	for len(src) > newPos {
		if src[newPos] == '\r' {
			newPos++
		}
		if src[newPos] == '\n' {
			newPos++
			return
		}
		line = append(line, src[newPos])
		newPos++
	}
	end = true
	return
}

func getNumbersInLine(src []byte) (numbers []int, err error) {
	pos := 0
	var numStr []byte
	for len(src) > pos {
		if src[pos] != ' ' {
			numStr = append(numStr, src[pos])
		} else {
			var num int
			num, err = strconv.Atoi(string(numStr))
			if err != nil {
				return
			}
			numbers = append(numbers, num)
			numStr = numStr[0:0]
		}
		pos++
	}
	var num int
	num, err = strconv.Atoi(string(numStr))
	if err != nil {
		return
	}
	numbers = append(numbers, num)
	return
}

func ppmChecker(ppmContent []byte) (width, height, maxPixel, pos int, err error) {
	for i := 0; i < 3; i++ {
		var line []byte
		var end bool
		line, end, pos = readLine(ppmContent, pos)
		if end {
			err = errors.New("invalid format")
			return
		}
		if i == 0 && string(line) != "P3" {
			err = errors.New("invalid format")
			return
		} else if i == 1 {
			var numbers []int
			numbers, err = getNumbersInLine(line)
			if err != nil {
				return
			}
			if len(numbers) != 2 {
				err = errors.New("invalid format")
				return
			}
			width = numbers[0]
			height = numbers[1]
		} else if i == 2 {
			var numbers []int
			numbers, err = getNumbersInLine(line)
			if err != nil {
				return
			}
			if len(numbers) != 1 {
				err = errors.New("invalid format")
				return
			}
			maxPixel = numbers[0]
		}
	}
	return
}

func PPMConvert(ppmContent []byte) (*image.RGBA, error) {
	width, height, maxPixel, pos, err := ppmChecker(ppmContent)
	if err != nil {
		return nil, err
	}
	var numbers []int
	end := false
	for !end {
		var line []byte
		line, end, pos = readLine(ppmContent, pos)
		if len(line) == 0 {
			continue
		}
		tmpNums, err := getNumbersInLine(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, tmpNums...)
	}
	if len(numbers) != width*height*3 {
		return nil, errors.New("invalid format")
	}
	pic := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: width, Y: height}})
	numPos := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := numbers[numPos]
			g := numbers[numPos+1]
			b := numbers[numPos+2]
			numPos += 3
			if r > maxPixel || g > maxPixel || b > maxPixel {
				return nil, errors.New("invalid format")
			}
			scale := float64(maxPixel+1) / 256.0
			r = int(float64(r) / scale)
			g = int(float64(g) / scale)
			b = int(float64(b) / scale)
			pic.Set(x, y, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255})
		}
	}
	return pic, nil
}
