package protobuf

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type Private struct {
	a int
}

type Public struct {
	A int
}

type Empty struct {
	Empty *string
}

func TestPrivate(t *testing.T) {
	s := Private{37}
	u := Public{37}
	str := "b"
	e := Empty{&str}

	bufS, errS := Encode(&s)
	bufU, errU := Encode(&u)
	bufE, errE := Encode(&e)

	t.Log(bufS, errS)
	t.Log(bufU, errU)
	t.Log(bufE, errE)

	assert.Error(t, errS)
	assert.NoError(t, errU)
	assert.NoError(t, errE)
}