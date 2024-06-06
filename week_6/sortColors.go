package week_6

/*
75. Sort Colors
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

Example 1:

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Example 2:

Input: nums = [2,0,1]
Output: [0,1,2]

Constraints:

n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.

Follow up: Could you come up with a one-pass algorithm using only constant extra space?
*/

func sortColors(nums []int) {
	i, j := 0, 0
	length := len(nums)

	for n := 2; n >= 1; n-- {
		for i, j = 0, 0; i < length; i++ {
			if nums[i] != n {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			}
		}
		length = j
	}
}

/*
func sortColors(nums []int)  {
    colors := make([]int, 3)
    for i := range nums {
        colors[nums[i]]++
    }

    index := 0
    for i := 0; i < len(colors); i++ {
        for j := 0; j < colors[i]; j++ {
            nums[index] = i
            index++
        }
    }
}
*/
