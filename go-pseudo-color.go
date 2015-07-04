package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/ww24/go-pseudo-color/pixel"
	"github.com/ww24/go-pseudo-color/pseudo"
)

func main() {

	flag.Parse()

	mode, filepathIn, filepathOut := flag.Arg(0), flag.Arg(1), flag.Arg(2)

	if mode == "" || filepathIn == "" || filepathOut == "" {
		fmt.Fprintln(os.Stderr, "usage: go-pseudo-color mode in_path out_path")
		fmt.Fprintln(os.Stderr, "mode = (linear, sigmoid, sin)")
		_main()
		return
	}

	file, err := os.Open(filepathIn)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	inPix := pixel.NewPixel(file)
	var outPix *pixel.Pixel

	switch mode {
	case "linear":
		outPix = inPix.Map(pseudo.ConvLinear)
	case "sigmoid":
		outPix = inPix.Map(pseudo.ConvSigmoid)
	case "sin":
		outPix = inPix.Map(pseudo.ConvSin)
	default:
		fmt.Fprintln(os.Stderr, "mode = (linear, sigmoid, sin)")
		return
	}

	file, err = os.Create(filepathOut)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	outPix.Save(file)
}

// line.png 生成
func _main() {
	pix := new(pixel.Pixel)
	rect := image.Rect(0, 0, 500, 60)
	pix.RGBA = image.NewRGBA(rect)
	pix.Image = pix.RGBA.SubImage(rect)
	pix = pix.Map(func(x int, _ int, _ color.Color) (c color.Color) {
		l := uint16(0xFFFF * x / 500)
		c = color.Gray16{l}
		return
	})

	file, err := os.Create("fixture/line.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pix.Save(file)
}
