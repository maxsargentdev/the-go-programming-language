package tutorial

import (
	"fmt"
	"os"
	"strings"
)

func Echo1() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
