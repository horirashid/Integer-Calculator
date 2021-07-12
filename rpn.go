package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
func CalcSub(a fraction, b fraction) fraction {
	var res fraction
	var temp int
	res.numerator = a.numerator*b.denominator - a.denominator*b.numerator
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

func String2Fraction(str string) fraction {
	n1, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return fraction{int(n1), 1}
	}
	return fraction{0, 0}
}

func RPN(rpn string) fraction {
	var res fraction
	arr := strings.Split(rpn, " ")
	opt := &Stack{
		maxNum: 10,
		top:    -1,
	}
	a, b := fraction{0, 0}, fraction{0, 0}
	for _, value := range arr {
		nuop := String2Fraction(value)
		temper := fraction{0, 0}
		if nuop == temper {
			a, _ = opt.Pop()
			b, _ = opt.Pop()
			switch value {
			case "+":
				res = CalcAdd(b, a)
			case "-":
				res = CalcSub(b, a)
			case "*":
				res = CalcMul(b, a)
			case "/":
				res = CalcDiv(b, a)
			}
			opt.Push(res)
		} else {
			opt.Push(nuop)
		}
	}
	a, _ = opt.Pop()
	return a
}

func main() {
	var r = bufio.NewScanner(os.Stdin)
	var rpn string
	r.Scan()
	rpn = r.Text()
	res := RPN(rpn)
	if res.denominator == 1 {
		fmt.Printf("%d\n", res.numerator)
	} else {
		fmt.Printf("%d/%d\n", res.numerator, res.denominator)
	}
}
