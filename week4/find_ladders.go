package main

import "math"

func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	// 计算wordList里面单词的个数
	n := len(wordList)
	// ids就是建立从给定word到它所在wordList 当前声明的是索引的map映射
	ids := map[string]int{}

	// 把元素压入到map中
	for i, s := range wordList {
		ids[s] = i
	}

	// 定义一下结果的基本结构
	ret := [][]string{}
	// 要是wordList里面没有endWord，直接返回
	if _, ok := ids[endWord]; !ok {
		return ret
	}

	// 确认beginWord在不在wordList里面 不在就给补进去
	if _, ok := ids[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		ids[beginWord] = n

		// 最后再更新一下wordList长度
		n++
	}

	// 生成一个cost切片 标识从beginWord只能改一个字母的变化
	cost := make([]int, n)
	// 变到这个单词的时候最少的变化次数
	edge := make([][]int, n)

	// 开始迭代wordList里面的每一个word，每次选一个word出来
	for i := range edge {
		// 然后从这个word的下一个word开始迭代，再选一个word出来
		for j := i + 1; j < n; j++ {
			// 对这两个word按照每个字节进行对比
			for b := 0; b < len(beginWord); b++ {
				// 假设这两个word在某一个字节处，不一样了
				if wordList[i][b] != wordList[j][b] {
					// 剩下部分一样 说明这两个word只有这个字节不一样
					if wordList[i][b+1:] == wordList[j][b+1:] {
						// 直接把这两个word对应位置的edge slice加上彼此
						edge[i] = append(edge[i], j)
						// 这样给定这两个word的任何一个word 就知道变换了
						edge[j] = append(edge[j], i)
					}
					// 无论剩下部分一不一样 这两个词就比对完了 直接跳出循环
					break
				}
			}
		}
		cost[i] = math.MaxInt64
	}

	// beginWord到beginWord，最短长度一定是0
	cost[ids[beginWord]] = 0

	// BFS的标准套路，就是把beginWord放到queue里面然后开始
	q := [][]int{{ids[beginWord]}}

	// 立刻开始BFS标准套路之迭代循环到queue里面没有可迭代的条件为止
	for i := 0; i < len(q); i++ {
		// BFS标准套路之从queue里面取当前node出来
		cur := q[i]
		// 看看当前node代表的path的最后一个访问到的word是啥
		curLast := cur[len(cur)-1]
		// 这个word要是endWord的话
		if wordList[curLast] == endWord {
			// 那把这个path转换成[]string 然后放到返回值里面
			p := []string{}
			for _, id := range cur {
				p = append(p, wordList[id])
			}
			ret = append(ret, p)
			// continue这个循环 不需要再往下进行下一层的搜索了
			continue
		}

		// 这个word不是endWord 那就把跟这个word能通过变换
		for _, to := range edge[curLast] {
			// 把能匹配到的所有word都给找出来 然后append进path 然后放到queue里面
			if cost[curLast]+1 <= cost[to] {
				cost[to] = cost[curLast] + 1
				path := make([]int, len(cur)+1)
				copy(path, cur)
				path[len(path)-1] = to
				// 然后BFS标准套路 把这个节点放进queue里面就不用管了
				q = append(q, path)
			}
		}
	}

	// 最后返回结果
	return ret
}
