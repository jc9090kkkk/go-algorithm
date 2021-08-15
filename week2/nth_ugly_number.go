// 我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
// 输入: n = 10
// 输出: 12
// 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
//
// 解题思路:要得到从小到大的第n个丑数，可以用最小堆来实现。
// 初始时堆为空。首先将最小的丑数 1 加入堆。
// 每次取出堆顶元素 x，则 x 是堆中最小的丑数，由于 2x, 3x, 5x也是丑数，因此将 2x, 3x, 5x 加入堆。
// 为了避免重复元素，可以使用哈希集合去重，避免相同元素多次加入堆。在排除重复元素的情况下，第 n 次从最小堆中取出的元素即为第 n 个丑数。
package main

import (
	"container/heap"
	"sort"
)

// 定义丑数因子
var factors = []int{2, 3, 5}

// 定义最小堆
type hp struct{ sort.IntSlice }
// Push 压入操作
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
// Pop 弹出操作
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func NthUglyNumber(n int) int {
	h := &hp{sort.IntSlice{1}}
	seen := map[int]struct{}{1: {}}
	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		if i == n {
			return x
		}
		for _, f := range factors {
			next := x * f
			if _, has := seen[next]; !has {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}
}