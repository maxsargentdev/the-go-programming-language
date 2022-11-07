package tutorial

import (
	"log"
	"net/http"
	"strconv"
)

const (
	defaultCycles int = 20
)

func Serve() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	cycles := r.URL.Query().Get("cycles")
	if cycles != "" {
		cycles, err := strconv.Atoi(cycles)
		if err != nil {
			log.Print(err)
			cycles = defaultCycles
		}
		CyclesArgMultiColouredLissajous(w, cycles)
	} else {
		CyclesArgMultiColouredLissajous(w, defaultCycles)
	}
}
