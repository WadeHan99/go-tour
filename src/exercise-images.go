package main

import (
	"image"
	"image/color"

	"code.google.com/p/go-tour/pic"
)

// create the custom image type
type Image struct {
	W, H   int
	pColor uint8
}

// set the image dimensions and location
func (self *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.W, self.H)
}

// set the image color model
func (self *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (self *Image) At(x, y int) color.Color {
	return color.RGBA{self.pColor + uint8(x), self.pColor + uint8(y), 255, 255}
}

type Image2 struct {
	W, H int
}

func (self Image2) Bounds() image.Rectangle {
	return image.Rect(0, 0, self.W, self.H)
}

func (self Image2) ColorModel() color.Model {
	return color.RGBAModel
}

func (self Image2) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100, 128}
	pic.ShowImage(&m)
	m2 := Image2{100, 100}
	pic.ShowImage(m2)
}
