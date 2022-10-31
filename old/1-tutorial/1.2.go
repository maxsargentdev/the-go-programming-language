package tutorial

import (
	"fmt"
	"os"
	"strings"
)

func Echo2() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
