package c1e1

import (
	"fmt"
	"os"
	"strings"
)

func EchoWithCommandName() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
