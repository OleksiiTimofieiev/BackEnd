package main

import (
	"Algo/arrayshashing"
	"log"
)

func main() {
	// Arrays & Hashing
	log.Println(arrayshashing.ContainsDuplicate([]int{1, 1, 2, 3, 4, 5}))
	log.Println(arrayshashing.IsAnagram("anagram","nagaram"))
	log.Println(arrayshashing.ReplaceElements([]int{17,18,5,4,6,1}))
	log.Println(arrayshashing.IsSubsequence("abc", "ahbgdc"))
	log.Println(arrayshashing.LengthOfLastWord("luffy is still joyboy"))


}
