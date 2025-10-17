package main

import "fmt"

func main() {
	iduser1 := 1
	user1 := &Reader{
		ID:        &iduser1,
		FirstName: "g",
		LastName:  "gg",
		isActive:  true,
	}
	fmt.Println(user1)

	book1 := &Book{
		ID:            1,
		Title:         "война и мир",
		Year:          1985,
		Author:        "tolstoy",
		isIssue:       false,
		ReaderTakerID: 0,
	}

	fmt.Println(book1)
	book1.IssueBook(user1)
	fmt.Println(book1)

	user1.Deactive()
	fmt.Println(user1)
	fmt.Println("---")
	book1.IssueBook(user1)

	Slice_notifer:=[]Notifer{}
	email:=EmailNotifier{EmailAddress:"kostik290820077@gmail.com"}
	phone:=SMSNotifier{PhoneNumber:"+79187025870"}

	Slice_notifer = append(Slice_notifer,email, phone)
	for i:=0;i<len(Slice_notifer);i++{
		Slice_notifer[i].Notify("Ваша книга просрочена!")
	}
}
