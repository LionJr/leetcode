package week_9

/*
1550. Three Consecutive Odds
Given an integer array arr, return true if there are three consecutive odd numbers in the array. Otherwise, return false.


Example 1:

Input: arr = [2,6,4,1]
Output: false
Explanation: There are no three consecutive odds.

Example 2:

Input: arr = [1,2,34,3,4,5,7,23,12]
Output: true
Explanation: [5,7,23] are three consecutive odds.


Constraints:

1 <= arr.length <= 1000
1 <= arr[i] <= 1000
*/

func threeConsecutiveOdds(arr []int) bool {
	oddCount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			oddCount++
			if oddCount == 3 {
				return true
			}
		} else {
			oddCount = 0
		}
	}

	return false
}
