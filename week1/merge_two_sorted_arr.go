// 给你两个有序整数数组nums1 和 nums2，请你将nums2合并到nums1中，使nums1成为一个有序数组。
//
// 初始化nums1和nums2的元素数量分别为m和n 。你可以假设nums1的空间大小等于m + n，这样它就有足够的空间保存来自nums2的元素。
//
// 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// 输出：[1,2,2,3,5,6]

// 解题思路:

package main

func MerTwoSortedArr(nums1 []int, m int, nums2 []int, n int) []int {

	for m > 0 || n > 0 {

		if n == 0 {
			break
		}

		if m == 0 {
			nums1[n-1] = nums2[n-1]
			n--
			continue
		}

		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	return nums1
}
