package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func RenderSuperSampled() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		supersamplefactor      = 2
	)

	supersamplewidth := width * supersamplefactor
	supersampleheight := height * supersamplefactor
	supersampleimg := image.NewRGBA(image.Rect(0, 0, supersamplewidth, supersampleheight))

	// create fractal at super sample factor
	for py := 0; py < supersampleheight; py++ {
		y := float64(py)/float64(supersampleheight)*(ymax-ymin) + ymin
		for px := 0; px < supersamplewidth; px++ {
			x := float64(px)/float64(supersamplewidth)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			supersampleimg.Set(px, py, mandelbrotSuperSampled(z))
		}
	}

	// now run super sample
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < supersampleheight; py = py + 2 {
		for px := 0; px < supersamplewidth; px = px + 2 {
			var c = []color.RGBA{
				supersampleimg.RGBAAt(px, py),
				supersampleimg.RGBAAt(px+1, py),
				supersampleimg.RGBAAt(px, py+1),
				supersampleimg.RGBAAt(px+1, py+1),
			}
			var pr, pg, pb, pa int
			for n := 0; n < 4; n++ {
				pr += int(c[n].R)
				pg += int(c[n].G)
				pb += int(c[n].B)
				pa += int(c[n].A)
			}
			img.SetRGBA(px/2, py/2, color.RGBA{uint8(pr / 4), uint8(pg / 4), uint8(pb / 4), uint8(pa / 4)})
		}
	}

	// super sample
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrotSuperSampled(z complex128) color.Color {
	const iterations = 200
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50: // dark red
				return color.RGBA{100, 0, 0, 255}
			default:
				// logarithmic blue gradient to show small differences on the
				// periphery of the fractal.
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}
