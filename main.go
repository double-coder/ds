package main

import (
	"fmt"
)

func main() {
	// arr1 := []int{100, 16, 9, 1, 0}
	x := Node{val: 21}
	x.Add(2)
	x.Add(3)
	x.Add(4)
	x.Add(5)
	fmt.Println(x.String())
	fmt.Println(generate(5))
	fmt.Println(lengthOfLongestSubstring("abcbbc"))

	fmt.Println(canConstruct("aa", "baa"))
	fmt.Println(isIsomorphic("egg", "add"))

	// reverse a linked list
	fmt.Println(reverse(&x))
}
