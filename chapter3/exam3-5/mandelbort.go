package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// 尽量多的定义常用的色彩
var Paletted = [16]color.Color{
	color.RGBA{66, 30, 15, 255},    // # brown 3
	color.RGBA{25, 7, 26, 255},     // # dark violett
	color.RGBA{9, 1, 47, 255},      //# darkest blue
	color.RGBA{4, 4, 73, 255},      //# blue 5
	color.RGBA{0, 7, 100, 255},     //# blue 4
	color.RGBA{12, 44, 138, 255},   //# blue 3
	color.RGBA{24, 82, 177, 255},   //# blue 2
	color.RGBA{57, 125, 209, 255},  //# blue 1
	color.RGBA{134, 181, 229, 255}, // # blue 0
	color.RGBA{211, 236, 248, 255}, // # lightest blue
	color.RGBA{241, 233, 191, 255}, // # lightest yellow
	color.RGBA{248, 201, 95, 255},  // # light yellow
	color.RGBA{255, 170, 0, 255},   // # dirty yellow
	color.RGBA{204, 128, 0, 255},   // # brown 0
	color.RGBA{153, 87, 0, 255},    // # brown 1
	color.RGBA{106, 52, 3, 255},    // # brown 2
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbort(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbort(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		// ƒ(z) = z^2 + c
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return Paletted[n%16]
		}
	}
	return color.Black
}
