/* Lissajous generates GIF animations of random Lissajous figures.
Uasge:
	go build lissajous.go
	./lissajous >out.gif
	Or:
		go build goBible/chapter1/exam1-5
		./exam1-5 >out.gif
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//!+ Only this different from exam1.4
var palette = []color.Color{color.Black, color.RGBA{0, 0xff, 0, 0xff}}

//!-

const (
	whiteIndex = iota // first color in the palette
	blackIndex        // next color in the palette
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:9000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5     // number of complete x oscillator revolutions
		res    = 0.001 // angular resolution
		size   = 100   //images canvas cover [-size...+size]
		nframs = 64    // number of animation frames
		delay  = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframs}
	phase := 0.0 // phase difference
	for i := 0; i < nframs; i++ {
		rect := image.Rect(0, 0, 2*size, 2*size)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
