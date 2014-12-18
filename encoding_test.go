package protobuf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	Encoding

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

func (i *Int) String() string {
	return fmt.Sprintf("Int: %d", i.N)
}

func (i *Int) Len() int {
	return 1
}

func (i *Int) Encode() []byte {
	return []byte{byte(i.N)}
}

func (i *Int) Decode(data []byte) error {
	i.N = int(data[0])
	return nil
}

var _ Encoding = (*Int)(nil)

// Validate that support for self-encoding via the Encoding
// interface works as expected
func TestEncoding(t *testing.T) {
	wrapper := Wrapper{NewNumber(99)}
	buf, err := Encode(&wrapper)
	assert.Nil(t, err)

	wrapper2 := Wrapper{NewNumber(0)}
	err = Decode(buf, &wrapper2)

	assert.Nil(t, err)
	assert.Equal(t, 99, wrapper2.N.Value())
}
