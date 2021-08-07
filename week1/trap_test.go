package main

import "testing"

func TestTrap(t *testing.T) {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	expected := 6
	res := Trap(nums)

	if res != expected {
		t.Errorf("test error")
	} else {
		t.Log("test success")
	}
}
