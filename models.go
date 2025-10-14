package main

import "fmt"

type Reader struct{
	ID int
	FirstName string
	LastName string
	isActive bool
}

func (r Reader) DisplayReader(){
	fmt.Printf("читатель:%s %s (ID: %d)\n",r.FirstName, r.LastName, r.ID)
}

func (r *Reader) Deactive(){
	r.isActive = false
}

func (r Reader) String() string{
	status :=""
	if(r.isActive){
		status="active"
	}else{
		status="noactive"
	}
	return fmt.Sprintf("пользователь %s %s, id: %d, %s",r.FirstName, r.LastName, r.ID,status)
}

type Book struct{
	ID int
	Title string
	Author string
	isIssue bool
}