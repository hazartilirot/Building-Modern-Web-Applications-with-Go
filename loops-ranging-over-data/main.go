package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	mySlice := []string{"dog", "cat", "horse", "fish", "banana"}
	
	for _, value := range mySlice {
		log.Println(value)
	}
	
	myMap := make(map[string]string)
	myMap["dog"] = "dog"
	myMap["fish"] = "fish"
	myMap["hat"] = "hat"
	
	for _, v := range myMap {
		log.Println(v)
	}
	
	type User struct {
		FirstName string
		LastName string
	}
	
	var userSlice []User
	
	u1 := User{
		FirstName: "Trevor",
	}
	u2 := User{
		FirstName: "Sam",
	}

	userSlice = append(userSlice, u1, u2)
	
	for _, v := range userSlice {
		log.Println(v.FirstName)
	}
}
