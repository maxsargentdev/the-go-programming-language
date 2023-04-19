package statefulmtsort

import (
	"fmt"
	"time"
)

func RunStatefulMTSort() {
	fmt.Println("I am sorting a table in a multi tier format, whilst retaining some state")
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type multiTierSort struct {
	t []*Track
	tier1
}

func (x multiTierSort) Len() int      { return len(x.t) }
func (x multiTierSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x multiTierSort) Less(i, j int) bool {
	if x.t != y.Title {
		return x.Title < y.Title
	}
	if x.Year != y.Year {
		return x.Year < y.Year
	}
	if x.Length != y.Length {
		return x.Length < y.Length
	}
	return false
}
