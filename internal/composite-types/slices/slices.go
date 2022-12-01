package slices

import "fmt"

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
	// reverse(s[:2])
	// reverse(s[2:])
	// reverse(s)
	// fmt.Println(s) // "[2 3 ` 5 0 1]"
	// for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	// 	s[i], s[j] = s[j], s[i]
	// }
	// create empty slice
	var result []int

	// take 0 to position count, i.e the first 2 (1,2)
	chunk := s[:positionCount]

	// append to the empty slice the end, i.e the last 3 (3,4,5)
	result = append(result, s[positionCount:]...)
	// append to the (3,4,5) slice, gives you (3,4,5,1,2)
	result = append(result, chunk...)

	fmt.Println(result)
}

func RemoveAdjacentDups(s []string) {
	slength := len(s)
	// out := s[:0]
	for i := 0; i < len(s)-1; i++ {
		backwards := i - 1
		frontwards := i + 1

		if backwards > 0 {
			if s[i] == s[backwards] {
				s = remove(s, backwards)
			}
		}
		if frontwards < slength {
			if s[i] == s[frontwards] {
				s = remove(s, frontwards)
			}
		}

	}

	fmt.Println(s)
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
