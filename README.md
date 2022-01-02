# NOTES 

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

uses **== - double equals** to compare values and types. Additionally 
you can't write an expression omitting curly braces. There is another one 
feature is that we don't wrap an expression with parenthesis like in JS or TS.

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


