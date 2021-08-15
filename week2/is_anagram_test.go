package main

import "testing"

func TestIsAnagram(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"

	res := IsAnagram(s1, s2)
	if res != true {
		t.Errorf("test error")
	} else {
		t.Log("test success")
	}
}
