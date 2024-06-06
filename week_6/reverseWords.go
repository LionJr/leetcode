package week_6

import "strings"

/*
151. Reverse Words in a String
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"

Example 2:

Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:

Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

Constraints:

1 <= s.length <= 104
s contains English letters (upper-case and lower-case), digits, and spaces ' '.
There is at least one word in s.

Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?
*/

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	index := strings.LastIndex(s, " ")
	array := []byte(s)

	for i, j := index, index; i >= 0; i-- {
		for i > 0 && int(array[i]) == 32 {
			i--
		}

		j = i + 1

		for i > 0 && int(array[i]) != 32 {
			i--
		}

		if i == 0 {
			array = append(array, byte(32))
		}

		array = append(array, array[i:j]...)
	}

	return string(array[index+1:])
}

/*
func reverseWords(s string) string {
	array := strings.Split(s, " ")
	for i := 0; i < len(array); i++ {
		for i < len(array) && array[i] == "" {
			array = append(array[:i], array[i+1:]...)
		}
	}

	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		for i < j && array[i] == "" {
			i++
		}

		for i < j && array[j] == "" {
			j--
		}

		array[i], array[j] = array[j], array[i]
	}

	return strings.Join(array, " ")
}
*/
