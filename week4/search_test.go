package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{4,5,6,7,0,1,2}
	target := 0
	expected := 4

	res := Search(nums, target)

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

	target = 3
	expected = -1

	res = Search(nums, target)

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

	nums = []int{1}
	target = 0
	expected = -1

	res = Search(nums, target)

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