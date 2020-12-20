package main

import "testing"

func TestMySum(t *testing.T) {
	v := mySum(1, 2, 3)
	if v != 6 {
		t.Error("Expected 6, but got", v)
	}

	v = mySum(4, 4)
	if v != 8 {
		t.Error("Expected 8, but got", v)
	}

	v = mySum(5, 95)
	if v != 100 {
		t.Error("Expected 100, but got", v)
	}
}
