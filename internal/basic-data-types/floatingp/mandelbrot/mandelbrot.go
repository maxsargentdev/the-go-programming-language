package mandelbrot

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"os"
	"time"
)

func Render() {
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
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
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

func RenderAll() {
	renderType("64")
	renderType("128")
	renderType("bigfloat")
	renderType("bigrat")
}

func renderType(rendertype string) {

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	start := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			switch {
			case rendertype == "64":
				img.Set(px, py, mandelbrot64(z))
			case rendertype == "128":
				img.Set(px, py, mandelbrot128(z))
			case rendertype == "bigfloat":
				img.Set(px, py, mandelbrotBigFloat(z))
			case rendertype == "bigrat":
				img.Set(px, py, mandelbrotRat(z))
				fmt.Printf("Pixel %d,%d complete\n", px, py)
			}
		}
	}
	outputFile, _ := os.Create(fmt.Sprintf("%s.png", rendertype))
	png.Encode(outputFile, img)
	outputFile.Close()
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", rendertype, elapsed)
	// write each one to a file with its name and time it
}

func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
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

func mandelbrot128(z complex128) color.Color {
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

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200
	zR := (&big.Float{}).SetFloat64(real(z))
	zI := (&big.Float{}).SetFloat64(imag(z))
	var vR, vI = &big.Float{}, &big.Float{}
	for i := uint8(0); i < iterations; i++ {
		// (r+i)^2 = r^2 + 2ri + i^2
		vR2, vI2 := &big.Float{}, &big.Float{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Float{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Float{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Float{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewFloat(4)) == 1 {
			switch {
			case i > 50: // dark red
				return color.RGBA{100, 0, 0, 255}
			default:
				// logarithmic blue gradient to show small differences on the
				// periphery of the fractal.
				logScale := math.Log(float64(i)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}

func mandelbrotRat(z complex128) color.Color {
	// High-resolution images take an extremely long time to render with
	// iterations = 200. Multiplying arbitrary precision numbers has
	// algorithmic complexity of at least O(n*log(n)*log(log(n)))
	// (https://en.wikipedia.org/wiki/Arbitrary-precision_arithmetic#Implementation_issues).
	const iterations = 200
	zR := (&big.Rat{}).SetFloat64(real(z))
	zI := (&big.Rat{}).SetFloat64(imag(z))
	var vR, vI = &big.Rat{}, &big.Rat{}
	for i := uint8(0); i < iterations; i++ {
		// (r+i)^2 = r^2 + 2ri + i^2
		vR2, vI2 := &big.Rat{}, &big.Rat{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Rat{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewRat(2, 1)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Rat{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Rat{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewRat(4, 1)) == 1 {
			switch {
			case i > 50: // dark red
				return color.RGBA{100, 0, 0, 255}
			default:
				// logarithmic blue gradient to show small differences on the
				// periphery of the fractal.
				logScale := math.Log(float64(i)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}
