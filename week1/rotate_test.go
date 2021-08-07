package main

import "testing"

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 5, 7, 9, 8}
	expectedArr := []int{7, 9, 8, 1, 2, 3, 5}

	// 翻转
	res := Rotate(nums, 3)

	// 比较翻转后的数组是否符合预期结果
	isSuccess := true
	for i, _ := range res {
		if res[i] != expectedArr[i] {
			isSuccess = false
			break
		}
	}
	if !isSuccess {
		t.Errorf("test error")
	} else {
		t.Log("test success")
	}
}
