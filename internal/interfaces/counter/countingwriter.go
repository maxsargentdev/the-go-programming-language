package counter

import "io"

type CWriter struct {
	Counter int64     // somewhere to keep track of the bytes
	Writer  io.Writer // wrap a writer
}

func (cw *CWriter) Write(p []byte) (int, error) {
	cw.Counter += int64(len(p)) // whenever we write, track the bytes used
	return cw.Writer.Write(p)   // call the write method of the underlying writer
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CWriter{0, w} // wrap w, initialize 0 counter
	return &cw, &cw.Counter
}
