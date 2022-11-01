package tutorial

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

func GreenBlackLissajous(out io.Writer) {
	var palette = []color.Color{color.RGBA{R: 0xAB, G: 0x25, B: 0xFF, A: 0xff}, color.RGBA{R: 0x00, G: 0xFF, B: 0x25, A: 0xff}, color.Black}
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 1024  // image canvas covers [-size..+size]
		nframes = 64    // number of frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Float64()*3.0))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
