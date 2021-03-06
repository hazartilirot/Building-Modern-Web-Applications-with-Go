# About the course (snippet)

This course is well-suited for both absolute beginners, and for developers who already know something about web development, but want to add Go to their toolbox.

We start with an overview of the Go language, and then cover everything you need to get started writing web applications, including an overview of HTML5, a survey of JavaScript and JavaScript modules, how to work with Cascading Style Sheets to make our application look the way we want, and much more.

The major project in this course is building a bookings and reservation system for a Bed & Breakfast. Visitors to our site will be able to search for accommodations by date and make an online reservation, and the site owner will be able to manage reservations from a secure back end.

# PERSONAL NOTES 

### Pointers

&-ampersand for references
```
func name(<&argument)
```

*-star for pointers
```
func name (<parameter *type>)
```

a short variable declaration
```
:= 
```

A variable scope seems to be somewhat similar to JS. At least what I'm shown 
in **types-structs** lesson

A letter case matters for exporting values and function. If a name begins 
with a capital letter it might be available outside the current package.

### struct 

looks like an interface in Typescript

### Slices 

are like Arrays in Typescript. There is a difference in naming, and a 
pair of curly brackets is used.

```
var stringSlice []string
	stringSlice = append(stringSlice, "Trevor")
```

Short way:

`stringSclie := {"one", "two", "three}`

### Maps, 

assignment is akin to square bracket notation

```
intMap := make(map[string]int)
	
	intMap["First"] = 1
```

### Conditional 

uses **== - double equals** to compare values and types. Additionally, 
you can't write an expression omitting curly braces. There is another one 
feature is that we don't wrap an expression with parentheses like in JS or TS.

```
    myNum := 100
	isTrue := false
	
	if myNum > 99 && !isTrue {
		log.Println("MyNum is greater than 99 and is True is set to true")
	}
    
```
### Switch

break a case isn't required. There is no indentation between switch and cases

```
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
```

### Loops and Ranging over data

**for** loop isn't wrapped in parentheses. In the second case a placeholder is 
used to omit an index value. Otherwise, a compiler would complain.

```
    for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	mySlice := []string{"dog", "cat", "horse", "fish", "banana"}
	
	for _, value := range mySlice {
		log.Println(value)
	}
```

### Interfaces

There isn't a key word or explicit implementation of an interface. Once a 
func parameter accepts the object type (or many func) and then, those funcs 
become a kind of methods of the object and if its funcs "implements" an 
interface they are suitable to be used as a polymorphism. A slightly 
different approach than once is used to think of an interface implementation. 

### Modules and Packages

Conventionally, after a key word "init" a go developer is used his GitHub repo.
````
go mod init github.com/username/project-name
````

### Channels

Is a weird method of passing values that must be lately reconsidered.

### 3-ty part libraries

Installing a 3-ty part libraries is followed by the command
````
go get github.com/bmizerany/pat
````
Don't use `http://` before a domain name in the command. Otherwise it fails while executing