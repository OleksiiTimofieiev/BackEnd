package sliding_window

import (
	"math"
)

func max[T int | float64](a T, b T) T {
	if a > b {cd ..
		return a
	}
	return b
}

func min[T int | float64](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func MaxSumSubarray(arr []int, k int) int {
	maxValue := math.MinInt
	currentRunningSum := 0

	for i := 0; i < len(arr); i++ {
		currentRunningSum += arr[i]

		if i >= k-1 {
			maxValue = max(maxValue, currentRunningSum)
			currentRunningSum -= arr[i-(k-1)]
		}
	}
	return maxValue

}

func SmallestSubarray(targetSum int, arr []int) int {
	minWindowSize := math.MaxInt
	currentWindowSum := 0;
	windowStart := 0;

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		currentWindowSum += arr[windowEnd];

		for(currentWindowSum >= targetSum) {
			minWindowSize = min(minWindowSize, windowEnd - windowStart + 1);
			currentWindowSum -= arr[windowStart];
			windowStart++;
		}
	}
	return minWindowSize;
}
 