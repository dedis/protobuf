package protobuf

import (
	"encoding"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type Number interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler

	Value() int
}

type Int struct {
	N int
}

type Wrapper struct {
	N Number
}

func NewNumber(n int) Number {
	return &Int{n}
}

func (i *Int) Value() int {
	return i.N
}

func (i *Int) MarshalBinary() ([]byte, error) {
	return []byte{byte(i.N)}, nil
}

func (i *Int) UnmarshalBinary(data []byte) error {
	i.N = int(data[0])
	return nil
}

// Check at compile time that we satisfy the interfaces.
var _ encoding.BinaryMarshaler = (*Int)(nil)
var _ encoding.BinaryUnmarshaler = (*Int)(nil)

// Validate that support for self-encoding via BinaryMarshaler
// works as expected.
func TestBinaryMarshaler(t *testing.T) {
	w := Wrapper{NewNumber(99)}
	_, err := Encode(&w)
	assert.Nil(t, err)
}

var testNumber = NewNumber(99)
var testBuf, _ = Encode(&Wrapper{testNumber})

func TestBinaryUnmarshaler(t *testing.T) {
	w2 := Wrapper{NewNumber(0)}
	err := Decode(testBuf, &w2)

	assert.Nil(t, err)
	assert.Equal(t, testNumber.Value(), w2.N.Value())
}

var aNumber Number
var tNumber = reflect.TypeOf(&aNumber).Elem()

type testCons struct{}

func (_ testCons) New(t reflect.Type) interface{} {
	switch t {
	case tNumber:
		return new(Int)
	default:
		return nil
	}
}

func TestBinaryUnmarshalerWithCons(t *testing.T) {
	w2 := new(Wrapper)
	e := Encoding{ Constructor: &testCons{}}
	err := e.Decode(testBuf, w2)

	assert.Nil(t, err)
	assert.Equal(t, testNumber.Value(), w2.N.Value())
}
