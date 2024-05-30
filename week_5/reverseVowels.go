package week_5

/*
345. Reverse Vowels of a String

Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

Example 1:

Input: s = "hello"
Output: "holle"

Example 2:

Input: s = "leetcode"
Output: "leotcede"


Constraints:

1 <= s.length <= 3 * 105
s consist of printable ASCII characters.
*/

func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	array := []byte(s)
	i, j := 0, len(s)-1

	for i < j {
		for i < j && !vowels[array[i]] {
			i++
		}

		for i < j && !vowels[array[j]] {
			j--
		}

		array[i], array[j] = array[j], array[i]

		i++
		j--
	}

	return string(array)
}
