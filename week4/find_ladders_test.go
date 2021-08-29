package main

import (
	"reflect"
	"testing"
)

func TestFindLadders(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	expected := [][]string{{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"}}

	res := FindLadders(beginWord, endWord, wordList)

	isSuccess := true
	// 判断与否与预期结果一致
	if !reflect.DeepEqual(res, expected) {
		isSuccess = false
	}

	if !isSuccess {
		t.Errorf("test error")
	} else {
		t.Log("test success")
	}
}
