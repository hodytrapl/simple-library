package main

import "fmt"

func main() {
	user1 :=Reader{
		ID:1,
		FirstName:"g",
		LastName:"gg",
		isActive:true,
	}

	user1.Deactive()
	fmt.Println(user1)
}