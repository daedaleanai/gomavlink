package mavlink

//lifted from
// https://github.com/mavlink/c_library_v2/blob/master/checksum.h#L25
// https://play.golang.org/p/ycYYW7bMChP
type crc16x25 uint16

func (acc *crc16x25) Update(b []byte) uint16 {
	for _, v := range b {
		t := v ^ byte(*acc)
		t ^= t << 4
		u := uint16(t)
		*acc = crc16x25(uint16(*acc)>>8 ^ u<<8 ^ u<<3 ^ u>>4)
	}
	return uint16(*acc)
}
