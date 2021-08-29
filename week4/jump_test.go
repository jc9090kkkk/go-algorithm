package main

import (
	"testing"
)

func TestJump(t *testing.T) {
	nums := []int{2,3,1,1,4}
	expected := 2

	res := Jump(nums)

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

	nums = []int{2,3,0,1,4}
	expected = 2

	res = Jump(nums)

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