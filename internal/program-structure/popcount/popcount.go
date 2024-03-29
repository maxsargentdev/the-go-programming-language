package popcount

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
} // pc is a table of results for popcount

func PopCount(x uint64) int { // 8 bytes in a uint64
	start := time.Now()
	popcount := int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
	fmt.Printf("PopCount - %d nanoseconds elapsed\n", time.Since(start).Nanoseconds())
	return popcount
}

func LoopPopCount(x uint64) int {
	start := time.Now()
	popcount := 0
	for i := 0; i < 8; i++ {
		popcount += int(pc[byte(x>>(i*8))])
	}
	fmt.Printf("LoopPopCount - %d nanoseconds elapsed\n", time.Since(start).Nanoseconds())
	return popcount
}

func BitShiftPopCount(x uint64) int {
	start := time.Now()
	// output := ""
	popcount := 0
	for i := 0; i < 64; i++ {
		if x&1 != 0 { // this is a bitwise AND operatian against 1 in binary
			popcount++ // it is equivalent to usig a bit mask!
		}
		// output += fmt.Sprint(x & 1) // same again here
		x = x >> 1
	}
	// fmt.Println(output)
	fmt.Printf("BitShiftPopcount - %d nanoseconds elapsed\n", time.Since(start).Nanoseconds())
	return popcount
}

// Bitwise AND of x & x-1
// Zeros out right most bit, see output below
func BitMaskPopCount(x uint64) int {
	start := time.Now()
	//output := ""
	popcount := 0
	for x != 0 {
		// fmt.Printf("x          = %b\n", x)
		x = x & (x - 1) // AND x with x-1, x-1 will always not mask with x and drop a bit
		//output += fmt.Sprint(x) // output here will be strange as not poaching the last bit anymore
		// it will just be a concatenation of numbers
		// fmt.Printf("x & (x -1) = %b\n", x)
		// fmt.Println("--------------------------------------------------------------")
		popcount++
	} // you essentially are just removing bits untul you get to 0
	fmt.Printf("BitMaskPopcount - %d nanoseconds elapsed\n", time.Since(start).Nanoseconds())
	return popcount
}

// x          = 111011100111000101000010100001
// x & (x -1) = 111011100111000101000010100000
// --------------------------------------------------------------
// x          = 111011100111000101000010100000
// x & (x -1) = 111011100111000101000010000000
// --------------------------------------------------------------
// x          = 111011100111000101000010000000
// x & (x -1) = 111011100111000101000000000000
// --------------------------------------------------------------
// x          = 111011100111000101000000000000
// x & (x -1) = 111011100111000100000000000000
// --------------------------------------------------------------
// x          = 111011100111000100000000000000
// x & (x -1) = 111011100111000000000000000000
// --------------------------------------------------------------
// x          = 111011100111000000000000000000
// x & (x -1) = 111011100110000000000000000000
// --------------------------------------------------------------
// x          = 111011100110000000000000000000
// x & (x -1) = 111011100100000000000000000000
// --------------------------------------------------------------
// x          = 111011100100000000000000000000
// x & (x -1) = 111011100000000000000000000000
// --------------------------------------------------------------
// x          = 111011100000000000000000000000
// x & (x -1) = 111011000000000000000000000000
// --------------------------------------------------------------
// x          = 111011000000000000000000000000
// x & (x -1) = 111010000000000000000000000000
// --------------------------------------------------------------
// x          = 111010000000000000000000000000
// x & (x -1) = 111000000000000000000000000000
// --------------------------------------------------------------
// x          = 111000000000000000000000000000
// x & (x -1) = 110000000000000000000000000000
// --------------------------------------------------------------
// x          = 110000000000000000000000000000
// x & (x -1) = 100000000000000000000000000000
// --------------------------------------------------------------
// x          = 100000000000000000000000000000
// x & (x -1) = 0
// --------------------------------------------------------------
