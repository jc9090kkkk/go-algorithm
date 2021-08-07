package main

import "testing"

func TestMoveZeros(t *testing.T) {
	nums := []int{1, 9, 0, 4, 6, 0, 12, 16, 0}
	expected := []int{1, 9, 4, 6, 12, 16, 0, 0 ,0}

	res := MoveZeros(nums)

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
