// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。
//
// 输入: s = "anagram", t = "nagaram"
// 输出: true
//
// 输入: s = "rat", t = "car"
// 输出: false
//
// 解题思路:从另一个角度考虑，t 是 s 的异位词等价于两个字符串中字符出现的种类和次数均相等。
// 由于字符串只包含 26 个小写字母，因此我们可以维护一个长度为 26 的频次数组table，先遍历记录字符串 s 中字符出现的频次，然后遍历字符串 t，
// 减去 table 中对应的频次，如果出现 table[i]<0，则说明 t 包含一个不在 s 中的额外字符，返回 false 即可
//

package main

func IsAnagram(s string, t string) bool {
	// 定义26个元素数组 来表示字频的hashTable
	var c1, c2 [26]int
	// 分别遍历s和t 将出现的字符存入hashTable中
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}

	// 如果两个hashTable相等 则代表出现的频次相等 则为字母异位词
	return c1 == c2
}