// Generated enums and structures for Mavlink dialect test #0 version 3
package test

import "math"

// Generated by gomavlink, DO NOT EDIT.

type Message interface {
	ID() int
	CRCExtra() uint16
	MarshalV1([]byte) []byte
	MarshalV2([]byte) []byte
}

func New(mid uint32) Message {
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

func (m *TestTypes) ID() int          { return 0 }
func (m *TestTypes) CRCExtra() uint16 { return 26368 }

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
