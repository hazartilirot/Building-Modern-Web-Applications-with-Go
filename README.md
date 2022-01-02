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

**struct** looks like an interface in Typescript

Slices are like Arrays in Typescript. There is a difference in naming, and a 
pair of curly brackets is used.

```
var stringSlice []string
	stringSlice = append(stringSlice, "Trevor")
```

Short way:

`stringSclie := {"one", "two", "three}`

Maps, assigning is akin to square bracket notation

```
intMap := make(map[string]int)
	
	intMap["First"] = 1
```