package main

import "fmt"

func main() {
	x := Node{val: 21}
	x.Add(2)

	fmt.Println(x.String())
}
