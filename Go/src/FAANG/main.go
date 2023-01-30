package main

import (
	"fmt"

	"faang/merge_intervals"
	"faang/sliding_window"
	"faang/two_pointer"
)

func printInColor() func(string) {
	Yellow := "\033[33m"
	// Blue := "\033[34m"
	Reset := "\033[0m"

	return func(message string) {
		fmt.Println("--- " + Yellow + message + Reset + " ---")
	}
}

func main() {
	logger := printInColor()

	// --- sliding window ---
	logger("Sliding window")
	fmt.Println("MaxSumSubarray: ", sliding_window.MaxSumSubarray([]int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 3))
	fmt.Println("SmallestSubarray: ", sliding_window.SmallestSubarray(8, []int{4, 2, 2, 7, 8, 1, 2, 8, 10}))
	fmt.Println("LengthOfLongestSubstring: ", sliding_window.LengthOfLongestSubstring("wertwteeeeeerrrtesta"))
	fmt.Println("ContainsNearbyDuplicate: ", sliding_window.ContainsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	fmt.Println("lengthOfLongestSubstring: ", sliding_window.LengthOfLongestSubstring("abcdeafbdgcbb"))

	// --- two pointers ---
	logger("Two Pointers")
	fmt.Println("TwoSumSorted: ", two_pointer.TwoSumSorted([]int{-3, 2, 3, 3, 6, 8, 15}, 6))
	fmt.Println("GetMaxSumSubArrayOfSizeKM2: ", two_pointer.GetMaxSumSubArrayOfSizeKM2([]int{1, 10, -1, -2, 7, 3, -1, 7}, 4))
	fmt.Println("MaxArea: ", two_pointer.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println("Intersect: ", two_pointer.Intersect([]int{1, 1, 2, 2}, []int{1, 2, 1, 2}))

	// --- merge intervals ---
	logger("Merge Intervals")
	fmt.Println("RemoveCoveredIntervals: ", merge_intervals.RemoveCoveredIntervals([][]int{{1, 4}, {3, 6}, {2, 8}}))
	fmt.Println("Merge: ",merge_intervals.Merge([][]int{{1, 4}, {3, 6}, {2, 8}}))
}
