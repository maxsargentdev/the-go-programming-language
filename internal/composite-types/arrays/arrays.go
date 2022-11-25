package arrays

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func SHABitCompare(stringA, stringB string) {
	c1 := sha256.Sum256([]byte(stringA))
	c2 := sha256.Sum256([]byte(stringB))

	diffBitCount := 0
	for i := 0; i < len(c2); i++ {
		diffBitCount += diffByte(c1[i], c2[i])
	}

	fmt.Println(diffBitCount)
}

func diffByte(b1, b2 byte) int {
	count := 0

	for i := uint(0); i < 8; i++ {
		// fmt.Printf("%b-%b\n", (b1 >> i), (b2 >> i))
		// fmt.Printf("%d-%d\n", (b1 >> i), (b2 >> i))
		bit1 := (b1 >> i) & 1 // you need to consider this a binary number, it is 00000001
		bit2 := (b2 >> i) & 1 // if this ands with 00000001 then it means that it ts first digit is 1, or its odd
		if bit1 != bit2 {
			fmt.Printf("%d-%d\n", (b1 >> i), (b2 >> i)) // these should all be pairs of numbers that are opposite odd and even
			count++
		}
		// } else {
		// 	fmt.Printf("%d-%d\n", (b1 >> i), (b2 >> i)) // these should all be pairs of numbers that are matching odds & even

		// }
	}
	return count
}

const (
	SHA256 string = "SHA256"
	SHA384 string = "SHA312"
	SHA512 string = "SHA512"
)

func SHAOutput(input string, shaSize string) {
	switch shaSize {
	case SHA256:
		h := sha256.New()
		h.Write([]byte(input))
		hash := hex.EncodeToString(h.Sum(nil))
		fmt.Println(hash)
	case SHA384:
		h := sha512.New384()
		h.Write([]byte(input))
		hash := hex.EncodeToString(h.Sum(nil))
		fmt.Println(hash)
	case SHA512:
		h := sha512.New()
		h.Write([]byte(input))
		hash := hex.EncodeToString(h.Sum(nil))
		fmt.Println(hash)
	default:
		h := sha256.New()
		h.Write([]byte(input))
		hash := hex.EncodeToString(h.Sum(nil))
		fmt.Println(hash)
	}
}
