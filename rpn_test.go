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
