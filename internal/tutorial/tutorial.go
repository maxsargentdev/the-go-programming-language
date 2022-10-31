package tutorial

import (
	"fmt"
	"os"
	"strings"
)

func Echo1() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func Echo2() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
