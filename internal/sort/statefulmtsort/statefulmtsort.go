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

type Pair[T, U any] struct {
	First  T
	Second U
}

type multiTierSort struct {
	t            []*Track
	primaryKey   string
	secondaryKey string
	tertiaryKey  string
}

func (x multiTierSort) Len() int      { return len(x.t) }
func (x multiTierSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x multiTierSort) Less(i, j int) bool {
	if x.t[i].Title != x.t[j].Title {
		return x.t[i].Title < x.t[j].Title
	}
	if x.t[i].Year != x.t[j].Year {
		return x.t[i].Year < x.t[j].Year
	}
	if x.t[i].Length != x.t[j].Length {
		return x.t[i].Length < x.t[j].Length
	}
	return false
}
