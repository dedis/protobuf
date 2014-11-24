package protobuf

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	w := &bytes.Buffer{}
	err := GenerateProtobufDefinition(w, []interface{}{test{}}, nil)
	assert.NoError(t, err)
	expected := `
message test {
  optional bool boolean = 1;
  required sint64 i = 2;
  required sint32 i32 = 3;
  required sint64 i64 = 4;
  required uint32 u32 = 5;
  required uint64 u64 = 6;
  required sfixed32 sx32 = 7;
  required sfixed64 sx64 = 8;
  required fixed32 ux32 = 9;
  required ufixed64 ux64 = 10;
  required float f32 = 11;
  required double f64 = 12;
  required bytes bytes = 13;
  required string string = 14;
  required emb struct = 15;
  optional bool obool = 50;
  optional sint32 oi32 = 51;
  optional sint64 oi64 = 52;
  optional uint32 ou32 = 53;
  optional uint64 ou64 = 54;
  optional float of32 = 55;
  optional double of64 = 56;
  optional bytes obytes = 57;
  optional string ostring = 58;
  optional test ostruct = 59;
  repeated bool sbool = 100;
  repeated sint32 si32 = 101;
  repeated sint64 si64 = 102;
  repeated uint32 su32 = 103;
  repeated uint64 su64 = 104;
  repeated sfixed32 ssx32 = 105;
  repeated sfixed64 ssx64 = 106;
  repeated fixed32 sux32 = 107;
  repeated ufixed64 sux64 = 108;
  repeated float sf32 = 109;
  repeated double sf64 = 110;
  repeated bytes sbytes = 111;
  repeated string sstring = 112;
  repeated emb sstruct = 113;
}

`
	assert.Equal(t, expected, w.String())
}

func TestGeneratePersonExample(t *testing.T) {
	w := &bytes.Buffer{}
	err := GenerateProtobufDefinition(w, []interface{}{Person{}, PhoneNumber{}}, nil)
	assert.NoError(t, err)
	expected := `
message Person {
  required string name = 1;
  required sint32 id = 2;
  optional string email = 3;
  repeated PhoneNumber phone = 4;
}

message PhoneNumber {
  required string number = 1;
  optional uint32 type = 2;
}

`
	assert.Equal(t, expected, w.String())
}
