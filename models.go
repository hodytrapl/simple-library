package main

import "fmt"

type Reader struct{
	ID int
	FirstName string
	LastName string
	isActive bool
}

func (r Reader) DisplayReader(){
	fmt.Printf("читатель:%s %s (ID: %s)\n",r.FirstName, r.LastName, r.ID)
}

func (r *Reader) Deactive(){
	r.isActive = false
}

type Book struct{
	ID int
	Title string
	Author string
	isIssue string
}