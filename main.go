package main

import "fmt"

func main() {
	x := Node{val: 21}
	x.Add(2)

	fmt.Println(x.String())
	fmt.Println(generate(5))
	fmt.Println(lengthOfLongestSubstring("abcbbc"))

	fmt.Println(canConstruct("aa", "baa"))
}
