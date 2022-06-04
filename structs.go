package main

import "fmt"

func main() {
	type user struct {
		ID int
		fn string
		ln string
	}

	var u user
	u.ID = 1
	u.fn = "aditya"
	u.ln = "pathania"

	fmt.Println(u)

	u2 := user{ID: 2, fn: "korewa", ln: "chinchin"}
	fmt.Println(u2)
}
