package main

import "testing"

func Test01(t *testing.T) {
	res := gcd(4, 6)
	if res != 2 {
		t.Errorf("test case 01 failed, output %d", res)
	}
}

func Test02(t *testing.T) {
	res := gcd(12, 8)
	if res != 4 {
		t.Errorf("test case 02 failed, output %d", res)
	}
}

func Test03(t *testing.T) {
	res := gcd(-12, 8)
	if res != 4 {
		t.Errorf("test case 03 failed, output %d", res)
	}
}

func Test04(t *testing.T) {
	a := fraction{1, 3}
	b := fraction{1, 6}
	c := fraction{1, 2}
	res := CalcAdd(a, b)
	if res != c {
		t.Errorf("test case 04 failed, output %d/%d", res.numerator, res.denominator)
	}
}

func Test05(t *testing.T) {
	a := fraction{1, 4}
	b := fraction{1, 2}
	c := fraction{1, 8}
	res := CalcMul(a, b)
	if res != c {
		t.Errorf("test case 05 failed, output %d/%d", res.numerator, res.denominator)
	}
}

func Test06(t *testing.T) {
	a := fraction{1, 8}
	b := fraction{1, 2}
	c := fraction{1, 4}
	res := CalcDiv(a, b)
	if res != c {
		t.Errorf("test case 06 failed, output %d/%d", res.numerator, res.denominator)
	}
}

func Test08(t *testing.T) {
	s := &Stack{
		maxNum: 10,
		top:    -1,
	}
	a := fraction{1, 8}
	s.Push(a)
	b, _ := s.Pop()
	if a != b {
		t.Errorf("test case 08 failed, output %d/%d", b.numerator, b.denominator)
	}
}
