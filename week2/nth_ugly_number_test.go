package main

import "testing"

func TestNthUglyNumber(t *testing.T) {
	nums := 10
	expected := 12

	res := NthUglyNumber(nums)
	if res != expected {
		t.Errorf("test error")
	} else {
		t.Log("test success")
	}
}
