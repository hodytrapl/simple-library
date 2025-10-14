package main

import "fmt"

type Reader struct{
	ID int
	FirstName string
	LastName string
	isActive bool
}

type Book struct{
	ID int
	Title string
	Author string
	isIssue string
}