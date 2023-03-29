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
