package main

import "testing"

func TestPlusOne(t *testing.T) {
	nums := []int{1, 2, 9, 9}
	expected := []int{1, 3, 0, 0}

	res := PlusOne(nums)

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
