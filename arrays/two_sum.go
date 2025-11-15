package arrays

/*
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: nums[0] + nums[1] == 9

Constraints:
- 2 <= nums.length <= 1e4
- -1e9 <= nums[i] <= 1e9
- -1e9 <= target <= 1e9

Solution:
Time Complexity: O(n)
Space Complexity: O(n)
*/

// TwoSum returns the indices of the two numbers in nums that add up to target.
// You may assume exactly one solution exists, and you may not use the same element twice.
// Time Complexity: O(n)
// Space Complexity: O(n)
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, n := range nums {
		rest := target - n

		if j, ok := m[rest]; ok {
			return []int{j, i}
		}

		m[n] = i
	}

	return nil
}
