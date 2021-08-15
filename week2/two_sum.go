// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
// 示例 1：
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
// 示例 2：
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
//
// 示例 3：
// 输入：nums = [3,3], target = 6
// 输出：[0,1]

// 解题思路：
// 如果用暴力解法，那就是使用双重循环，把所有数字的加和结果全部计算出来，不过这并不是我们想要的解题方式，因为效率太差，时间复杂度为O(N^2)，
// 更好的方式是构造一个hash表，将nums数组中的元素作为key，索引作为value，利用一个for循环，就可以判断出来，这是利用了hash表的键值唯一性
// 时间复杂度: O(N)
// 空间复杂度: O(N)

package main

func TwoSum(nums []int, target int) []int {
	// 定义一个hashTable
	hashTable := map[int]int{}

	for i,p := range nums {
		// 判断当前元素是否存在hashTable中
		e, ok := hashTable[target - p]
		// 如果存在 则说明存在两数之和 直接返回下标
		if ok {
			return []int{e, i}
		}

		// 否则将当前元素的索引值加入到hashTable中 用于后续判断
		hashTable[p] = i
	}

	return nil
}