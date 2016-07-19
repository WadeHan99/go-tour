package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// initialize the image
	image := make([][]uint8, dy)

	// iterate over each pixel
	for x := 0; x < dy; x++ {
		image[x] = make([]uint8, dx)
		for y := 0; y < dx; y++ {

			// set a color value for each pixel
			image[x][y] = uint8((x ^ y) * (x ^ y))
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}
