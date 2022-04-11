package main

import "log"

func main() {
	var whatToSay string
	var i int
	whatToSay, _ = saySomething("Hi")

	log.Println(whatToSay)

	saySomethingElse, _ := saySomething("Goodbay")

	log.Println(saySomethingElse)

	log.Println(saySomething("Hello "))

	log.Println(i) // would be "0" since no value has been assigned
}

func saySomething(word string) (string, string) {
	return word, "World"
}
