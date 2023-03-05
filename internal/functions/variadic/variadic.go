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

func Join(sep string, strings ...string) string {

	if sep == "" {
		fmt.Println("no input provided")
		return ""
	}
	if len(strings) == 0 {
		fmt.Println("no input provided")
		return ""
	}

	var returnMe string
	for i, v := range strings {
		if i > 0 {
			returnMe = fmt.Sprintf("%s%s%s", returnMe, sep, v)
		} else {
			returnMe = fmt.Sprintf("%s%s", returnMe, v)
		}
	}
	fmt.Println(returnMe)
	return returnMe
}
