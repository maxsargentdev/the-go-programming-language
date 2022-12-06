package slices

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func ReverseUsingArrayPointer(ap *[5]int) { // they probably want us to do an in place swap here instead of just appending....

	for i, j := 0, len(ap)-1; i < j; i, j = i+1, j-1 {
		ap[i], ap[j] = ap[j], ap[i]
	}
	fmt.Println(ap)
}

func RotateLeft(positionCount int, s []int) {
	// create empty slice
	var result []int

	// take 0 to position count, i.e. the first 2 (1,2)
	chunk := s[:positionCount]

	// append to the empty slice the end, i.e. the last 3 (3,4,5)
	result = append(result, s[positionCount:]...)

	// append to the (3,4,5) slice, gives you (3,4,5,1,2)
	result = append(result, chunk...)

	fmt.Println(result)
}

func RemoveAdjacentDupes(s []string) {
	length := len(s)
	for i := 0; i < len(s)-1; i++ {
		backwards, frontwards := i-1, i+1

		if backwards > 0 {
			if s[i] == s[backwards] {
				s = remove(s, backwards)
			}
		}
		if frontwards < length {
			if s[i] == s[frontwards] {
				s = remove(s, frontwards)
			}
		}

	}

	fmt.Println(s)
}

func remove(slice []string, x int) []string {
	copy(slice[x:], slice[x+1:])
	return slice[:len(slice)-1]
}

func removeByte(slice []byte, x int) []byte {
	copy(slice[x:], slice[x+1:])
	return slice[:len(slice)-1]
}

func SquashAdjacentUnicodeSpaces(s []byte) {
	for pos, char := range s {
		if unicode.IsSpace(rune(char)) { // 32 is the int that corresponds to whitespace
			for {
				frontwards := pos + 1
				if unicode.IsSpace(rune(s[frontwards])) {
					s = removeByte(s, frontwards)
				} else {
					break
				}

			}
		}
	}
	fmt.Println(string(s))
}

func ReverseByteSlice(in []byte) {
	// first treat as non utf8-encoded data
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}

	// try to decode according to utf8, then fix error
	i := 0
	for i < len(in) {
		var tryTwo, tryThree, tryFour bool
		for {
			r, s := utf8.DecodeRune(in[i:])
			if r != utf8.RuneError {
				i += s
				break
			} else {
				// try two byte length, swap two bytes
				if !tryTwo {
					tryTwo = true
					in[i], in[i+1] = in[i+1], in[i]
					continue
				}

				// try three byte length, swap three bytes
				if !tryThree {
					// cancel tryTwo side effect
					in[i], in[i+1] = in[i+1], in[i]
					tryThree = true
					in[i], in[i+2] = in[i+2], in[i]
					continue
				}

				// try four byte length, swap four bytes
				if !tryFour {
					// cancel tryThree side effect
					in[i], in[i+1], in[i+2] = in[i+2], in[i+1], in[i]

					tryFour = true
					in[i], in[i+1], in[i+2], in[i+3] = in[i+3], in[i+2], in[i+1], in[i]
					continue
				}

				// should not be here
				panic("Should not be here!")
			}
		}
	}
	fmt.Printf("%s", in)

}
