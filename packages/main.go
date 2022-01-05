package main

import (
 "github.com/username/project-name/helpers"
 "log"
)

func main() {
 log.Println("Hello")
 
 var myVar helpers.SomeType
 
 myVar.TypeName = "Some name"
 
 log.Println(myVar.TypeName)
} 