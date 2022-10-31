package comma

import (
	"bytes"
	"fmt"
)

func Comma(inputString string) {
	var buf bytes.Buffer
	var sLength = len(inputString)

	const chunkSize = 3

	if sLength <= chunkSize {
		fmt.Println(inputString)
		return
	}

	chunkCount := sLength / chunkSize
	remainderCount := sLength % chunkSize

	for i := 0; i < chunkCount; i++ {

		// Run this on every iteration
		frontIndex := i * chunkSize
		rearIndex := frontIndex + chunkSize
		completedLoop := (i == chunkCount-1)

		buf.WriteString(string(inputString[frontIndex:rearIndex])) // write chunks of 3

		// Break out of last iteration
		if completedLoop {
			break
		}

		// Dont run this on the last iteration
		buf.WriteByte(',')

	}

	if remainderCount > 0 {
		frontIndex := chunkCount * chunkSize
		rearIndex := sLength
		buf.WriteString(string(inputString[frontIndex:rearIndex])) // write the remaining parts
	}

	fmt.Println(buf.String())
	return
}
