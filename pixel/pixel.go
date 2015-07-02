package pixel

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

// Pixel utility
type Pixel struct {
	Image image.Image
	RGBA  *image.RGBA
}

// HSV is color space
type HSV struct {
	H     uint16
	S     uint16
	V     uint16
	Alpha uint16
}

// RGBA64 is alias of color.RGBA
type RGBA64 color.RGBA64

// ToRGBA64 is converter from HSV
func (hsv *HSV) ToRGBA64() (rgba *RGBA64) {

	return
}

// ToHSV is converter from RGBA64
func (rgba *RGBA64) ToHSV() (hsv *HSV) {
	max := math.Max(float64(rgba.R), math.Max(float64(rgba.G), float64(rgba.B)))
	min := math.Min(float64(rgba.R), math.Min(float64(rgba.G), float64(rgba.B)))
	hue := max - min
	hsv.S = 0
	hsv.V = uint16(max)
	hsv.Alpha = rgba.A

	if max != 0 {
		hsv.S = uint16(0xFFFF * (hue / max))
	}

	switch uint16(max) {
	case rgba.R:
		hue = float64(uint16(60*float64(rgba.G-rgba.B)/hue+360) % 360)
	case rgba.G:
		hue = 60*float64(rgba.B-rgba.R)/hue + 120
	case rgba.B:
		hue = 60*float64(rgba.R-rgba.G)/hue + 240
	}

	hsv.H = uint16(hue)

	return
}

// NewPixel is Pixel constructor
func NewPixel(file *os.File) *Pixel {
	defer file.Close()

	img, err := png.Decode(file)

	if err != nil {
		panic(err)
	}

	pixel := &Pixel{img, nil}

	rect := img.Bounds()
	pixel.RGBA = image.NewRGBA(rect)
	pixel.Each(pixel.RGBA.Set)

	return pixel
}

// Each method
func (pixel *Pixel) Each(f func(int, int, color.Color)) {
	rect := pixel.Image.Bounds()

	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			f(x, y, pixel.Image.At(x, y))
		}
	}

	pixel.Image = pixel.RGBA.SubImage(rect)
}

// Map method
func (pixel *Pixel) Map(f func(int, int, color.Color) color.Color) {
	rect := pixel.Image.Bounds()

	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			col := f(x, y, pixel.Image.At(x, y))
			pixel.RGBA.Set(x, y, col)
		}
	}

	pixel.Image = pixel.RGBA.SubImage(rect)
}

// Save method
func (pixel *Pixel) Save(file *os.File) {
	defer file.Close()

	png.Encode(file, pixel.Image)
}
