package mandelbrot

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	defaultZoom string = "medium"
	defaultY           = 1
	defaultX           = 1
)

func Serve() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	var x float64 = 1
	var y float64 = 1
	var zoom string = "medium"

	zoomQueryParam := r.URL.Query().Get("zoom")
	if zoomQueryParam != "" {
		zoom = zoomQueryParam
	}
	fmt.Println(zoom)

	xQueryParam := r.URL.Query().Get("x")
	if xQueryParam != "" {
		x, _ = strconv.ParseFloat(xQueryParam, 64)
	}

	yQueryParam := r.URL.Query().Get("y")
	if yQueryParam != "" {
		y, _ = strconv.ParseFloat(yQueryParam, 64)
	}

	renderTypeWithParams(w, x, y, zoom)
}
