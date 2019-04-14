package mavlink

import (
	"fmt"
	"io"
)

// An Encoder can serialize Messages on a io.Writer
type Encoder struct {
	// w is the io.Writer that Encode serializes to.  Each frame will be written in a single call to w.Write.
	w io.Writer
	// Protocol switches between the formats used when serializing a message.
	// This can be reset freely between calls to Encode.
	Protocol Protocol
	// SeqNr counts the number of messages sent, the lower 8 bits are used to generate the packet sequence number.
	// This can be reset freely between calls to Encode.
	SeqNr uint64
	// The value of CompatFlags will be copied to the 3rd byte of the V2 messages.
	CompatFlags byte
}

func NewEncoder(w io.Writer) *Encoder { return &Encoder{w: w} }

func (e *Encoder) Encode(sysid, compid byte, m Message) error {
	buf := make([]byte, 279)
	mid := m.ID()
	switch e.Protocol {
	case V2:
		buf = append(buf, 0xFD, 0, 0, e.CompatFlags, byte(e.SeqNr), sysid, compid, byte(mid), byte(mid>>8), byte(mid>>16))
	case V2Signed:
		buf = append(buf, 0xFD, 0, 1, e.CompatFlags, byte(e.SeqNr), sysid, compid, byte(mid), byte(mid>>8), byte(mid>>16))
	case V1:
		if mid > 255 {
			return fmt.Errorf("Cannot encode %T as V1 frame, message id %d too large.", m, mid)
		}
		buf = append(buf, 0xFE, 0, byte(e.SeqNr), sysid, compid, byte(mid))
	}

	mark := len(buf)
	if e.Protocol == V1 {
		buf = m.MarshalV1(buf)
	} else {
		buf = m.MarshalV2(buf)
		// chop trailing zeroes
		for len(buf) > 0 && buf[len(buf)-1] == 0 {
			buf = buf[:len(buf)-1]
		}
	}
	if len(buf)-mark > 255 {
		return fmt.Errorf("Cannot encode %T, payload %d bytes > 255.", m, len(buf)-mark)
	}
	buf[1] = byte(len(buf) - mark)

	x := crc16x25(0xffff)
	x.Update(buf[1:])
	c := m.CRCExtra()
	x.Update([]byte{byte(c), byte(c >> 8)})
	buf = append(buf, byte(x), byte(x>>8))

	if e.Protocol == V2Signed {
		// TODO(lvd) append signature
		return fmt.Errorf("sorry, didn't get around to implementing signing yet.")
	}

	_, err := e.w.Write(buf)

	return err
}
