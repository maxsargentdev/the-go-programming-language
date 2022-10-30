package anagram

import (
	"fmt"
	"sort"
	"strings"
)

type runeSlice []rune

func Anagram(stringOne string, stringTwo string) {
	fmt.Println(stringToRuneList(stringOne, stringTwo))

}

func stringToRuneList(inputString1 string, inputString2 string) bool {
	len1 := len(inputString1)
	len2 := len(inputString2)

	if len1 != len2 {
		return false
	}

	r1 := []byte(inputString1)
	r2 := []byte(inputString2)

	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})
	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})

	for i := 0; i < len1; i++ {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
	// runeList := []rune
	// for _, v := range inputString {
	// 	fmt.Println(string(v))
	// }
}

func stringToRuneList2(inputString1 string, inputString2 string) bool {
	aa := strings.Split(inputString1, "")
	sort.Strings(aa)
	aaa := strings.Join(aa, "")

	bb := strings.Split(inputString2, "")
	sort.Strings(bb)
	bbb := strings.Join(bb, "")

	return aaa == bbb
}
