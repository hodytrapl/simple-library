package main

import "fmt"

type Notifer interface {
	Notify(message string)
}

type EmailNotifier struct {
	EmailAddress string
}
type SMSNotifier struct {
	PhoneNumber string
}

func (e EmailNotifier) Notify(message string) {
	fmt.Printf("Отправляю email на %s: \"%s\".\n", e.EmailAddress, message)
}
func (s SMSNotifier) Notify(message string) {
	fmt.Printf("Отправляю SMS на %s: \"%s\".\n", s.PhoneNumber, message)
}
