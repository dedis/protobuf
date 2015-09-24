package protobuf

// Message fields declared to have exactly this type
// will be transmitted as fixed-size 32-bit unsigned integers.
type Ufixed32 uint32

// Message fields declared to have exactly this type
// will be transmitted as fixed-size 64-bit unsigned integers.
type Ufixed64 uint64

// Message fields declared to have exactly this type
// will be transmitted as fixed-size 32-bit signed integers.
type Sfixed32 int32

// Message fields declared to have exactly this type
// will be transmitted as fixed-size 64-bit signed integers.
type Sfixed64 int64

// Protobufs enums are transmitted as unsigned varints;
// using this type alias is optional but recommended
// to ensure they get the correct type.
type Enum uint32

// Encoding is an object that generically represents the protobuf encoding,
// and may be used to set options such as a Constructor to use while decoding.
type Encoding struct {
	Constructor // How to instantiate unknown types while decoding

	// prevent clients from depending on the exact set of fields,
	// to reserve the right to extend in backward-compatible ways.
	foo struct{}
}

