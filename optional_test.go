package optl_test

import (
	"bytes"
	"encoding/json"
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

func TestF_Of(t *testing.T) {

	t.Run("new optional string", func(t *testing.T) {
		op := optl.Of("Hello optional")
		if !op.IsPresent() || op.IsEmpty() || op.Get() != "Hello optional" {
			t.Fail()
		}
	})

	t.Run("new optional int", func(t *testing.T) {
		op := optl.Of(int(1))
		if !op.IsPresent() || op.IsEmpty() || op.Get() != int(1) {
			t.Fail()
		}
	})

	t.Run("new optional int8", func(t *testing.T) {
		op := optl.Of(int8(1))
		if !op.IsPresent() || op.IsEmpty() || op.Get() != int8(1) {
			t.Fail()
		}
	})

	t.Run("new optional int16", func(t *testing.T) {
		op := optl.Of(int16(1))
		if !op.IsPresent() || op.IsEmpty() || op.Get() != int16(1) {
			t.Fail()
		}
	})

	t.Run("new optional int32", func(t *testing.T) {
		op := optl.Of(int32(1))
		if !op.IsPresent() || op.IsEmpty() || op.Get() != int32(1) {
			t.Fail()
		}
	})

	t.Run("new optional int64", func(t *testing.T) {
		op := optl.Of(int64(1))
		if !op.IsPresent() || op.IsEmpty() || op.Get() != int64(1) {
			t.Fail()
		}
	})

	t.Run("new optional float32", func(t *testing.T) {
		op := optl.Of(float32(1))
		if !op.IsPresent() || op.IsEmpty() || op.Get() != float32(1) {
			t.Fail()
		}
	})

	t.Run("new optional float64", func(t *testing.T) {
		op := optl.Of(float64(1))
		if !op.IsPresent() || op.IsEmpty() || op.Get() != float64(1) {
			t.Fail()
		}
	})

	t.Run("new optional bool", func(t *testing.T) {
		op := optl.Of(true)
		if !op.IsPresent() || op.IsEmpty() || op.Get() != true {
			t.Fail()
		}
	})

	t.Run("new optional time.Time", func(t *testing.T) {
		ti := time.Now()
		op := optl.Of(ti)
		if !op.IsPresent() || op.IsEmpty() || op.Get() != ti {
			t.Fail()
		}
	})

	t.Run("new optional array", func(t *testing.T) {
		a := [...]byte{1, 2, 3}
		op := optl.Of(a)
		if !op.IsPresent() || op.IsEmpty() || op.Get()[1] != 2 {
			t.Fail()
		}
	})

	t.Run("new optional slice", func(t *testing.T) {
		a := [...]byte{1, 2, 3}
		op := optl.Of(a[1:2])
		if !op.IsPresent() || op.IsEmpty() || op.Get()[0] != 2 {
			t.Fail()
		}
	})

	t.Run("new optional struct", func(t *testing.T) {
		dto := optl.Of(dto{
			Field1: optl.Of("text..."),
		})

		if !dto.IsPresent() || dto.IsEmpty() || !dto.Get().Field1.IsPresent() || dto.Get().Field1.Get() != "text..." {
			t.Fail()
		}

		if !dto.IsPresent() || dto.IsEmpty() || dto.Get().Field2.IsPresent() {
			t.Fail()
		}
	})
}

func TestF_OfNillable(t *testing.T) {

	t.Run("new optional nillable nil", func(t *testing.T) {
		var s *string = nil
		op := optl.OfNillable(s)
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
	})

	t.Run("new optional nillable not nil", func(t *testing.T) {
		var s string = "Hello optional"
		op := optl.OfNillable(&s)
		if !op.IsPresent() || op.IsEmpty() || op.Get() != s {
			t.Fail()
		}
	})

}

func TestF_Empty(t *testing.T) {
	t.Run("new optional empty string", func(t *testing.T) {
		op := optl.Empty[string]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})

	t.Run("new optional empty int", func(t *testing.T) {
		op := optl.Empty[int]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})

	t.Run("new optional empty int8", func(t *testing.T) {
		op := optl.Empty[int8]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})

	t.Run("new optional empty int16", func(t *testing.T) {
		op := optl.Empty[int16]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})

	t.Run("new optional empty int32", func(t *testing.T) {
		op := optl.Empty[int32]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})

	t.Run("new optional empty int64", func(t *testing.T) {
		op := optl.Empty[int64]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})

	t.Run("new optional empty float32", func(t *testing.T) {
		op := optl.Empty[float32]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})

	t.Run("new optional empty float64", func(t *testing.T) {
		op := optl.Empty[float64]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})

	t.Run("new optional empty bool", func(t *testing.T) {
		op := optl.Empty[bool]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})

	t.Run("new optional empty time.Time", func(t *testing.T) {
		op := optl.Empty[time.Time]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})

	t.Run("new optional empty array", func(t *testing.T) {
		op := optl.Empty[[]byte]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})

	t.Run("new optional empty struct", func(t *testing.T) {
		op := optl.Empty[dto]()
		if op.IsPresent() || !op.IsEmpty() {
			t.Fail()
		}
		assertPanic(t, op.Get)
	})
}

func TestM_Type_OrElse(t *testing.T) {
	t.Run("OrElse with value present", func(t *testing.T) {
		if optl.Of("Hello optional").OrElse("Hello else") != "Hello optional" {
			t.Fail()
		}
	})

	t.Run("OrElse with value not present", func(t *testing.T) {
		if optl.Empty[string]().OrElse("Hello else") != "Hello else" {
			t.Fail()
		}
	})
}

func TestM_Type_OrElseGet(t *testing.T) {
	t.Run("OrElseGet with value present", func(t *testing.T) {
		if optl.Of("Hello optional").OrElseGet(func() string { return "Hello else" }) != "Hello optional" {
			t.Fail()
		}
	})

	t.Run("OrElseGet with value not present", func(t *testing.T) {
		if optl.Empty[string]().OrElseGet(func() string { return "Hello else" }) != "Hello else" {
			t.Fail()
		}
	})
}

func TestM_Type_IfPresent(t *testing.T) {
	t.Run("IfPresent with value present", func(t *testing.T) {
		execute := false
		optl.Of("Hello optional").IfPresent(func(v string) {
			if v != "Hello optional" {
				t.Fail()
			}
			execute = true
		})
		if !execute {
			t.Fail()
		}
	})

	t.Run("IfPresent with value not present", func(t *testing.T) {
		optl.Empty[string]().IfPresent(func(v string) {
			t.Fail()
		})
	})
}

func TestM_Type_Filter(t *testing.T) {
	t.Run("Filter off", func(t *testing.T) {
		op := optl.Of("Hello optional").Filter(func(v string) bool {
			if v != "Hello optional" {
				t.Fail()
			}
			return true
		})
		if op.Get() != "Hello optional" {
			t.Fail()
		}
	})

	t.Run("Filter on", func(t *testing.T) {
		op := optl.Of("Hello optional").Filter(func(v string) bool {
			return false
		})
		if op.IsPresent() {
			t.Fail()
		}
	})

	t.Run("Filter empty", func(t *testing.T) {
		optl.Empty[string]().Filter(func(v string) bool {
			t.Fail()
			return false
		})
	})

}

func TestM_Type_MarshallingJSON(t *testing.T) {

	t.Run("Marshalling string", func(t *testing.T) {
		op := optl.Of("Hello optional")
		var op2 optl.Type[string]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if op2.Get() != "Hello optional" {
			t.Fail()
		}
	})

	t.Run("Marshalling nil string", func(t *testing.T) {
		op := optl.Empty[string]()
		op2 := optl.Of("Hello optional")

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ string) {
			t.Fail()
		})
	})

	t.Run("Marshalling int", func(t *testing.T) {
		op := optl.Of(1)
		var op2 optl.Type[int]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if op2.Get() != 1 {
			t.Fail()
		}
	})

	t.Run("Marshalling nil int", func(t *testing.T) {
		op := optl.Empty[int]()
		op2 := optl.Of(1)

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ int) {
			t.Fail()
		})
	})

	t.Run("Marshalling int8", func(t *testing.T) {
		op := optl.Of(int8(1))
		var op2 optl.Type[int8]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if op2.Get() != int8(1) {
			t.Fail()
		}
	})

	t.Run("Marshalling nil int8", func(t *testing.T) {
		op := optl.Empty[int8]()
		op2 := optl.Of(int8(1))

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ int8) {
			t.Fail()
		})
	})

	t.Run("Marshalling int16", func(t *testing.T) {
		op := optl.Of(int16(1))
		var op2 optl.Type[int16]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if op2.Get() != int16(1) {
			t.Fail()
		}
	})

	t.Run("Marshalling nil int16", func(t *testing.T) {
		op := optl.Empty[int16]()
		op2 := optl.Of(int16(1))

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ int16) {
			t.Fail()
		})
	})

	t.Run("Marshalling int32", func(t *testing.T) {
		op := optl.Of(int32(1))
		var op2 optl.Type[int32]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if op2.Get() != int32(1) {
			t.Fail()
		}
	})

	t.Run("Marshalling nil int32", func(t *testing.T) {
		op := optl.Empty[int32]()
		op2 := optl.Of(int32(1))

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ int32) {
			t.Fail()
		})
	})

	t.Run("Marshalling int64", func(t *testing.T) {
		op := optl.Of(int64(1))
		var op2 optl.Type[int64]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if op2.Get() != int64(1) {
			t.Fail()
		}
	})

	t.Run("Marshalling nil int64", func(t *testing.T) {
		op := optl.Empty[int64]()
		op2 := optl.Of(int64(1))

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ int64) {
			t.Fail()
		})
	})

	t.Run("Marshalling float32", func(t *testing.T) {
		op := optl.Of(float32(1.5))
		var op2 optl.Type[float32]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if op2.Get() != float32(1.5) {
			t.Fail()
		}
	})

	t.Run("Marshalling nil float32", func(t *testing.T) {
		op := optl.Empty[float32]()
		op2 := optl.Of(float32(1.5))

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ float32) {
			t.Fail()
		})
	})

	t.Run("Marshalling float64", func(t *testing.T) {
		op := optl.Of(float64(1.5))
		var op2 optl.Type[float64]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if op2.Get() != float64(1.5) {
			t.Fail()
		}
	})

	t.Run("Marshalling nil float64", func(t *testing.T) {
		op := optl.Empty[float64]()
		op2 := optl.Of(float64(1.5))

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ float64) {
			t.Fail()
		})
	})

	t.Run("Marshalling bool", func(t *testing.T) {
		op := optl.Of(true)
		var op2 optl.Type[bool]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if op2.Get() != true {
			t.Fail()
		}
	})

	t.Run("Marshalling nil bool", func(t *testing.T) {
		op := optl.Empty[bool]()
		op2 := optl.Of(true)

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ bool) {
			t.Fail()
		})
	})

	t.Run("Marshalling time.Time", func(t *testing.T) {
		ti := time.Now()
		op := optl.Of(ti)
		var op2 optl.Type[time.Time]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if !op2.Get().Equal(ti) {
			t.Fail()
		}
	})

	t.Run("Marshalling nil time.Time", func(t *testing.T) {
		op := optl.Empty[time.Time]()
		op2 := optl.Of(time.Now())

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ time.Time) {
			t.Fail()
		})
	})

	t.Run("Marshalling array", func(t *testing.T) {
		a := []byte{1, 2, 3}
		op := optl.Of(a)
		var op2 optl.Type[[]byte]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if !bytes.Equal(op2.Get(), a) {
			t.Fail()
		}
	})

	t.Run("Marshalling nil array", func(t *testing.T) {
		op := optl.Empty[[]byte]()
		op2 := optl.Of([]byte{1, 2, 3})

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ []byte) {
			t.Fail()
		})
	})

	t.Run("Marshalling slice", func(t *testing.T) {
		a := []byte{1, 2, 3}
		op := optl.Of(a[1:2])
		var op2 optl.Type[[]byte]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if len(op2.Get()) != 1 || op2.Get()[0] != 2 {
			t.Fail()
		}
	})

	t.Run("Marshalling struct", func(t *testing.T) {
		dto1 := optl.Of(dto{
			Field1: optl.Of("text..."),
			Field2: optl.Of(int32(1)),
		})

		dto2 := optl.Empty[dto]()

		v, _ := json.Marshal(dto1)
		json.Unmarshal(v, &dto2)

		if dto1.Get().Field1.Get() != dto2.Get().Field1.Get() {
			t.Fail()
		}

		if dto1.Get().Field2.Get() != dto2.Get().Field2.Get() {
			t.Fail()
		}

	})

	t.Run("Marshalling struct with nil", func(t *testing.T) {
		dto1 := optl.Of(dto{
			Field1: optl.Of("text..."),
			Field2: optl.Empty[int32](),
		})

		dto2 := optl.Empty[dto]()

		v, _ := json.Marshal(dto1)
		json.Unmarshal(v, &dto2)

		if dto1.Get().Field1.Get() != dto2.Get().Field1.Get() {
			t.Fail()
		}

		if dto2.Get().Field2.IsPresent() {
			t.Fail()
		}

	})

	t.Run("Marshalling nil struct", func(t *testing.T) {

		dto1 := optl.Empty[dto]()

		dto2 := optl.Of(dto{
			Field1: optl.Of("text..."),
			Field2: optl.Empty[int32](),
		})

		v, _ := json.Marshal(dto1)
		json.Unmarshal(v, &dto2)

		dto2.IfPresent(func(d dto) {
			t.Fail()
		})

	})

}

func TestM_Type_Value(t *testing.T) {

	t.Run("SQL string value", func(t *testing.T) {
		op := optl.Of("Hello optional")
		v, _ := op.Value()

		if v != "Hello optional" {
			t.Fail()
		}
	})

	t.Run("SQL nil string value", func(t *testing.T) {
		op := optl.Empty[string]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL int value", func(t *testing.T) {
		op := optl.Of(1)
		v, _ := op.Value()

		if v != int64(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil int value", func(t *testing.T) {
		op := optl.Empty[int]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL int8 value", func(t *testing.T) {
		op := optl.Of(int8(1))
		v, _ := op.Value()

		if v != int64(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil int8 value", func(t *testing.T) {
		op := optl.Empty[int8]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL int16 value", func(t *testing.T) {
		op := optl.Of(int16(1))
		v, _ := op.Value()

		if v != int64(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil int16 value", func(t *testing.T) {
		op := optl.Empty[int16]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL int32 value", func(t *testing.T) {
		op := optl.Of(int32(1))
		var op2 optl.Type[int32]

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		if op2.Get() != int32(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil int32 value", func(t *testing.T) {
		op := optl.Empty[int32]()
		op2 := optl.Of(int32(1))

		v, _ := json.Marshal(op)
		json.Unmarshal(v, &op2)

		op2.IfPresent(func(_ int32) {
			t.Fail()
		})
	})

	t.Run("SQL int64 value", func(t *testing.T) {
		op := optl.Of(int64(1))
		v, _ := op.Value()

		if v != int64(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil int64 value", func(t *testing.T) {
		op := optl.Empty[int64]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL float32 value", func(t *testing.T) {
		op := optl.Of(float32(1))
		v, _ := op.Value()

		if v != float64(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil float32 value", func(t *testing.T) {
		op := optl.Empty[float32]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL float64 value", func(t *testing.T) {
		op := optl.Of(float64(1))
		v, _ := op.Value()

		if v != float64(1) {
			t.Fail()
		}
	})

	t.Run("SQL nil float64 value", func(t *testing.T) {
		op := optl.Empty[float64]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL bool value", func(t *testing.T) {
		op := optl.Of(true)
		v, _ := op.Value()

		if v != true {
			t.Fail()
		}
	})

	t.Run("SQL nil bool value", func(t *testing.T) {
		op := optl.Empty[bool]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL time.Time value", func(t *testing.T) {
		ti := time.Now()
		op := optl.Of(ti)
		v, _ := op.Value()

		if !ti.Equal(v.(time.Time)) {
			t.Fail()
		}
	})

	t.Run("SQL nil time.Time value", func(t *testing.T) {
		op := optl.Empty[time.Time]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL array of bytes value", func(t *testing.T) {
		a := []byte{1, 2, 3}
		op := optl.Of(a)
		v, _ := op.Value()

		if !bytes.Equal(a, v.([]byte)) {
			t.Fail()
		}
	})

	t.Run("SQL nil array of bytes value", func(t *testing.T) {
		op := optl.Empty[[]byte]()
		v, _ := op.Value()

		if v != nil {
			t.Fail()
		}
	})

	t.Run("SQL invalid array type", func(t *testing.T) {
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

	t.Run("SQL invalid struct type", func(t *testing.T) {
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

	t.Run("SQL scan string", func(t *testing.T) {
		op := optl.Empty[string]()
		op.Scan("Hello optional")
		if op.Get() != "Hello optional" {
			t.Fail()
		}
	})

	t.Run("SQL scan nil string", func(t *testing.T) {
		op := optl.Of("s")
		op.Scan(nil)
		op.IfPresent(func(v string) {
			t.Fail()
		})
	})

	t.Run("SQL scan int", func(t *testing.T) {
		op := optl.Empty[int]()
		op.Scan(int64(1))
		if op.Get() != 1 {
			t.Fail()
		}
	})

	t.Run("SQL scan nil int", func(t *testing.T) {
		op := optl.Of(int(1))
		op.Scan(nil)
		op.IfPresent(func(v int) {
			t.Fail()
		})
	})

	t.Run("SQL scan int8", func(t *testing.T) {
		op := optl.Empty[int8]()
		op.Scan(int64(1))
		if op.Get() != int8(1) {
			t.Fail()
		}
	})

	t.Run("SQL scan nil int8", func(t *testing.T) {
		op := optl.Of(int8(1))
		op.Scan(nil)
		op.IfPresent(func(v int8) {
			t.Fail()
		})
	})

	t.Run("SQL scan int16", func(t *testing.T) {
		op := optl.Empty[int16]()
		op.Scan(int64(1))
		if op.Get() != int16(1) {
			t.Fail()
		}
	})

	t.Run("SQL scan nil int16", func(t *testing.T) {
		op := optl.Of(int16(1))
		op.Scan(nil)
		op.IfPresent(func(v int16) {
			t.Fail()
		})
	})

	t.Run("SQL scan int32", func(t *testing.T) {
		op := optl.Empty[int32]()
		op.Scan(int64(1))
		if op.Get() != int32(1) {
			t.Fail()
		}
	})

	t.Run("SQL scan nil int32", func(t *testing.T) {
		op := optl.Of(int32(1))
		op.Scan(nil)
		op.IfPresent(func(v int32) {
			t.Fail()
		})
	})

	t.Run("SQL scan int64", func(t *testing.T) {
		op := optl.Empty[int16]()
		op.Scan(int64(1))
		if op.Get() != int16(1) {
			t.Fail()
		}
	})

	t.Run("SQL scan nil int64", func(t *testing.T) {
		op := optl.Of(int64(1))
		op.Scan(nil)
		op.IfPresent(func(v int64) {
			t.Fail()
		})
	})

	t.Run("SQL scan float32", func(t *testing.T) {
		op := optl.Empty[float32]()
		op.Scan(float64(1.5))
		if op.Get() != float32(1.5) {
			t.Fail()
		}
	})

	t.Run("SQL scan nil float32", func(t *testing.T) {
		op := optl.Of(float32(1.5))
		op.Scan(nil)
		op.IfPresent(func(v float32) {
			t.Fail()
		})
	})

	t.Run("SQL scan float64", func(t *testing.T) {
		op := optl.Empty[float64]()
		op.Scan(float64(1.5))
		if op.Get() != float64(1.5) {
			t.Fail()
		}
	})

	t.Run("SQL scan nil float64", func(t *testing.T) {
		op := optl.Of(float64(1.5))
		op.Scan(nil)
		op.IfPresent(func(v float64) {
			t.Fail()
		})
	})

	t.Run("SQL scan bool", func(t *testing.T) {
		op := optl.Empty[bool]()
		op.Scan(true)
		if op.Get() != true {
			t.Fail()
		}
	})

	t.Run("SQL scan nil bool", func(t *testing.T) {
		op := optl.Of(true)
		op.Scan(nil)
		op.IfPresent(func(v bool) {
			t.Fail()
		})
	})

	t.Run("SQL scan time.Time", func(t *testing.T) {
		op := optl.Empty[time.Time]()
		ti := time.Now()
		op.Scan(ti)
		if !ti.Equal(op.Get()) {
			t.Fail()
		}
	})

	t.Run("SQL scan nil time.Time", func(t *testing.T) {
		op := optl.Of(time.Now())
		op.Scan(nil)
		op.IfPresent(func(v time.Time) {
			t.Fail()
		})
	})

	t.Run("SQL scan array of bytes", func(t *testing.T) {
		op := optl.Empty[[]byte]()
		a := []byte{1, 2, 3}
		op.Scan(a)

		if !bytes.Equal(a, op.Get()) {
			t.Fail()
		}
	})

	t.Run("SQL scan nil array of bytes", func(t *testing.T) {
		op := optl.Of([]byte{1, 2, 3})
		op.Scan(nil)
		op.IfPresent(func([]byte) {
			t.Fail()
		})
	})

}
