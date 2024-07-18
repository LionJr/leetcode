package week_10

import "strings"

/*
1190. Reverse Substrings Between Each Pair of Parentheses
You are given a string s that consists of lower case English letters and brackets.

Reverse the strings in each pair of matching parentheses, starting from the innermost one.

Your result should not contain any brackets.



Example 1:

Input: s = "(abcd)"
Output: "dcba"
Example 2:

Input: s = "(u(love)i)"
Output: "iloveu"
Explanation: The substring "love" is reversed first, then the whole string is reversed.
Example 3:

Input: s = "(ed(et(oc))el)"
Output: "leetcode"
Explanation: First, we reverse the substring "oc", then "etco", and finally, the whole string.


Constraints:

1 <= s.length <= 2000
s only contains lower case English characters and parentheses.
It is guaranteed that all parentheses are balanced.
*/

func reverseParentheses(s string) string {
	array := []byte(s)
	stack := make([]int, 0, len(array)/2)

	for i := range array {
		if array[i] == '(' {
			stack = append(stack, i)
		}

		if array[i] == ')' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			array[i], array[j] = ' ', ' '

			for i > j {
				array[i], array[j] = array[j], array[i]
				i--
				j++
			}
		}
	}

	return strings.ReplaceAll(string(array), " ", "")
}
