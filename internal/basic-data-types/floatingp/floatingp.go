package floatingp

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

const (
	mountainEnum = "mountain"
	eggboxEnum   = "eggbox"
)

// pointFunctions take 2 floats and return a single float
type pointFunction func(float64, float64) float64

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func Serve() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	Surface(w, mountainEnum, true)
}

func Surface(out io.Writer, choice string, color bool) {
	var function pointFunction
	var zmax, zmin float64

	switch choice {
	case mountainEnum:
		function = mountain
	case eggboxEnum:
		function = eggbox
	default:
		function = mountain
	}

	zmax, zmin = getMinMax(function)

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _, err := corner(function, i+1, j)
			if err != nil {
				continue
			}
			bx, by, _, err := corner(function, i, j)
			if err != nil {
				continue
			}
			cx, cy, _, err := corner(function, i, j+1)
			if err != nil {
				continue
			}
			dx, dy, z, err := corner(function, i+1, j+1)
			if err != nil {
				continue
			}
			if color {
				fmt.Fprintf(out, "<polygon style='fill:%s;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					getColor(z, zmin, zmax), ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(fn pointFunction, i int, j int) (float64, float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := fn(x, y)
	if math.IsInf(z, 0) {
		return 0, 0, 0, errors.New("non finite polygon generated")
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

func mountain(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 {
	r := math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 12 // distance from (0,0)
	return r
}

func getColor(z float64, zmin float64, zmax float64) string { // put zmin and zmax in globals

	if z > 0 {
		gradient := math.Exp(math.Abs(z)) / math.Exp(math.Abs(zmax)) * 255
		return fmt.Sprintf("#%02x0000", int(math.Round(gradient)))
	} else {
		gradient := math.Exp(math.Abs(z)) / math.Exp(math.Abs(zmin)) * 255
		return fmt.Sprintf("#0000%02x", int(math.Round(gradient)))
	}

}

func getMinMax(fn pointFunction) (float64, float64) {
	min := math.NaN()
	max := math.NaN()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xoff := 0; xoff <= 1; xoff++ {
				for yoff := 0; yoff <= 1; yoff++ {
					x := xyrange * (float64(i+xoff)/cells - 0.5)
					y := xyrange * (float64(j+yoff)/cells - 0.5)
					z := fn(x, y)
					if math.IsNaN(min) || z < min {
						min = z
					}
					if math.IsNaN(max) || z > max {
						max = z
					}
				}
			}
		}
	}

	return max, min
}
