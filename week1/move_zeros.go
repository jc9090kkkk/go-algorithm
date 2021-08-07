// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序
// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
//
// 解题思路:定义两个指针分别为left和right，其中right指针指用来记录有多少个非0元素，每当碰到非0元素时，将数组往左边挪，

package main

func MoveZeros(nums []int) []int {
	// 定义双指针和数组长度
	left, right, n := 0, 0, len(nums)

	// 右指针的遍历范围小于数组长度
	for right < n {
		// 如果碰到非0元素 则交互左右指针对应的元素 并向右侧移动
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		// 只移动有指针 用于扫描非0元素
		right++
	}

	return nums
}
