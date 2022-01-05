package main

import (
	"github.com/username/project-name/helpers"
	"log"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)
	
	go CalculateValue(intChan)
	
	num := <- intChan
	
	log.Println("Random number:", num)
}
