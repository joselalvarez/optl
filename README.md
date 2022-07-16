![Go version](https://img.shields.io/github/go-mod/go-version/joselalvarez/optl)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

The `optl package` is an Go implementation of the [null object pattern](https://en.wikipedia.org/wiki/Null_object_pattern).
This implementation is based on the Java 8 `Optional` class.

***
## [Features](#usage-overview)
* Uses the new 'Generics' feature added in Go version 1.18, all types are supported.
* The optl.Type implements the JSON Marshaller/Unmarshaller and the SQL Scanner/Valuer interfaces.
* The optl.Type is immutable except for the UnmarshalJSON and Scan methods, these methods should not be used directly.

***
## [Installation](#installation)

```
go get github.com/joselalvarez/optl@v1.0.0
```

***
## [Usage Overview](#usage-overview)

### **Import package**

```go
import (
	"github.com/joselalvarez/optl"
)
```

### **Variables and Type declaration**

```go
type Sample struct{
    Field1 optl.Type[string] // Optional string field
    Field2 optl.Type[int]    // Optional integer field
}
...
```

```go
var bar optl.Type[Sample] // Optional struct variable
```

```go
var foo optl.Type[float64] // Optional float variable

foo.IsEmpty() // true, by default the optional is empty
```

### **Constructors**

```go
foo := optl.Of("Hello optional") // Optional with value 
```

```go
bar := optl.Empty[Sample]() // Empty optional eq. -> bar := optl.Type[Sample]
```

```go
var ref *float64
...
baz := optl.OfNillable(ref) // Nillable optional
```

### **Get(), IsPresent() and IsEmpty() methods**

```go
foo := optl.Of("Hello optional")

foo.IsPresent() // true

bar := optl.Empty[int]()

bar.IsEmpty() // true

bar.Get() //Panic error!!
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

foo.OrElse(0) // 0
```

### **OrElseGet(...) method**
```go
func Zero() int{
    return 0
}
...

foo := optl.Empty[int]() 

foo.OrElseGet(Zero) // 0
```

### **Filter(...) method**
```go

func Odd(int i) bool{
    return i % 2 != 0
}
...

foo := optl.Of(1)
 
foo.Filter(Odd).IsPresent() // true
```

### **JSON Marshalling**

```go
	sample := Sample{
		Field1: optl.Of("Text ..."),
		Field2: optl.Of(1),
	}

	bytes, _ := json.Marshal(sample)

	fmt.Println(string(bytes)) // {"Field1":"Text ...","Field2":1}

	sample = Sample{
		Field1: optl.Of("Other Text ..."),
		Field2: optl.Empty[int](),
	}

	bytes, _ = json.Marshal(sample)

	fmt.Println(string(bytes)) // {"Field1":"Other Text ...","Field2":null}
```

### **JSON Unmarshalling**

```go
	var sample Sample
	json.Unmarshal([]byte(`{"Field1":"Text ...","Field2":1}`), &sample)

	fmt.Println(sample.Field1.Get()) // Text ...
	fmt.Println(sample.Field2.Get()) // 1

    var sample2 Sample
	json.Unmarshal([]byte(`{"Field1":"Other Text ...","Field2":null}`), &sample2)

	fmt.Println(sample2.Field1.Get())     // Other Text ...
	fmt.Println(sample2.Field2.IsPresent()) // false

```

### **SQL Scanner and Valuer**
The `optl.Type` implements the SQL `Scan(...)` anf `Value()` interfaces. The use of optionals as SQL values ​​is restricted to the following types:
* optl.Type[bool]
* optl.Type[string]
* optl.Type[int | int8 | int16 | int32 | int64]
* optl.Type[float32 | float64]
* optl.Type[time.Time]
* optl.Type[[]byte]

