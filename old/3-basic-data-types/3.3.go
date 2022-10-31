package basicdatatypes

import (
	"errors"
	"fmt"
	"math"
)

var zmax, zmin float64

func init() {
	min := math.NaN()
	max := math.NaN()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xoff := 0; xoff <= 1; xoff++ {
				for yoff := 0; yoff <= 1; yoff++ {
					x := xyrange * (float64(i+xoff)/cells - 0.5)
					y := xyrange * (float64(j+yoff)/cells - 0.5)
					z := f(x, y)
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

	zmax, zmin = max, min
}

func Surface3() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z, err := corner3(i+1, j)
			if err != nil {
				continue
			}
			bx, by, z, err := corner3(i, j)
			if err != nil {
				continue
			}
			cx, cy, z, err := corner3(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, z, err := corner3(i+1, j+1)
			if err != nil {
				continue
			}
			fmt.Printf("<polygon style='fill:%s;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				getColor(z, zmin, zmax), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner3(i, j int) (float64, float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f3(x, y)
	if math.IsInf(z, 0) {
		return 0, 0, 0, errors.New("non finite polygon generated")
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

func f3(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
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
