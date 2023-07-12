package c1e2

import (
	"fmt"
	"os"
)

func EchoWithIndexAndArguments() {
	for i := 0; i < len(os.Args); i++ {
		index := i
		argument := os.Args[i]
		fmt.Printf("%d - %s\n", index, argument)
	}
}
