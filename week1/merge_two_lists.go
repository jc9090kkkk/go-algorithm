// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
// 输入：l1 = [], l2 = [0]
// 输出：[0]
// 解题思路：拿到两个链表里所有的值，然后直接排序，用一个新的链表存一下并返回即可
// 通过分别对两个链表进行比较，比较时，值小的那个被存进链表，然后该链表向后一格再做比较，直至某一条链表到达尾部
// 时间复杂度: O(M+N)
// 空间复杂度: O(M+N)

package main

type ListNode struct {
	Val int
	Next *ListNode
}

// MakeListNode 生成链表
func MakeListNode(s []int) *ListNode {
	if len(s) == 0 {
		return nil
	}
	node := &ListNode{
		Val: s[0],
	}
	tmp := node
	for i := 1; i < len(s); i++ {
		tmp.Next = &ListNode{
			Val: s[i],
		}
		tmp = tmp.Next
	}
	return node
}

// MergeTwoLists 合并两个链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{Val:0}
	res := head
	for {
		if l1 == nil{
			res.Next = l2
			break
		}

		if l2 == nil {
			res.Next = l1
			break
		}

		if l1 != nil && l2 != nil {
			if l1.Val>l2.Val {
				test := ListNode{Val : l2.Val}
				res.Next = &test
				l2 = l2.Next
			} else {
				res.Next = &ListNode{Val : l1.Val}
				l1 = l1.Next
			}
			res = res.Next
		}
	}

	return head.Next
}

// GetListNodes 获取链表节点元素
func (list *ListNode) GetListNodes() []int {
	var s []int
	ptr := list
	for ptr != nil {
		s = append(s, ptr.Val)
		ptr = ptr.Next
	}

	var listNodesElements []int
	for _, v := range s {
		listNodesElements = append(listNodesElements, v)
	}

	return listNodesElements
}