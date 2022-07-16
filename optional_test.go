package optl_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/joselalvarez/optl"
)

type dto struct {
	Field1 optl.Type[string]
	Field2 optl.Type[int32]
}

func assertPanic[T interface{}](t *testing.T, f func() T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()
	f()
}

func getValues() []interface{} {
	values := make([]interface{}, 12)
	values[0] = "Hello optional"
	values[1] = 1
	values[2] = int8(1)
	values[3] = int16(1)
	values[4] = int32(1)
	values[5] = int64(1)
	values[6] = float32(1.5)
	values[7] = float64(1.5)
	values[8] = true
	values[9] = time.Now()
	values[10] = []byte{1, 2, 3}
	values[11] = dto{
		Field1: optl.Of("Text ..."),
		Field2: optl.Of(int32(1)),
	}
	return values
}

func getElseValues() []interface{} {
	values := make([]interface{}, 12)
	values[0] = "Bye optional"
	values[1] = 2
	values[2] = int8(2)
	values[3] = int16(2)
	values[4] = int32(2)
	values[5] = int64(2)
	values[6] = float32(2.5)
	values[7] = float64(2.5)
	values[8] = false
	values[9] = time.Now()
	values[10] = []byte{4, 5, 6}
	values[11] = dto{
		Field1: optl.Of("Text 2 ..."),
		Field2: optl.Of(int32(2)),
	}
	return values
}

func TestF_Of(t *testing.T) {

	for _, value := range getValues() {
		t.Run("new optional of type "+reflect.TypeOf(value).String(), func(t *testing.T) {
			op := optl.Of(value)
			if !op.IsPresent() || op.IsEmpty() || !reflect.DeepEqual(op.Get(), value) {
				t.Fail()
			}
		})
	}
}

func TestF_OfNillable(t *testing.T) {

	for _, value := range getValues() {
		t.Run("new optional nillable of type "+reflect.TypeOf(value).String(), func(t *testing.T) {
			op := optl.OfNillable(&value)
			if !op.IsPresent() || op.IsEmpty() || !reflect.DeepEqual(op.Get(), value) {
				t.Fail()
			}
		})
	}

	t.Run("new optional nillable nil", func(t *testing.T) {
		op := optl.OfNillable[interface{}](nil)
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
	})
}

func TestF_Empty(t *testing.T) {

	t.Run("new optional empty", func(t *testing.T) {
		op := optl.Empty[interface{}]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})
}

func TestM_Type_OrElse(t *testing.T) {

	elseValues := getElseValues()
	for i, value := range getValues() {
		t.Run("OrElse with value present of type "+reflect.TypeOf(value).String(), func(t *testing.T) {
			if !reflect.DeepEqual(optl.Of(value).OrElse(elseValues[i]), value) {
				t.Fail()
			}
		})
	}

	for i, value := range getValues() {
		t.Run("OrElse with value not present of type "+reflect.TypeOf(value).String(), func(t *testing.T) {
			if !reflect.DeepEqual(optl.Empty[interface{}]().OrElse(elseValues[i]), elseValues[i]) {
				t.Fail()
			}
		})
	}

}

func TestM_Type_OrElseGet(t *testing.T) {
	elseValues := getElseValues()
	for i, value := range getValues() {
		t.Run("OrElseGet with value present of type "+reflect.TypeOf(value).String(), func(t *testing.T) {
			if !reflect.DeepEqual(optl.Of(value).OrElseGet(func() interface{} { return elseValues[i] }), value) {
				t.Fail()
			}
		})
	}

	for i, value := range getValues() {
		t.Run("OrElseGet with value not present of type "+reflect.TypeOf(value).String(), func(t *testing.T) {
			if !reflect.DeepEqual(optl.Empty[interface{}]().OrElseGet(func() interface{} { return elseValues[i] }), elseValues[i]) {
				t.Fail()
			}
		})
	}
}

func TestM_Type_IfPresent(t *testing.T) {
	for _, value := range getValues() {
		t.Run("IfPresent with value present of type"+reflect.TypeOf(value).String(), func(t *testing.T) {
			execute := false
			optl.Of(value).IfPresent(func(v interface{}) {
				if !reflect.DeepEqual(value, v) {
					t.Fail()
				}
				execute = true
			})
			if !execute {
				t.Fail()
			}
		})
	}

	t.Run("IfPresent with value not present", func(t *testing.T) {
		optl.Empty[interface{}]().IfPresent(func(v interface{}) {
			t.Fail()
		})
	})
}

func TestM_Type_Filter(t *testing.T) {

	for _, value := range getValues() {
		t.Run("Filter not match of type "+reflect.TypeOf(value).String(), func(t *testing.T) {
			op := optl.Of(value).Filter(func(v interface{}) bool {
				if !reflect.DeepEqual(value, v) {
					t.Fail()
				}
				return false
			})
			if op.IsPresent() {
				t.Fail()
			}
		})
	}

	for _, value := range getValues() {
		t.Run("Filter match of type "+reflect.TypeOf(value).String(), func(t *testing.T) {
			op := optl.Of(value).Filter(func(v interface{}) bool {
				if !reflect.DeepEqual(value, v) {
					t.Fail()
				}
				return true
			})
			if op.IsEmpty() {
				t.Fail()
			}
		})
	}

	t.Run("Filter empty", func(t *testing.T) {
		optl.Empty[string]().Filter(func(v string) bool {
			t.Fail()
			return false
		})
	})

}

func TestM_Type_MarshallingJSON(t *testing.T) {

	for _, value := range getValues() {
		t.Run("Marshalling optional of type "+reflect.TypeOf(value).String(), func(t *testing.T) {
			op := optl.Of(value)
			var op2 optl.Type[interface{}]

			v1, _ := json.Marshal(op)
			json.Unmarshal(v1, &op2)

			v2, _ := json.Marshal(op2)

			if !reflect.DeepEqual(v1, v2) {
				t.Fail()
			}
		})
	}

	t.Run("Marshalling empty optional", func(t *testing.T) {
		op := optl.Empty[interface{}]()
		op2 := optl.Of("Hello optional")

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ string) {
			t.Fail()
		})
	})
}

func TestM_Type_Value(t *testing.T) {

	t.Run("SQL string valuer", func(t *testing.T) {
		op := optl.Of("Hello optional")
		v, _ := op.Value()

		if v != "Hello optional" {
			t.Fail()
		}
	})

	t.Run("SQL nil string valuer", func(t *testing.T) {
		op := optl.Empty[string]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL int valuer", func(t *testing.T) {
		op := optl.Of(1)
		v, _ := op.Value()

		if v != int64(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil int valuer", func(t *testing.T) {
		op := optl.Empty[int]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL int8 valuer", func(t *testing.T) {
		op := optl.Of(int8(1))
		v, _ := op.Value()

		if v != int64(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil int8 valuer", func(t *testing.T) {
		op := optl.Empty[int8]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL int16 valuer", func(t *testing.T) {
		op := optl.Of(int16(1))
		v, _ := op.Value()

		if v != int64(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil int16 valuer", func(t *testing.T) {
		op := optl.Empty[int16]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL int32 valuer", func(t *testing.T) {
		op := optl.Of(int32(1))
		var op2 optl.Type[int32]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if op2.Get() != int32(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil int32 valuer", func(t *testing.T) {
		op := optl.Empty[int32]()
		op2 := optl.Of(int32(1))

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ int32) {
			t.Fail()
		})
	})

	t.Run("SQL int64 valuer", func(t *testing.T) {
		op := optl.Of(int64(1))
		v, _ := op.Value()

		if v != int64(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil int64 valuer", func(t *testing.T) {
		op := optl.Empty[int64]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL float32 valuer", func(t *testing.T) {
		op := optl.Of(float32(1))
		v, _ := op.Value()

		if v != float64(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil float32 valuer", func(t *testing.T) {
		op := optl.Empty[float32]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL float64 valuer", func(t *testing.T) {
		op := optl.Of(float64(1))
		v, _ := op.Value()

		if v != float64(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil float64 valuer", func(t *testing.T) {
		op := optl.Empty[float64]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL bool valuer", func(t *testing.T) {
		op := optl.Of(true)
		v, _ := op.Value()

		if v != true {
			t.Fail()
		}
	})

	t.Run("SQL nil bool valuer", func(t *testing.T) {
		op := optl.Empty[bool]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL time.Time valuer", func(t *testing.T) {
		ti := time.Now()
		op := optl.Of(ti)
		v, _ := op.Value()

		if !ti.Equal(v.(time.Time)) {
			t.Fail()
		}
	})

	t.Run("SQL nil time.Time valuer", func(t *testing.T) {
		op := optl.Empty[time.Time]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL array of bytes valuer", func(t *testing.T) {
		a := []byte{1, 2, 3}
		op := optl.Of(a)
		v, _ := op.Value()

		if !bytes.Equal(a, v.([]byte)) {
			t.Fail()
		}
	})

	t.Run("SQL nil array of bytes valuer", func(t *testing.T) {
		op := optl.Empty[[]byte]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL invalid array typer", func(t *testing.T) {
		op1 := optl.Of([]int{1, 2, 3})
		op2 := optl.Of([]string{"a", "b"})

		_, e := op1.Value()
		if e == nil {
			t.Fail()
		}
		_, e = op2.Value()
		if e == nil {
			t.Fail()
		}
	})

	t.Run("SQL invalid struct typer", func(t *testing.T) {
		op := optl.Of(dto{
			Field1: optl.Of("Hello optional"),
		})
		_, e := op.Value()
		if e == nil {
			t.Fail()
		}
	})
}

func TestM_Type_Scan(t *testing.T) {

	t.Run("SQL scanner string", func(t *testing.T) {
		op := optl.Empty[string]()
		op.Scan("Hello optional")
		if op.Get() != "Hello optional" {
			t.Fail()
		}
	})

	t.Run("SQL scanner nil string", func(t *testing.T) {
		op := optl.Of("s")
		op.Scan(nil)
		op.IfPresent(func(v string) {
			t.Fail()
		})
	})

	t.Run("SQL scanner int", func(t *testing.T) {
		op := optl.Empty[int]()
		op.Scan(int64(1))
		if op.Get() != 1 {
			t.Fail()
		}
	})

	t.Run("SQL scanner nil int", func(t *testing.T) {
		op := optl.Of(int(1))
		op.Scan(nil)
		op.IfPresent(func(v int) {
			t.Fail()
		})
	})

	t.Run("SQL scanner int8", func(t *testing.T) {
		op := optl.Empty[int8]()
		op.Scan(int64(1))
		if op.Get() != int8(1) {
			t.Fail()
		}
	})

	t.Run("SQL scanner nil int8", func(t *testing.T) {
		op := optl.Of(int8(1))
		op.Scan(nil)
		op.IfPresent(func(v int8) {
			t.Fail()
		})
	})

	t.Run("SQL scanner int16", func(t *testing.T) {
		op := optl.Empty[int16]()
		op.Scan(int64(1))
		if op.Get() != int16(1) {
			t.Fail()
		}
	})

	t.Run("SQL scanner nil int16", func(t *testing.T) {
		op := optl.Of(int16(1))
		op.Scan(nil)
		op.IfPresent(func(v int16) {
			t.Fail()
		})
	})

	t.Run("SQL scanner int32", func(t *testing.T) {
		op := optl.Empty[int32]()
		op.Scan(int64(1))
		if op.Get() != int32(1) {
			t.Fail()
		}
	})

	t.Run("SQL scanner nil int32", func(t *testing.T) {
		op := optl.Of(int32(1))
		op.Scan(nil)
		op.IfPresent(func(v int32) {
			t.Fail()
		})
	})

	t.Run("SQL scanner int64", func(t *testing.T) {
		op := optl.Empty[int16]()
		op.Scan(int64(1))
		if op.Get() != int16(1) {
			t.Fail()
		}
	})

	t.Run("SQL scanner nil int64", func(t *testing.T) {
		op := optl.Of(int64(1))
		op.Scan(nil)
		op.IfPresent(func(v int64) {
			t.Fail()
		})
	})

	t.Run("SQL scanner float32", func(t *testing.T) {
		op := optl.Empty[float32]()
		op.Scan(float64(1.5))
		if op.Get() != float32(1.5) {
			t.Fail()
		}
	})

	t.Run("SQL scanner nil float32", func(t *testing.T) {
		op := optl.Of(float32(1.5))
		op.Scan(nil)
		op.IfPresent(func(v float32) {
			t.Fail()
		})
	})

	t.Run("SQL scanner float64", func(t *testing.T) {
		op := optl.Empty[float64]()
		op.Scan(float64(1.5))
		if op.Get() != float64(1.5) {
			t.Fail()
		}
	})

	t.Run("SQL scanner nil float64", func(t *testing.T) {
		op := optl.Of(float64(1.5))
		op.Scan(nil)
		op.IfPresent(func(v float64) {
			t.Fail()
		})
	})

	t.Run("SQL scanner bool", func(t *testing.T) {
		op := optl.Empty[bool]()
		op.Scan(true)
		if op.Get() != true {
			t.Fail()
		}
	})

	t.Run("SQL scanner nil bool", func(t *testing.T) {
		op := optl.Of(true)
		op.Scan(nil)
		op.IfPresent(func(v bool) {
			t.Fail()
		})
	})

	t.Run("SQL scanner time.Time", func(t *testing.T) {
		op := optl.Empty[time.Time]()
		ti := time.Now()
		op.Scan(ti)
		if !ti.Equal(op.Get()) {
			t.Fail()
		}
	})

	t.Run("SQL scanner nil time.Time", func(t *testing.T) {
		op := optl.Of(time.Now())
		op.Scan(nil)
		op.IfPresent(func(v time.Time) {
			t.Fail()
		})
	})

	t.Run("SQL scanner array of bytes", func(t *testing.T) {
		op := optl.Empty[[]byte]()
		a := []byte{1, 2, 3}
		op.Scan(a)

		if !bytes.Equal(a, op.Get()) {
			t.Fail()
		}
	})

	t.Run("SQL scanner nil array of bytes", func(t *testing.T) {
		op := optl.Of([]byte{1, 2, 3})
		op.Scan(nil)
		op.IfPresent(func([]byte) {
			t.Fail()
		})
	})

}
