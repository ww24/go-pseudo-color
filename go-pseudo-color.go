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

	pixel := pixel.NewPixel(file)

	switch mode {
	case "linear":
		pixel.Map(pseudo.Convlinear)
	case "sigmoid":
		pixel.Map(pseudo.Convsigmoid)
	case "sin":
		pixel.Map(pseudo.Convsin)
	default:
		fmt.Fprintln(os.Stderr, "mode = (linear, sigmoid, sin)")
		return
	}

	file, err = os.Create(filepathOut)
	if err != nil {
		panic(err)
	}

	pixel.Save(file)
}

// line.png 生成
func _main() {
	pixel := new(pixel.Pixel)
	rect := image.Rect(0, 0, 500, 60)
	pixel.RGBA = image.NewRGBA(rect)
	pixel.Image = pixel.RGBA.SubImage(rect)
	pixel.Map(func(x int, _ int, _ color.Color) (c color.Color) {
		l := uint16(0xFFFF * x / 500)
		c = color.Gray16{l}
		return
	})

	file, err := os.Create("fixture/line.png")
	if err != nil {
		panic(err)
	}

	pixel.Save(file)
}
