package variadic

import "fmt"

func Max(integers ...int) {

	if len(integers) == 0 {
		fmt.Println("no input provided")
		return
	}

	max := integers[0]
	for _, v := range integers {
		if v > max {
			max = v
		}
	}
	fmt.Printf("%d\n", max)
}

func Min(integers ...int) {

	if len(integers) == 0 {
		fmt.Println("no input provided")
		return
	}

	min := integers[0]
	for _, v := range integers {
		if v < min {
			min = v
		}
	}
	fmt.Printf("%d\n", min)
}
