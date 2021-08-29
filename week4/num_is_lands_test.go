package main

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := [][]byte{{'1','1','1','1','0'}, {'1','1','0','1','0'}, {'1','1','0','0','0'}, {'0','0','0','0','0'}}
	expected := 1

	res := NumIslands(grid)

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