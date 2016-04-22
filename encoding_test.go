package protobuf

import (
	"encoding"
	"testing"

	"github.com/stretchr/testify/assert"
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

// Validate that support for self-encoding via the Encoding
// interface works as expected
func TestBinaryMarshaler(t *testing.T) {
	wrapper := Wrapper{NewNumber(99)}
	buf, err := Encode(&wrapper)
	assert.Nil(t, err)

	wrapper2 := Wrapper{NewNumber(0)}
	err = Decode(buf, &wrapper2)

	assert.Nil(t, err)
	assert.Equal(t, 99, wrapper2.N.Value())
}

type Wrapper2 struct {
	W *Wrapper3
}

type Wrapper3 struct {
	A int
	B int
}

// MarshalBinary adds one to the data to check if we passed here
func (w *Wrapper3) MarshalBinary() ([]byte, error) {
	return []byte{byte(w.A + 1), byte(w.B + 1)}, nil
}

// UnmarshalBinary swaps the data to check if we passed here
func (w *Wrapper3) UnmarshalBinary(data []byte) error {
	w.A = int(data[1])
	w.B = int(data[0])
	return nil
}

var _ encoding.BinaryMarshaler = (*Wrapper3)(nil)

// Validate that support for self-encoding via the Encoding
// interface works as expected
func TestBinaryMarshalerStruct(t *testing.T) {
	// TODO: Encode a structure that supports self-encoding
	wrapper3 := &Wrapper3{A: 1, B: 4}
	buf, err := Encode(wrapper3)
	assert.Nil(t, err)

	wrapper4 := &Wrapper3{}
	err = Decode(buf, wrapper4)

	assert.Nil(t, err)
	assert.Equal(t, 5, wrapper4.A)
	assert.Equal(t, 2, wrapper4.B)

	// Working: a structure holding a structure that supports
	// self-encoding
	wrapper := Wrapper2{&Wrapper3{A: 1, B: 4}}
	buf, err = Encode(&wrapper)
	assert.Nil(t, err)

	wrapper2 := Wrapper2{&Wrapper3{}}
	err = Decode(buf, &wrapper2)

	assert.Nil(t, err)
	assert.Equal(t, 5, wrapper2.W.A)
	assert.Equal(t, 2, wrapper2.W.B)
}
