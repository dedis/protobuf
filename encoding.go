package protobuf

// Encoding is an object that generically represents the protobuf encoding,
// and may be used to set options such as a Constructor to use while decoding.
type Encoding struct {
	cons Constructor
}

func (e Encoding) SetConstructor(cons Constructor) Encoding {
	e.cons = cons
	return e
}
