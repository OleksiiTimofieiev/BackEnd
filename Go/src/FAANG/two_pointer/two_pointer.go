package two_pointer

import (
	"math"
	"sort"

	"faang/utils"
)

// https://codeshare.io/X8V9n8
// You are given an integer array height of length n.
// There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// .
// .   .
// . . .
// . . .
// . . .
// . . .
// . . .

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

// func MaxArea(height []int) int {
// ...
// }

func TwoSumSorted(nums []int, target int) [2]int {
	var start int
	end := len(nums) - 1
	var res [2]int

	for start < end {
		sum := nums[start] + nums[end]
		if sum == target {
			res[0] = start + 1
			res[1] = end + 1
			break
		} else if sum < target {
			start++
		} else {
			end--
		}
	}
	return res
}

func GetMaxSumSubArrayOfSizeKM2(A []int, k int) int {
	var windowSum int
	maxSum := math.MinInt

	for i := 0; i < k; i++ {
		windowSum += A[i]
	}
	maxSum = utils.Max(maxSum, windowSum)
	// added as a part of correction
	for windowEndIndex := k; windowEndIndex < len(A); windowEndIndex++ {
		windowSum += A[windowEndIndex] - A[windowEndIndex-k]
		maxSum = utils.Max(maxSum, windowSum)
	}
	return maxSum
}

func MaxArea(height []int) int {
	maxarea := 0
	left := 0
	right := len(height) - 1
	for left < right {
		width := right - left
		maxarea = utils.Max(maxarea, utils.Min(height[left], height[right])*width)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxarea
}

// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each element in the result must appear as many times as it shows in both arrays
// and you may return the result in any order.
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]

// func intersect(nums1 []int, nums2 []int) []int {
// }

// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

func Intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)

	sort.Ints(nums1)
	sort.Ints(nums2)

	for p1, p2 := 0, 0; p1 < len(nums1) && p2 < len(nums2); {
		if nums1[p1] == nums2[p2] {
			result = append(result, nums1[p1])
			p1, p2 = p1+1, p2+1
		} else if nums1[p1] < nums2[p2] {
			p1++
		} else {
			p2++
		}
	}

	return result
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func RemoveDuplicates(nums []int) int {
	j := 1

	prev := nums[0]

	for i := 1; i < len(nums); i++ {
		if prev != nums[i] {
			j++
			nums[j-1] = nums[i]
		}
		prev = nums[i]
	}
	for i := j; i < len(nums); i++ {
		nums[i] = '_'
	}
	return j
}
