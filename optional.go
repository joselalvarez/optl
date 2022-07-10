package optl

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

//The Optional Type.
//
//Example: var foo optl.Type[string]
type Type[T interface{}] struct {
	value   T
	present bool
}

//Constructor of an optional with value.
//
//Example 1: foo := optl.Of("Hello optional")
//
//Example 2: bar := optl.Of(2.3)
//
//Example 3: baz := optl.Of[int64](3)
//
//Example 4: quux := optl.Of(MyStruct{...})
func Of[T interface{}](v T) Type[T] {
	return Type[T]{value: v, present: true}
}

//Constructor of an optional with nillable value (pointer to value).
//
//var bar * string
//
//...
//
//Example: foo := optl.OfNillable(bar)
func OfNillable[T interface{}](v *T) Type[T] {
	if v != nil {
		return Type[T]{value: *v, present: true}
	}
	return Empty[T]()
}

//Constructor of an empty optional.
//
//Sample 1: foo := optl.Empty[string]()
//
//Sample 2: bar := optl.Empty[MyStruct]()
func Empty[T interface{}]() Type[T] {
	return Type[T]{value: *new(T), present: false}
}

//Return true if value is present.
func (o Type[T]) IsPresent() bool {
	return o.present
}

//Return true if value is not present.
func (o Type[T]) IsEmpty() bool {
	return !o.present
}

//Return the value of the optional, if it is not present throw a panic error.
func (o Type[T]) Get() T {
	if o.IsPresent() {
		return o.value
	}
	panic("Value is not present (nil)")
}

//If value is present return the current value, else return the value passed as parameter.
func (o Type[T]) OrElse(v T) T {
	if o.IsPresent() {
		return o.value
	}
	return v
}

//If value is present return the current value, else return the value returned by lambda function passed as parameter.
func (o Type[T]) OrElseGet(f func() T) T {
	if o.IsPresent() {
		return o.value
	}
	return f()
}

//If the value is present executes the lambda function passed as parameter with the current value.
func (o Type[T]) IfPresent(f func(T)) {
	if o.IsPresent() {
		f(o.value)
	}
}

//If the value is present executes the lambda function passed as parameter with the current value.
//If this function return false return a empty optional else return the current optional.
func (o Type[T]) Filter(f func(T) bool) Type[T] {
	if o.IsPresent() && f(o.value) {
		return o
	}
	return Empty[T]()
}

//JSON marshaller
func (o Type[T]) MarshalJSON() ([]byte, error) {
	if o.IsPresent() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

//JSON unmarshaller
func (o *Type[T]) UnmarshalJSON(data []byte) error {

	o.value = *new(T)
	o.present = false

	if string(data) == "null" {
		return nil
	}

	var value T

	err := json.Unmarshal(data, &value)

	if err != nil {
		return err
	}

	o.value = value
	o.present = true

	return nil
}

//SQL valuer
func (o Type[T]) Value() (driver.Value, error) {

	if o.IsPresent() {

		t := reflect.TypeOf(o.value)
		k := t.Kind()
		ser, e := json.Marshal(o.value)

		if e != nil {
			return nil, e
		}

		if k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 {
			var v int64
			e = json.Unmarshal(ser, &v)
			if e != nil {
				return nil, e
			}
			return v, nil
		} else if k == reflect.Float32 || k == reflect.Float64 {
			var v float64
			e = json.Unmarshal(ser, &v)
			if e != nil {
				return nil, e
			}
			return v, nil
		} else if k == reflect.Bool {
			var v bool
			e = json.Unmarshal(ser, &v)
			if e != nil {
				return nil, e
			}
			return v, nil
		} else if k == reflect.String {
			var v string
			e = json.Unmarshal(ser, &v)
			if e != nil {
				return nil, e
			}
			return v, nil
		} else if (k == reflect.Slice || k == reflect.Array) && t.Elem().Kind() == reflect.Uint8 {
			var sl []byte
			e = json.Unmarshal(ser, &sl)
			if e != nil {
				return nil, e
			}
			return sl, nil
		} else if t == reflect.TypeOf(time.Now()) {
			var ti time.Time
			e = json.Unmarshal(ser, &ti)
			if e != nil {
				return nil, e
			}
			return ti, nil
		}

		return nil, fmt.Errorf("sql value type '%s' of kind '%s' not supported", t.String(), k.String())
	}
	return nil, nil
}

//SQL scanner
func (o *Type[T]) Scan(value interface{}) error {
	if value != nil {
		ser, e := json.Marshal(value)
		if e != nil {
			return e
		}

		var v T
		e = json.Unmarshal(ser, &v)
		if e != nil {
			return e
		}
		o.value = v
		o.present = true
		return nil
	}
	o.value = *new(T)
	o.present = false
	return nil
}
