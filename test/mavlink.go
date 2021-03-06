// Generated enums and structures for Mavlink dialect test #0 version 3
package test

// Generated by gomavlink, DO NOT EDIT.

import (
	"math"

	mavlink "github.com/daedaleanai/gomavlink"
)

// The Dialect factory function constructs the proper empty message given the message ID.
func Dialect(mid int) mavlink.Message {
	switch mid {
	case 0:
		return &TestTypes{}
	}
	return nil
}

/* Test all field types */
type TestTypes struct {
	/* uint64_t */
	U64 uint64

	/* int64_t */
	S64 int64

	/* double */
	D float64

	/* uint64_t_array */
	U64Array [3]uint64

	/* int64_t_array */
	S64Array [3]int64

	/* double_array */
	DArray [3]float64

	/* uint32_t */
	U32 uint32

	/* int32_t */
	S32 int32

	/* float */
	F float32

	/* uint32_t_array */
	U32Array [3]uint32

	/* int32_t_array */
	S32Array [3]int32

	/* float_array */
	FArray [3]float32

	/* uint16_t */
	U16 uint16

	/* int16_t */
	S16 int16

	/* uint16_t_array */
	U16Array [3]uint16

	/* int16_t_array */
	S16Array [3]int16

	/* char */
	C byte

	/* string */
	S [10]byte

	/* uint8_t */
	U8 byte

	/* int8_t */
	S8 int8

	/* uint8_t_array */
	U8Array [3]byte

	/* int8_t_array */
	S8Array [3]int8
}

func (m *TestTypes) ID() int        { return 0 }
func (m *TestTypes) CRCExtra() byte { return 103 }

func (m *TestTypes) MarshalV1(buf []byte) []byte {
	buf = marshalUint64(buf, (m.U64))
	buf = marshalInt64(buf, (m.S64))
	buf = marshalFloat64(buf, (m.D))
	for _, v := range m.U64Array {
		buf = marshalUint64(buf, (v))
	}
	for _, v := range m.S64Array {
		buf = marshalInt64(buf, (v))
	}
	for _, v := range m.DArray {
		buf = marshalFloat64(buf, (v))
	}
	buf = marshalUint32(buf, (m.U32))
	buf = marshalInt32(buf, (m.S32))
	buf = marshalFloat32(buf, (m.F))
	for _, v := range m.U32Array {
		buf = marshalUint32(buf, (v))
	}
	for _, v := range m.S32Array {
		buf = marshalInt32(buf, (v))
	}
	for _, v := range m.FArray {
		buf = marshalFloat32(buf, (v))
	}
	buf = marshalUint16(buf, (m.U16))
	buf = marshalInt16(buf, (m.S16))
	for _, v := range m.U16Array {
		buf = marshalUint16(buf, (v))
	}
	for _, v := range m.S16Array {
		buf = marshalInt16(buf, (v))
	}
	buf = marshalByte(buf, (m.C))
	for _, v := range m.S {
		buf = marshalByte(buf, (v))
	}
	buf = marshalByte(buf, (m.U8))
	buf = marshalInt8(buf, (m.S8))
	for _, v := range m.U8Array {
		buf = marshalByte(buf, (v))
	}
	for _, v := range m.S8Array {
		buf = marshalInt8(buf, (v))
	}

	return buf
}

func (m *TestTypes) MarshalV2(buf []byte) []byte {
	buf = m.MarshalV1(buf)

	return buf
}

func (m *TestTypes) UnmarshalV1(buf []byte) []byte {

	buf, m.U64 = unmarshalUint64(buf)

	buf, m.S64 = unmarshalInt64(buf)

	buf, m.D = unmarshalFloat64(buf)

	for i, _ := range m.U64Array {
		buf, m.U64Array[i] = unmarshalUint64(buf)
	}

	for i, _ := range m.S64Array {
		buf, m.S64Array[i] = unmarshalInt64(buf)
	}

	for i, _ := range m.DArray {
		buf, m.DArray[i] = unmarshalFloat64(buf)
	}

	buf, m.U32 = unmarshalUint32(buf)

	buf, m.S32 = unmarshalInt32(buf)

	buf, m.F = unmarshalFloat32(buf)

	for i, _ := range m.U32Array {
		buf, m.U32Array[i] = unmarshalUint32(buf)
	}

	for i, _ := range m.S32Array {
		buf, m.S32Array[i] = unmarshalInt32(buf)
	}

	for i, _ := range m.FArray {
		buf, m.FArray[i] = unmarshalFloat32(buf)
	}

	buf, m.U16 = unmarshalUint16(buf)

	buf, m.S16 = unmarshalInt16(buf)

	for i, _ := range m.U16Array {
		buf, m.U16Array[i] = unmarshalUint16(buf)
	}

	for i, _ := range m.S16Array {
		buf, m.S16Array[i] = unmarshalInt16(buf)
	}

	buf, m.C = unmarshalByte(buf)

	for i, _ := range m.S {
		buf, m.S[i] = unmarshalByte(buf)
	}

	buf, m.U8 = unmarshalByte(buf)

	buf, m.S8 = unmarshalInt8(buf)

	for i, _ := range m.U8Array {
		buf, m.U8Array[i] = unmarshalByte(buf)
	}

	for i, _ := range m.S8Array {
		buf, m.S8Array[i] = unmarshalInt8(buf)
	}

	return buf
}

func (m *TestTypes) UnmarshalV2(buf []byte) []byte {
	buf = m.UnmarshalV1(buf)

	return buf
}

// These will be inlined.
func marshalByte(b []byte, v byte) []byte     { return append(b, v) }
func marshalInt8(b []byte, v int8) []byte     { return append(b, byte(v)) }
func marshalInt16(b []byte, v int16) []byte   { return append(b, byte(v), byte(v>>8)) }
func marshalUint16(b []byte, v uint16) []byte { return append(b, byte(v), byte(v>>8)) }
func marshalInt32(b []byte, v int32) []byte {
	return append(b, byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
}
func marshalUint32(b []byte, v uint32) []byte {
	return append(b, byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
}
func marshalInt64(b []byte, v int64) []byte {
	return append(b, byte(v), byte(v>>8), byte(v>>16), byte(v>>24), byte(v>>32), byte(v>>40), byte(v>>48), byte(v>>56))
}
func marshalUint64(b []byte, v uint64) []byte {
	return append(b, byte(v), byte(v>>8), byte(v>>16), byte(v>>24), byte(v>>32), byte(v>>40), byte(v>>48), byte(v>>56))
}
func marshalFloat32(b []byte, v float32) []byte { return marshalUint32(b, math.Float32bits(v)) }
func marshalFloat64(b []byte, v float64) []byte { return marshalUint64(b, math.Float64bits(v)) }

func unmarshalByte(b []byte) ([]byte, byte)     { return b[1:], b[0] }
func unmarshalInt8(b []byte) ([]byte, int8)     { return b[1:], int8(b[0]) }
func unmarshalInt16(b []byte) ([]byte, int16)   { return b[2:], int16(b[0]) | int16(b[1])<<8 }
func unmarshalUint16(b []byte) ([]byte, uint16) { return b[2:], uint16(b[0]) | uint16(b[1])<<8 }
func unmarshalInt32(b []byte) ([]byte, int32) {
	return b[4:], int32(b[0]) | int32(b[1])<<8 | int32(b[2])<<16 | int32(b[3])<<24
}
func unmarshalUint32(b []byte) ([]byte, uint32) {
	return b[4:], uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}
func unmarshalInt64(b []byte) ([]byte, int64) {
	return b[8:], int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 | int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
}
func unmarshalUint64(b []byte) ([]byte, uint64) {
	return b[8:], uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}
func unmarshalFloat32(b []byte) ([]byte, float32) {
	b, v := unmarshalUint32(b)
	return b, math.Float32frombits(v)
}
func unmarshalFloat64(b []byte) ([]byte, float64) {
	b, v := unmarshalUint64(b)
	return b, math.Float64frombits(v)
}
