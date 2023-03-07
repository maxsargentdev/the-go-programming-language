package panicandrecover

import "fmt"

func PanicAndRecover() (returnMe string) {
	defer func() {
		panicValue := recover()
		returnMe = fmt.Sprintf("I am panicking, %s", panicValue)
	}()

	panic("HOUSE IS ON FIRE!")
}
