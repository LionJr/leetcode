package leetcode

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true

Example 2:

Input: s = "rat", t = "car"
Output: false

Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap1 := make(map[byte]int)
	charMap2 := make(map[byte]int)
	for i := range s {
		charMap1[s[i]]++
		charMap2[t[i]]++
	}

	for key, value := range charMap1 {
		if value != charMap2[key] {
			return false
		}
	}
	return true
}
