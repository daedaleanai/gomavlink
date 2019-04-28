package mavlink

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
)

// A Decoder can deserialize Messages from a io.Reader
type Decoder struct {
	r       *bufio.Reader
	dialect func(int) Message
}

// NewDecoder constructs a new decoder from a reader and a dialect.
// The dialect is specified by a factory function that can construct empty Messages from a given Message Id.
// Mavgen generates the subpackage Dialect functions for this purpose.
func NewDecoder(r io.Reader, dialect func(int) Message) *Decoder {
	return &Decoder{bufio.NewReaderSize(r, 280), dialect}
}

// Resync discards bytes until the next byte is the V1(0xfe) or V2 (0xfd) STX marker.
// It returns the number of bytes discarded and any error encountered calling ReadByte
func (d *Decoder) Resync() (n int, err error) {
	var b byte
	for {
		b, err = d.r.ReadByte()
		if err != nil {
			return
		}
		n++
		if b == 0xFD || b == 0xFE {
			err = d.r.UnreadByte()
			break
		}
	}
	return
}

var (
	ErrMustSync = errors.New("Stream must resync")
	ErrBadCRC   = errors.New("Bad x25 CRC")
)

func (d *Decoder) Decode() (msg Message, str StreamID, err error) {

	var (
		stx, paylen, incomp byte
		msgId               int
		pld                 = make([]byte, 256) // zero filled
		chksum              [2]byte
		signature           [13]byte
		crc                 crc16x25
	)

	defer func() {
		if r, _ := recover().(error); r != nil {
			err = r
		}
	}()

	get := func() byte {
		b, err := d.r.ReadByte()
		if err != nil {
			panic(err)
		}
		crc.Update([]byte{b})
		return b
	}

	stx = get()
	switch stx {
	case 0xFD, 0xFE:
		// nix
	default:
		log.Println("stx", stx)
		return nil, 0, ErrMustSync
	}

	crc = 0xffff
	paylen = get()
	if stx == 0xFD {
		incomp = get()
		_ = get() // compat field, not used by anyone
	}
	_ = get() // seq field not used yet, TODO use for stats
	sysid := get()
	compid := get()
	str = Stream(sysid, compid, 0)
	msgId = int(get())
	if stx == 0xFD {
		msgId |= int(get()) << 8
		msgId |= int(get()) << 16
	}

	pld = pld[:paylen]
	if _, err = io.ReadFull(d.r, pld); err != nil {
		panic(err)
	}
	crc.Update(pld)

	if _, err = io.ReadFull(d.r, chksum[:]); err != nil {
		panic(err)
	}

	if incomp&1 != 0 {
		if _, err = io.ReadFull(d.r, signature[:]); err != nil {
			panic(err)
		}

		// TODO check signature and set Link in strId
	}

	msg = d.dialect(msgId)

	if msg != nil {
		crc.Update([]byte{msg.CRCExtra()})
	}
	if chk := uint16(chksum[1])<<8 | uint16(chksum[0]); uint16(crc) != chk {
		panic(ErrBadCRC)
	}

	switch {
	case msg == nil:
		panic(fmt.Errorf("Cannot decode message type id %d", msgId))
	case stx == 0xFE:
		//log.Printf("Decoding %T %d bytes (V1) %v", msg, len(pld), pld)
		msg.UnmarshalV1(pld)
	case stx == 0xFD:
		//log.Printf("Decoding %T %d bytes (V2) %v", msg, len(pld), pld)
		msg.UnmarshalV2(pld[:256]) // zero padded
	}

	return
}
