package protobuf

import (
	"fmt"
	"testing"
)

type Msg struct {
	Data [][]string
}

func Test_PriFiRocks_4(t *testing.T) {

	data := make([][]string, 1)
	data[0] = make([]string, 2)
	data[0][0] = "PRIFI"
	data[0][1] = "ROCKS"

	source := &Msg{data}

	fmt.Println(*source) // [[PRIFI ROCKS]]

	b, err := Encode(source)
	if err != nil {
		t.Error("Encode returned an error,",err)
	}

	decoded := &Msg{}
	err = Decode(b, decoded)
	if err != nil {
		t.Error("Decode returned an error,",err)
	}

	fmt.Println(*decoded) // [[PRIFI] [ROCKS]]

	if len(source.Data) != len(decoded.Data) {
		t.Error("Length on first dimension don't match", len(source.Data), "!=", len(decoded.Data))
		t.Error("	", source.Data, "!=", decoded.Data)
	}
	if len(source.Data[0]) != len(decoded.Data[0]) {
		t.Error("Length on first dimension don't match", len(source.Data[0]), "!=", len(decoded.Data[0]))
		t.Error("	", source.Data[0], "!=", decoded.Data[0])
	}
}
