package tlog

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"time"

	mavlink "github.com/daedaleanai/gomavlink"
	"github.com/daedaleanai/gomavlink/ardupilotmega"
)

// A Decoder decodes the (big endian 100ns since unix epoch, ardupilot mavlink message) pairs of a tlog file.
type Decoder struct {
	r *bufio.Reader
	d *mavlink.Decoder
}

// NewDecoder constructs a new decoder from a reader and a dialect.
// The dialect is specified by a factory function that can construct empty Messages from a given Message Id.
// Mavgen generates the subpackage Dialect functions for this purpose.
func NewDecoder(r io.Reader) *Decoder {
	rr := bufio.NewReaderSize(r, 300)
	// this relies on mavlinkDecoder reusing rr
	return &Decoder{rr, mavlink.NewDecoder(rr, ardupilotmega.Dialect)}
}

// A Record is a Time, a StreamID and a Message
type Record struct {
	time.Time
	mavlink.StreamID
	mavlink.Message
}

func (r *Record) String() string {
	return fmt.Sprintf("%s %v %#v", r.Time.Format("2006-01-02T15:04:05.999999999"), r.StreamID, r.Message)
}

func (d *Decoder) Decode() (rec Record, err error) {
	var ts uint64
	err = binary.Read(d.r, binary.BigEndian, &ts)
	if err != nil {
		return
	}
	rec.Time = time.Unix(0, int64(1000*ts))
	rec.Message, rec.StreamID, err = d.d.Decode()
	return
}
