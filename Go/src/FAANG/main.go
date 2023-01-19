package main

import (
	"fmt"

	"faang/sliding_window"
)

func main() {
	fmt.Println("MaxSumSubarray: ", sliding_window.MaxSumSubarray([]int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 3))
	fmt.Println("SmallestSubarray: ", sliding_window.SmallestSubarray(8, []int{4, 2, 2, 7, 8, 1, 2, 8, 10}))

}
