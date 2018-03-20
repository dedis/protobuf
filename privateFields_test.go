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

func TestPrivate(t *testing.T) {
	s := Private{37}
	u := Public{37}

	bufS, errS := Encode(&s)
	bufU, errU := Encode(&u)

	t.Log(bufS, errS)
	t.Log(bufU, errU)

	assert.Error(t, errS, "")
	assert.NoError(t, errU, "")
}