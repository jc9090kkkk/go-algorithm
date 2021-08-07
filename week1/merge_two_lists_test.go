package main

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var (
		l1 *ListNode
		l2 *ListNode
	)

	s1 := []int{1, 2, 4}
	s2 := []int{1, 3, 4}
	expected := []int{1, 1, 2, 3, 4, 4}

	// 创建链表
	l1 = MakeListNode(s1)
	l2 = MakeListNode(s2)

	// 合并链表 并获取链表元素
	res := MergeTwoLists(l1, l2).GetListNodes()

	// 判断与否与预期结果一致
	isSuccess := true
	for i,_ := range res {
		if expected[i] != res[i] {
			isSuccess = false
		}
	}

	if !isSuccess {
		t.Errorf("test error")
	} else {
		t.Log("test success")
	}
}
