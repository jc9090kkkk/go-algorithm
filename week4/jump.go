// 给你一个非负整数数组nums，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
// 假设你总是可以到达数组的最后一个位置。
//
// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
// 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
// 解题思路: 贪心算法，通过局部最优解得到全局最优解，如果进行贪心地正向查找，每次找到可到达的最远位置，就可以在线性时间内得到最少的跳跃次数
//
// 时间复杂度:O(N)
// 空间复杂度:O(1)

package main

func Jump(nums []int) int {
	// 获取数组长度
	length := len(nums)
	end := 0
	// 最大位置
	maxPosition := 0
	// 步数
	steps := 0
	for i := 0; i < length - 1; i++ {
		// 计算最大位置
		maxPosition = Max(maxPosition, i + nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
