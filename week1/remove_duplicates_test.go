package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 3, 3, 4, 5, 6, 6}
	if res := RemoveDuplicates(nums); res != 6 {
		t.Errorf("%d expected be 6, but %d got", nums, res)
	} else {
		t.Log("test success")
	}
}
