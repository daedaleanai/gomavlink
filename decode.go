package mavlink

import "io"

type Decoder struct {
	r       io.Reader
	dialect func(uint32) Message
}

func NewDecoder(r io.Reader, dialect func(uint32) Message) *Decoder {
	return &Decoder{r, dialect}
}
