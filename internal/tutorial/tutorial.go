package tutorial

import (
	"fmt"
	"os"
	"strings"
)

func EchoWithCommandName() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func EchoWithIndexAndValue() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("index[%d] - arg[%s]\n", i, os.Args[i])
	}
}
