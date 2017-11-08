package protobuf

import (
	"testing"

	"github.com/dedis/kyber"
	"github.com/dedis/kyber/group/edwards25519"
)

var Suite = edwards25519.NewAES128SHA256Ed25519()

type OneDimension struct {
	Points []kyber.Point
}

type TwoDimension struct {
	Points []OneDimension
}

type Message struct {
	Bytes []byte
}

func BenchmarkEncode(b *testing.B) {
	x := 100
	y := 10000
	big := &TwoDimension{
		Points: make([]OneDimension, x),
	}
	for i := range big.Points {
		big.Points[i].Points = make([]kyber.Point, y)
		for j := range big.Points[i].Points {
			big.Points[i].Points[j] = Suite.Point().Base()
		}

	}
	Msg := &Message{}
	var err error
	Msg.Bytes, err = Encode(big)
	if err != nil {
		b.Fatal(err)
	}
	if len(Msg.Bytes) < x*y {
		b.Fatal("Not enough data")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes, err := Encode(Msg)
		if err != nil {
			b.Fatal(err)
		}
		if len(bytes) < x*y {
			b.Fatal("Message not long enough")
		}
	}
}
