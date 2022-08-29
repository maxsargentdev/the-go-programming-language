package basicdatatypes

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

func Serve() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	height := r.URL.Query().Get("height")
	if height != "" {
		height, err := strconv.Atoi(height)
		if err != nil {
			log.Print(err)
			height = 300
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		Surface4(w, height)
	} else {
		w.Header().Set("Content-Type", "image/svg+xml")
		Surface4(w, 300)
	}

}
func Surface4(out io.Writer, height int) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z, err := corner4(i+1, j)
			if err != nil {
				continue
			}
			bx, by, z, err := corner4(i, j)
			if err != nil {
				continue
			}
			cx, cy, z, err := corner4(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, z, err := corner4(i+1, j+1)
			if err != nil {
				continue
			}
			fmt.Fprintf(out, "<polygon style='fill:%s;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				getColor2(z, zmin, zmax), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner4(i, j int) (float64, float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f4(x, y)
	if math.IsInf(z, 0) {
		return 0, 0, 0, errors.New("non finite polygon generated")
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

func f4(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func getColor2(z float64, zmin float64, zmax float64) string { // put zmin and zmax in globals

	if z > 0 {
		gradient := math.Exp(math.Abs(z)) / math.Exp(math.Abs(zmax)) * 255
		return fmt.Sprintf("#%02x0000", int(math.Round(gradient)))
	} else {
		gradient := math.Exp(math.Abs(z)) / math.Exp(math.Abs(zmin)) * 255
		return fmt.Sprintf("#0000%02x", int(math.Round(gradient)))
	}

}
