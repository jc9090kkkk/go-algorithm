// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
// 使用空间复杂度为 O(1) 的 原地 算法解决这个问题
//
// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]

// 解题思路: 首先对整个数组进行翻转，然后按照k的位置，将整个数组分割成2个子部分，然后分别将其再进行翻转，最后拼接在一起即可
// 需要额外再定一个用于翻转数组的函数来提供辅助

package main

func Rotate(nums []int, k int) []int {
	// 求出分割位置
	k %= len(nums)

	// 翻转整个数组
	reverse(nums)

	// 翻转0-k位置的数组元素
	reverse(nums[:k])
	// 翻转k-len位置的数组元素
	reverse(nums[k:])

	return nums
}

// reverse 数组翻转
func reverse(arr []int) {
	length := len(arr)
	// 计算出中间位置后 将中间位置左右两侧的元素俩俩替换
	for i := 0; i < length/2; i++ {
		temp := arr[length-1-i]
		arr[length-1-i] = arr[i]
		arr[i] = temp
	}
}
