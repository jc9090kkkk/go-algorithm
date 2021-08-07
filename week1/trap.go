// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解题思路: 1.动态规划 2.栈

package main

// Max 计算最大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min 计算最小值
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Trap 计算接雨水
func Trap(height []int) int {
	length := len(height)

	left := make([]int, length)
	right := make([]int, length)

	for i := 1; i < length; i++ {
		left[i] = Max(left[i-1], height[i-1])
	}

	for i := length - 2; i >= 0; i-- {
		right[i] = Max(right[i+1], height[i+1])
	}

	res := 0
	for i := 0; i < length; i++ {
		hwm := Min(left[i], right[i])
		res += Max(0, hwm-height[i])
	}

	return res
}