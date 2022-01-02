package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName string
}

func main() {
	stringMap := make(map[string]string)
	
	stringMap["dog"] = "Samson"
	stringMap["other-dog"] = "Cassie"
	
	intMap := make(map[string]int)
	
	intMap["First"] = 1
	intMap["Second"] = 2
	
	userMap := make(map[string]User)
	
	me := User {
		FirstName: "Trevor",
		LastName: "Sawler",
	}
	
	userMap["me"] = me
	
	log.Println(stringMap["dog"], stringMap["other-dog"])
	log.Println(intMap["First"], intMap["Second"])
	log.Println(userMap["me"].FirstName)
	
	var stringSlice []string
	stringSlice = append(stringSlice, "Trevor")
	stringSlice = append(stringSlice, "John")
	
	var intSlice []int
	intSlice = append(intSlice, 2)
	intSlice = append(intSlice, 3)
	intSlice = append(intSlice, 1)

	sort.Ints(intSlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	log.Println(stringSlice)
	log.Println(intSlice)
	log.Println(numbers[0:2])
}
