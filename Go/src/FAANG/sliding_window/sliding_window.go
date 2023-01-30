package sliding_window

import (
	"math"

	"faang/utils"
)

func MaxSumSubarray(arr []int, k int) int {
	maxValue := math.MinInt
	currentRunningSum := 0

	for i := 0; i < len(arr); i++ {
		currentRunningSum += arr[i]

		if i >= k-1 {
			maxValue = utils.Max(maxValue, currentRunningSum)
			currentRunningSum -= arr[i-(k-1)]
		}
	}
	return maxValue

}

func SmallestSubarray(targetSum int, arr []int) int {
	minWindowSize := math.MaxInt
	currentWindowSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		currentWindowSum += arr[windowEnd]

		for currentWindowSum >= targetSum {
			minWindowSize = utils.Min(minWindowSize, windowEnd-windowStart+1)
			currentWindowSum -= arr[windowStart]
			windowStart++
		}
	}
	return minWindowSize
}

// func LengthOfLongestSubstring(str string) int {
// 	i := 0
// 	j := 0
// 	max := 0
// 	set := make(map[uint8]bool)

// 	for i < len(str) {
// 		ch := str[i]
// 		for {
// 			if _, ok := set[ch]; ok {
// 				delete(set, ch)
// 				j = j + 1
// 			} else {
// 				break
// 			}
// 		}
// 		set[ch] = true
// 		max = utils.Max(max, i-j+1)
// 		i = i + 1
// 	}
// 	return max
// }

// https://leetcode.com/problems/contains-duplicate-ii/
func ContainsNearbyDuplicate(nums []int, k int) bool {
	mapUnique := make(map[interface{}]int)

	for i := 0; i < len(nums); i++ {
		digit := nums[i]

		if v, exists := mapUnique[digit]; exists {
			if i-v <= k {
				return true
			}
		}
		mapUnique[digit] = i

	}
	return false
}

func LengthOfLongestSubstring(s string) int {
	mp := make(map[uint8]int)
	i := 0
	ans := 0

	for j := 0; j < len(s); j++ {
		if v, exists := mp[s[j]]; exists {
			i = utils.Max(v, i)
		}
		ans = utils.Max(ans, j-i+1)
		mp[s[j]] = j + 1
	}
	return ans
}

// https://codeshare.io/OadVqXN
// class Solution {
// 	public:
// 		int longestBeautifulSubstring(string word) {
// 			const auto n = word.size();

// 			int cnt = 1;
// 			int len = 1;
// 			int max_len = 0;
// 			for (int i = 1; i != n; ++i) {
// 				if (word[i - 1] == word[i]) {
// 					++len;
// 				} else if (word[i - 1] < word[i]) {
// 					++len;
// 					++cnt;
// 				} else {
// 					cnt = 1;
// 					len = 1;
// 				}

// 				if (cnt == 5) {
// 					max_len = max(max_len, len);
// 				}
// 			}
// 			return max_len;
// 		}
// 	};
