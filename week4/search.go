// 整数数组 nums 按升序排列，数组中的值 互不相同 。
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//
// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4
//
// 解题思路:
// 二分查找，虽然常规的二分查找只适用于有序集合或者数组
// 但是对于二分查找的思想是分治的思想，只要一组元素有相似的规律或者对立的规则，根据二分查询都可以通过边界条件将数据分而治之，最后达到目的。
// 最关键是middle和左右边界的判断如何确定，使其最终落在想要的范围内
//
// 时间复杂度:O(log N)
// 空间复杂度:O(1)

package main

func Search(nums []int, target int) int {
	// 定义左右边界
	left, right := 0, len(nums) - 1

	// 左边界不能大于右边界 不能越界
	for left <= right {
		// 计算中间值
		mid := (right - left) / 2 + left
		// 如果能直接匹配到目标值 直接返回
		if nums[mid] == target {
			return mid
		}
		// 左边界的值小于中间值
		if nums[mid] >= nums[left] {
			// 中间值大于目标值 并且 目标值大于等于左边界的值 右边界变成中间值-1的位置 否则变成中间值+1的位置
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 中间值小于目标值 并且 目标值小于等于左边界的值 左边界变成中间值+1的位置 否则变成中间值-1的位置
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
