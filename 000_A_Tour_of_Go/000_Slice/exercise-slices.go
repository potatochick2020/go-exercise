package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// uint range [0 : 255]

	// make 2D array.
	returnValue := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		returnValue[i] = make([]uint8, dx)
	}

	// generate value
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			returnValue[i][j] = uint8(i ^ j)
		}
	}
	return returnValue
}

func main() {
	pic.Show(Pic)
}
