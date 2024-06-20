package week_7

/*
560. Subarray Sum Equals K

Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.

Example 1:

Input: nums = [1,1,1], k = 2
Output: 2

Example 2:

Input: nums = [1,2,3], k = 3
Output: 2

Constraints:

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/

func subarraySum(nums []int, k int) int {
	previousSubsSum := make(map[int]int)
	previousSubsSum[0] = 1

	result, sum := 0, 0
	for i := range nums {
		sum += nums[i]
		if prevSum, ok := previousSubsSum[sum-k]; ok {
			result += prevSum
		}

		previousSubsSum[sum]++
	}

	return result
}
