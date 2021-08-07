package main

import (
	"testing"
)

func TestMerTwoSortedArr(t *testing.T) {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	expected := []int{1, 2, 2, 3, 5, 6}

	res := MerTwoSortedArr(a, 3, b, 3)

	// 判断与否与预期结果一致
	isSuccess := true
	for i,_ := range res {
		if expected[i] != res[i] {
			isSuccess = false
		}
	}

	if !isSuccess {
		t.Errorf("test error")
	} else {
		t.Log("test success")
	}
}
