package main

import (
	"PPM2PNG/converter"
	"PPM2PNG/reader"
	"PPM2PNG/writer"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	ppmFile    = flag.String("ppm-file", "", "Path of ppm file")
	outputFile = flag.String("output-file", "", "File to output")
	outputType = flag.String("output-type", "png", "Type of output picture.[PNG/JPG]")
)

func main() {
	help := flag.Bool("help", false, "Show Usage")
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if *ppmFile == "" || *outputFile == "" {
		if *ppmFile == "" {
			fmt.Print("ppm ")
		} else {
			fmt.Println("output ")
		}
		fmt.Println("File path is empty.")
		flag.Usage()
		os.Exit(1)
	}
	ppmContent, err := reader.ReadPPMFile(*ppmFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pic, err := converter.PPMConvert(ppmContent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	picType := strings.ToLower(*outputType)
	if picType == "png" {
		err = writer.DrawPngPic(*outputFile, pic)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if picType == "jpg" {
		err = writer.DrawJpgPic(*outputFile, pic)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Unknown picture type")
		flag.Usage()
		os.Exit(1)
	}
}
