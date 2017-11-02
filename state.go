package chief

import (
	"encoding/binary"
	"io"
)

// State current state of a Pool
type State struct {
	Seed        uint64
	ShardLen    uint64
	ShardStart  uint64
	StripeLen   uint64
	StripeStart uint64
	Index       uint64
}

// DecodeState decode a state from a reader
func DecodeState(r io.Reader, s *State) (err error) {
	err = binary.Read(r, binary.BigEndian, s)
	return
}

// EncodeState encode a state from a reader
func EncodeState(w io.Writer, s *State) (err error) {
	err = binary.Write(w, binary.BigEndian, s)
	return
}
