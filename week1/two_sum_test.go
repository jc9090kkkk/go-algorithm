package main

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 11, 7, 15}
	target := 9
	expected := []int{0, 2}

	res := TwoSum(nums, target)

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
