package counter

import (
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) { // the int returned here is the number of bytes read!
	byteCount := len(p)
	wordCount := len(strings.Fields(string(p)))
	*c += WordCounter(wordCount) // convert int to WordCounter
	return byteCount, nil
}
