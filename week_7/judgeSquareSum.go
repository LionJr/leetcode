package week_7

import "math"

/*
633. Sum of Square Numbers

Given a non-negative integer c, decide whether there're two integers a and b such that a2 + b2 = c.

Example 1:

Input: c = 5
Output: true
Explanation: 1 * 1 + 2 * 2 = 5

Example 2:

Input: c = 3
Output: false

Constraints:

0 <= c <= 231 - 1
*/

func judgeSquareSum(c int) bool {
	i, j, n := 0, int(math.Sqrt(float64(c))), 0
	for i <= j {
		n = i*i + j*j
		if n > c {
			j--
		} else if n < c {
			i++
		} else {
			return true
		}
	}

	return false
}
