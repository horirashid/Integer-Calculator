package main

import (
//"bufio"
//"fmt"
//"os"
//"strconv"
//"strings"
)

type fraction struct {
	numerator   int
	denominator int
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func gcd(x, y int) int {
	if x == 0 || y == 0 {
		return 1
	}
	x = abs(x)
	y = abs(y)
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}

func main() {
}
