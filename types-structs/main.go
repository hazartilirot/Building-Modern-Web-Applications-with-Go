package main

import (
	"log"
	"time"
)

var s = "seven"

func saySomething(s string) (string, string) {
	log.Print("s from the saySomething func is ", s)
	return s, "World"
}

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	var s2 = "six"
	log.Print("s is ", s)
	log.Print("s2 is ", s2)
	
	saySomething("xxx")
	
	user := User {
		FirstName: "Trevor",
		LastName: "Sawler",
	}
	
	log.Println(user.FirstName, user.LastName)
}