// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
// 你可以假设除了整数 0 之外，这个整数不会以零开头。
//
// 解题思路: 先从末位数开始计算，如果没有进位，则可以直接返回，如果有进位，则继续循环计算前一位数，以此类推，如果计算到最后发现需要整体进位，
// 则构造一个头部的1，然后把剩余部分的数组追加在尾部即可
// 时间复杂度: O(N)
// 空间复杂度: O(1)

package main

import "fmt"

// PlusOne 加1
func PlusOne(digits []int) []int {
	length := len(digits)

	for i := length - 1; i >= 0; i-- {
		// 当前元素+1 当表元素的后位数有进位 需要提前计算
		digits[i]++
		// 计算是否为10的余数 用于判断是否进位
		digits[i] %= 10
		// 没有进位 直接返回
		if digits[i] != 0 {
			return digits
		}
	}

	// 如果有进位
	newDigits := []int{1}
	// 将新增的1补在头部
	newDigits = append(newDigits, digits...)

	return newDigits
}

func main() {
	nums := []int{9, 9, 6, 9}

	res := PlusOne(nums)

	fmt.Println(res)
}
