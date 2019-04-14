package mavlink

import "bufio"

type Decoder struct {
	r       bufio.Reader
	dialect func(uint32) Message
	// TODO make this map[StreamID]struct{stats}
	SeqNr  uint64 // number of messages received
	SeqGap uint64 // number of messages apparently dropped
}

/*
// NewDecoder constructs a new decoder from a reader and a dialect.
// The dialect is specified by a factory function that can construct empty Messages from a given Message Id.
// Mavgen generates the subpackage New functions for this purpose.
func NewDecoder(r io.Reader, dialect func(uint32) Message) *Decoder {
	return &Decoder{bufio.NewReader(r), dialect}
}

func (d *Decoder) Decode() (msg Message, str StreamID, err error) {

}
*/
