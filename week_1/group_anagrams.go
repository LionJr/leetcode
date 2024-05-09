package week_1

import "sort"

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:

Input: strs = [""]
Output: [[""]]

Example 3:

Input: strs = ["a"]
Output: [["a"]]

Constraints:

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/

func groupAnagrams(strs []string) [][]string {

	strsMap := make(map[string][]string)
	for _, str := range strs {
		key := sortBytes(str)
		strsMap[key] = append(strsMap[key], str)
	}

	results := make([][]string, len(strsMap))
	index := 0
	for _, array := range strsMap {
		results[index] = array
		index++
	}

	return results
}

func sortBytes(s string) string {
	array := []byte(s)
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	return string(array)
}
