package main

import (
	"testing"
)

func TestLemonadeChange(t *testing.T) {
	bills := []int{5, 5, 5, 10, 20}
	expected := true

	res := LemonadeChange(bills)

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
}