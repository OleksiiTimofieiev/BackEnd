package merge_intervals

import (
	"sort"
)

// Given an array intervals where intervals[i] = [li, ri] represent the interval [li, ri), remove all intervals that are covered by another interval in the list.
// The interval [a, b) is covered by the interval [c, d) if and only if c <= a and b <= d.
// Return the number of remaining intervals.

// Input: intervals = [[1,4],[3,6],[2,8]]
// Output: 2
// Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.

// https://leetcode.com/problems/

/
// sort.Slice(intervals, func(i, j int) bool {return intervals[i][0] < intervals[j][0]})

func RemoveCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	// fmt.Println(intervals)

	end, res := 0, 0

	for _, interval := range intervals {
		if interval[1] > end {
			end = interval[1]
			res++
		}
	}

	return res
}

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for k := 1; k < len(intervals); k++ {
		lastEnd := res[len(res)-1][1]
		if intervals[k][0] > lastEnd {
			res = append(res, intervals[k])
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[k][1])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
