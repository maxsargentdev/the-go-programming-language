package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

func RunOutline(inputString string) {
	reader := NewReader(inputString)

	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

// we need to use this function to change the above RunOutline function so that it takes input from a string
// rather than a file (os.Stdin)

type Reader struct {
	s string
	i int64
}

func NewReader(s string) *Reader {
	return &Reader{s: s, i: 0}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	// handle zero length string case
	if len(p) == 0 {
		return 0, nil
	}

	// jump into the cursors location
	n = copy(p, r.s[r.i:])

	// start reading to the end of file
	if r.i += int64(n); r.i >= int64(len(r.s)) {
		err = io.EOF
	}
	return
}

type LimitedReader struct {
	R io.Reader
	N int64
}

func (lr *LimitedReader) Read(p []byte) (n int, err error) {

	// if the 0 or less requested to be read, just return 0,EOF
	if lr.N <= 0 {
		return 0, io.EOF
	}

	// if the length of the byte slice is greater than the limit, just read what we can
	if int64(len(p)) > lr.N {
		p = p[0:lr.N]
	}

	// read the bytes
	n, err = lr.R.Read(p)

	// update our limit so if we read again its not over
	lr.N -= int64(n)

	return

}

func LimitReader(r io.Reader, n int64) io.Reader {
	limitedReader := &LimitedReader{
		R: r,
		N: n,
	}
	return limitedReader
}
