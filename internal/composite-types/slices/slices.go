package slices

import "fmt"

func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}

func Reverse2(ap *[5]int) { // they probably want us to do an in place swap here instead of just appending....
	result := make([]int, 0) // create result slice

	for i := 4; i >= 0; i-- { // go from right to left of the ap array, appending to result slice
		result = append(result, ap[i])
	}

	fmt.Println(result) // print
}

func ReverseUsingArrayPointer(ap *[5]int) { // they probably want us to do an in place swap here instead of just appending....

	for i, j := 0, len(ap)-1; i < j; i, j = i+1, j-1 {
		ap[i], ap[j] = ap[j], ap[i]
	}
	fmt.Println(ap)
}
