package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	bills := []int{7,1,5,3,6,4}
	expected := 7

	res := MaxProfit(bills)

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