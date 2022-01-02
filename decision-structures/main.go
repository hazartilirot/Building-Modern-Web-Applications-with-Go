package main

import "log"

func main() {
	var isTrue bool
	
	isTrue = true
	
	if isTrue {
		log.Println("isTrue is ", isTrue)
	} else {
		log.Println("isTrue is ", isTrue)
	}
	
	cat := "cat"
	
	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Print("Cat is not cat")
	}
	
	myNum := 100
	isTrue = false
	
	if myNum > 99 && !isTrue {
		log.Println("MyNum is greater than 99 and is True is set to true")
	}
	
	myVar := "dog"
	
	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")

	default:
		log.Println("cat is something else")
	}
}
