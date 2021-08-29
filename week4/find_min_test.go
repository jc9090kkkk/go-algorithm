package main

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	nums := []int{3,4,5,1,2}
	expected := 1

	res := FindMin(nums)

	// 判断与否与预期结果一致
	isSuccess := true
	if res != expected {
		isSuccess = false
	}

	if !isSuccess {
		t.Errorf("test error")
	} else {
		t.Log("test success")
	}

	nums = []int{4,5,6,7,0,1,2}
	expected = 0

	res = FindMin(nums)

	// 判断与否与预期结果一致
	isSuccess = true
	if res != expected {
		isSuccess = false
	}

	if !isSuccess {
		t.Errorf("test error")
	} else {
		t.Log("test success")
	}

	nums = []int{11,13,15,17}
	expected = 11

	res = FindMin(nums)

	// 判断与否与预期结果一致
	isSuccess = true
	if res != expected {
		isSuccess = false
	}

	if !isSuccess {
		t.Errorf("test error")
	} else {
		t.Log("test success")
	}
}