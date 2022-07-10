![Go version](https://img.shields.io/github/go-mod/go-version/joselalvarez/optl)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

`optl package` is an Go implementation of the [null object pattern](https://en.wikipedia.org/wiki/Null_object_pattern).


***
## [Features](#usage-overview)
* Uses the 'Generic Types' feature added to Go in version 1.18. All basic types are supported, plus structs, arrays and slices.
* The implementation style is based on the Java 8 optional type. It supports all methods except "map()", "flatMap()" (Go does not support generic types in methods) and "orElseThrow" (In this case I think it does not suit the style of Go). As in the case of the java optional, it is an immutable (except for the JSON and SQL serialization methods).
* Implementation of the JSON Marshaller and JSON Unmarshaller interfaces.
* Implementation of the SQL Scaner and SQL Valuer interfaces.

***
## [Installation](#installation)

```
go get github.com/joselalvarez/optl@v0.1.0
```

***
## [Usage Overview](#usage-overview)

### **Variable definition**

```go
import (
	"github.com/joselalvarez/optl"
)
```

```go
optl.Type[<type>] // Type declaration
```
Examples:
```go
var foo optl.Type[string] // Optional string variable
```

```go
type Sample struct{
    Field1 optl.Type[string] // Optional string field
    Field2 optl.Type[int] // Optional integer field
}
```

```go
var bar optl.Type[Sample] // Optional struct variable
```

### **Constructors**

```go
foo := optl.Of("Hello optional") // Optional with value 
```

```go
bar := optl.Empty[Sample]() // Empty optional
```

```go
var pointer *float64
...
baz := optl.OfNillable(pointer) // Nillable optional
```

### **Get(), IsPresent() and IsEmpty() methods**
```go
foo := optl.Of("Hello optional")

if foo.IsPresent() {
    fmt.Println(foo.Get())
}

bar := optl.Empty[int]()

if bar.IsEmpty(){
    fmt.Println("Optional 'bar' is empty")
}
```

### **IfPresent(...) method**
```go
foo := optl.Of("Hello optional")

foo.IfPresent(func(value string){
    fmt.Println(value)
})
```

### **OrElse(...) method**
```go
foo := optl.Empty[int]()

fmt.Println(foo.OrElse(0))
```

### **OrElseGet(...) method**
```go
func Zero() int{
    return 0
}

foo := optl.Empty[int]() 

fmt.Println(foo.OrElseGet(Zero))
```

### **Filter(...) method**
```go

func Odd(i int) bool{
    return i % 2 != 0
}

foo := optl.Of(1)
 
fmt.Println("Is odd value?", foo.Filter(Odd).IsPresent())
```
