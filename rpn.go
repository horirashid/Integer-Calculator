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

func CalcAdd(a fraction, b fraction) fraction {
	var res fraction
	var temp int
	res.numerator = a.denominator*b.numerator + a.numerator*b.denominator
	res.denominator = a.denominator * b.denominator
	temp = gcd(res.numerator, res.denominator)
	res.numerator /= temp
	res.denominator /= temp
	return res
}

func CalcMul(a fraction, b fraction) fraction {
	var newN = a.numerator * b.numerator
	var newD = a.denominator * b.denominator
	var z = gcd(newN, newD)
	return fraction{numerator: newN / z, denominator: newD / z}
}

func CalcDiv(a fraction, b fraction) fraction {
	var res fraction
	var c int
	res.numerator = a.numerator * b.denominator
	res.denominator = a.denominator * b.numerator
	c = gcd(res.numerator, res.denominator)
	res.numerator /= c
	res.denominator /= c
	return res
}

func main() {

}
