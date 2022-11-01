package tutorial

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

//1.1
func EchoWithCommandName() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

//

//1.2
func EchoWithIndexAndValue() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("index[%d] - arg[%s]\n", i, os.Args[i])
	}
}

//

//1.3
func timedBadEcho() {
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

func timedGoodEcho() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	total := time.Since(start)
	fmt.Printf("good echo took %d microseconds \n", total.Microseconds())
}

func TimedEchoTest() {
	timedBadEcho()
	timedGoodEcho()
}

//

//1.4
func Dup2() {

	linemap := make(map[string]map[string]int)

	files := os.Args[2:] // Use two here not one, changes because of cobra cmd

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(file, filename, linemap)
		file.Close()
	}

	for line, filenames := range linemap {

		filecount := len(filenames)

		if filecount == 0 {
			continue
		}
		if filecount == 1 {
			keys := reflect.ValueOf(filenames).MapKeys()
			fmt.Printf("\tFound %s in %s with 1 hit(s)\n", line, keys[0])
		}
		if filecount > 1 {
			fmt.Printf("\tFound %s in %d files:\n", line, filecount)
			for name, count := range filenames {
				fmt.Printf("\t\t%d hit(s) in %s\n", count, name)
			}
		}

	}

}

func countLines(file *os.File, filename string, linemap map[string]map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		if linemap[input.Text()] == nil {
			linemap[input.Text()] = make(map[string]int)
		}
		linemap[input.Text()][filename]++
	}
}

//
