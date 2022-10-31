package comma

import (
	"bytes"
	"fmt"
	"strings"
)

func CommaFloatingP(inputString string) {
	const chunkSize = 3
	var buf bytes.Buffer
	var sLength = len(inputString)
	var suffix = ""
	var firstCharacter = inputString[0:1]
	var signed = strings.ContainsAny(firstCharacter, "+-")
	var floatingPoint = strings.ContainsAny(inputString, ".")
	var floatingPointIndex = strings.Index(inputString, ".")
	var chunkCount = sLength / chunkSize
	var remainderCount = sLength % chunkSize

	if signed {
		buf.WriteString(firstCharacter)
		inputString = inputString[1:]
		sLength--
	}

	if floatingPoint {
		suffix = inputString[floatingPointIndex-1:]
		inputString = inputString[:sLength-len(suffix)]
		sLength -= len(suffix)
	}

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

	if floatingPoint {
		buf.WriteString(suffix)
	}

	fmt.Println(buf.String())
}
