package basicdatatypes

import (
	"errors"
	"fmt"
	"math"
)

func Surface2() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner2(i+1, j)
			if err != nil {
				continue
			}
			bx, by, err := corner2(i, j)
			if err != nil {
				continue
			}
			cx, cy, err := corner2(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, err := corner2(i+1, j+1)
			if err != nil {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner2(i, j int) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f2(x, y)
	if math.IsInf(z, 0) {
		return 0, 0, errors.New("non finite polygon generated")
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f2(x, y float64) float64 {
	r := math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 12 // distance from (0,0)
	return r
}
