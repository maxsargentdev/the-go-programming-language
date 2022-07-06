package tutorial

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func badEcho() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	total := time.Since(start)
	fmt.Printf("bad echo took %d microseconds \n", total.Microseconds())
}

func goodEcho() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	total := time.Since(start)
	fmt.Printf("good echo took %d microseconds \n", total.Microseconds())
}

func Echo3() {
	badEcho()
	goodEcho()
}
