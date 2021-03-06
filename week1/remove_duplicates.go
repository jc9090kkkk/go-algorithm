// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// 输入：nums = [1,1,2]
// 输出：2, nums = [1,2]
//
// 输入：nums = [0,0,1,1,1,2,2,3,3,4]
// 输出：5, nums = [0,1,2,3,4]
// 解题思路:
// 比较 slow 和 fast 位置的元素是否相等。
//
// 如果相等，fast 后移 1 位
// 如果不相等，将 fast 位置的元素复制到 slow+1 位置上，slow 后移一位，fast 后移 1 位，重复上述过程，直到 fast 等于数组长度。
// 返回 slow + 1，即为新数组长度
// 时间复杂度: O(N)
// 空间复杂度: O(1)

package main

func RemoveDuplicates(nums []int) int {
	// 获取数组长度
	length := len(nums)
	// 数组为0直接
	if length == 0 {
		return 0
	}

	// 设置一快一慢的两个指针，设置起始位置为1
	slow := 1
	// 遍历数组 通过快慢指针来判断对应的元素是否相等
	for fast := 1; fast < length; fast++ {
		// 如果元素不相等 快指针赋值给慢指针 慢指针后移等 否则快指针后移 直到快指针循环越界
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}