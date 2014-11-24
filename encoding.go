package protobuf

/*
Encoding is a basic interface representing fixed-length (or known-length)
cryptographic objects or structures having a built-in binary encoding.
*/
type Encoding interface {
	String() string

	// Encoded length of this object in bytes.
	Len() int

	// Encode the content of this object into a slice,
	// whose length must be exactly Len().
	Encode() []byte

	// XXX EncodeTo(w io.Writer) error
	// XXX WriteTo(w io.Writer) (n,error)

	// Decode the content of this object from a slice,
	// whose length must be exactly Len().
	Decode(buf []byte) error

	// Decode the content of this object by reading from an io.Reader.
	// If r is also a cipher.Stream (e.g., a RandomReader),
	// then picks a valid object [pseudo-]randomly from that stream,
	// which may entail reading more than Len bytes due to retries.
	// XXX DecodeFrom(r io.Reader) error
	// XXX ReadFrom(w io.Writer) (n,error)
}
