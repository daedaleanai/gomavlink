// Generated enums and structures for Mavlink dialect icarous #0 version 0
package icarous

import "math"

// Generated by gomavlink, DO NOT EDIT.
//go:generate stringer -output strings.go -type=IcarousTrackBandTypes,IcarousFmsState

type IcarousTrackBandTypes uint32

const (
	ICAROUS_TRACK_BAND_TYPE_NONE IcarousTrackBandTypes = 0

	ICAROUS_TRACK_BAND_TYPE_NEAR IcarousTrackBandTypes = 1

	ICAROUS_TRACK_BAND_TYPE_RECOVERY IcarousTrackBandTypes = 2
)

type IcarousFmsState uint32

const (
	ICAROUS_FMS_STATE_IDLE IcarousFmsState = 0

	ICAROUS_FMS_STATE_TAKEOFF IcarousFmsState = 1

	ICAROUS_FMS_STATE_CLIMB IcarousFmsState = 2

	ICAROUS_FMS_STATE_CRUISE IcarousFmsState = 3

	ICAROUS_FMS_STATE_APPROACH IcarousFmsState = 4

	ICAROUS_FMS_STATE_LAND IcarousFmsState = 5
)

type Message interface {
	ID() int
	CRCExtra() uint16
	MarshalV1([]byte) []byte
	MarshalV2([]byte) []byte
}

func New(mid uint32) Message {
	switch mid {
	case 42000:
		return &IcarousHeartbeat{}
	case 42001:
		return &IcarousKinematicBands{}
	}
	return nil
}

/* ICAROUS heartbeat */
type IcarousHeartbeat struct {
	/* See the FMS_STATE enum. */
	Status IcarousFmsState // byte

}

func (m *IcarousHeartbeat) ID() int          { return 42000 }
func (m *IcarousHeartbeat) CRCExtra() uint16 { return 41282 }

func (m *IcarousHeartbeat) MarshalV1(buf []byte) []byte {
	buf = marshalByte(buf, byte(m.Status))

	return buf
}
func (m *IcarousHeartbeat) MarshalV2(buf []byte) []byte {
	buf = m.MarshalV1(buf)

	return buf
}

/* Kinematic multi bands (track) output from Daidalus */
type IcarousKinematicBands struct {
	/* min angle (degrees) */
	Min1 float32

	/* max angle (degrees) */
	Max1 float32

	/* min angle (degrees) */
	Min2 float32

	/* max angle (degrees) */
	Max2 float32

	/* min angle (degrees) */
	Min3 float32

	/* max angle (degrees) */
	Max3 float32

	/* min angle (degrees) */
	Min4 float32

	/* max angle (degrees) */
	Max4 float32

	/* min angle (degrees) */
	Min5 float32

	/* max angle (degrees) */
	Max5 float32

	/* Number of track bands */
	Numbands int8

	/* See the TRACK_BAND_TYPES enum. */
	Type1 IcarousTrackBandTypes // byte

	/* See the TRACK_BAND_TYPES enum. */
	Type2 IcarousTrackBandTypes // byte

	/* See the TRACK_BAND_TYPES enum. */
	Type3 IcarousTrackBandTypes // byte

	/* See the TRACK_BAND_TYPES enum. */
	Type4 IcarousTrackBandTypes // byte

	/* See the TRACK_BAND_TYPES enum. */
	Type5 IcarousTrackBandTypes // byte

}

func (m *IcarousKinematicBands) ID() int          { return 42001 }
func (m *IcarousKinematicBands) CRCExtra() uint16 { return 55351 }

func (m *IcarousKinematicBands) MarshalV1(buf []byte) []byte {
	buf = marshalFloat32(buf, (m.Min1))
	buf = marshalFloat32(buf, (m.Max1))
	buf = marshalFloat32(buf, (m.Min2))
	buf = marshalFloat32(buf, (m.Max2))
	buf = marshalFloat32(buf, (m.Min3))
	buf = marshalFloat32(buf, (m.Max3))
	buf = marshalFloat32(buf, (m.Min4))
	buf = marshalFloat32(buf, (m.Max4))
	buf = marshalFloat32(buf, (m.Min5))
	buf = marshalFloat32(buf, (m.Max5))
	buf = marshalInt8(buf, (m.Numbands))
	buf = marshalByte(buf, byte(m.Type1))
	buf = marshalByte(buf, byte(m.Type2))
	buf = marshalByte(buf, byte(m.Type3))
	buf = marshalByte(buf, byte(m.Type4))
	buf = marshalByte(buf, byte(m.Type5))

	return buf
}
func (m *IcarousKinematicBands) MarshalV2(buf []byte) []byte {
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
