package counter

import "strings"

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) { // the int returned here is the number of bytes read!
	byteCount := len(p)
	lineCount := len(strings.Split(strings.ReplaceAll(string(p), "\r\n", "\n"), "\n"))
	*c += LineCounter(lineCount) // convert int to WordCounter
	return byteCount, nil
}
