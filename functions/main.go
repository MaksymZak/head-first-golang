package main

import (
	"errors"
	"fmt"
	"log"
)

func paintNeeded(width, height float64) (float64, error) {
	err := errors.New("height must be greater than 0")
	if height <= 0 {
		return 0, err
	}

	if width <= 0 {
		return 0, fmt.Errorf("width must be greater than 0")
	}

	area := width * height
	return area / 10.0, nil
}

func main() {

	width := 5.1
	height := 4.2
	paint, err := paintNeeded(width, height)
	if err != nil {
		log.Fatal(err)
	}

	paintw, err := paintNeeded(width, -5.2)
	if err != nil {
		fmt.Println(err)
	}

	paint3, err := paintNeeded(-10, 20)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%0.2f liters needed", paint)
	fmt.Printf("%0.2f liters needed", paintw)
	fmt.Printf("%0.2f liters needed", paint3)
}
