// Generated enums and structures for Mavlink dialect minimal #0 version 3
package minimal

// Generated by gomavlink, DO NOT EDIT.

import (
	"math"

	mavlink "github.com/daedaleanai/gomavlink"
)

//go:generate stringer -output strings.go -type=MavAutopilot,MavType,MavModeFlag,MavModeFlagDecodePosition,MavState,MavComponent

/* Micro air vehicle / autopilot classes. This identifies the individual model. */
type MavAutopilot uint32

const (
	/* Generic autopilot, full support for everything */
	MAV_AUTOPILOT_GENERIC MavAutopilot = 0

	/* Reserved for future use. */
	MAV_AUTOPILOT_RESERVED MavAutopilot = 1

	/* SLUGS autopilot, http://slugsuav.soe.ucsc.edu */
	MAV_AUTOPILOT_SLUGS MavAutopilot = 2

	/* ArduPilot - Plane/Copter/Rover/Sub/Tracker, http://ardupilot.org */
	MAV_AUTOPILOT_ARDUPILOTMEGA MavAutopilot = 3

	/* OpenPilot, http://openpilot.org */
	MAV_AUTOPILOT_OPENPILOT MavAutopilot = 4

	/* Generic autopilot only supporting simple waypoints */
	MAV_AUTOPILOT_GENERIC_WAYPOINTS_ONLY MavAutopilot = 5

	/* Generic autopilot supporting waypoints and other simple navigation commands */
	MAV_AUTOPILOT_GENERIC_WAYPOINTS_AND_SIMPLE_NAVIGATION_ONLY MavAutopilot = 6

	/* Generic autopilot supporting the full mission command set */
	MAV_AUTOPILOT_GENERIC_MISSION_FULL MavAutopilot = 7

	/* No valid autopilot, e.g. a GCS or other MAVLink component */
	MAV_AUTOPILOT_INVALID MavAutopilot = 8

	/* PPZ UAV - http://nongnu.org/paparazzi */
	MAV_AUTOPILOT_PPZ MavAutopilot = 9

	/* UAV Dev Board */
	MAV_AUTOPILOT_UDB MavAutopilot = 10

	/* FlexiPilot */
	MAV_AUTOPILOT_FP MavAutopilot = 11

	/* PX4 Autopilot - http://px4.io/ */
	MAV_AUTOPILOT_PX4 MavAutopilot = 12

	/* SMACCMPilot - http://smaccmpilot.org */
	MAV_AUTOPILOT_SMACCMPILOT MavAutopilot = 13

	/* AutoQuad -- http://autoquad.org */
	MAV_AUTOPILOT_AUTOQUAD MavAutopilot = 14

	/* Armazila -- http://armazila.com */
	MAV_AUTOPILOT_ARMAZILA MavAutopilot = 15

	/* Aerob -- http://aerob.ru */
	MAV_AUTOPILOT_AEROB MavAutopilot = 16

	/* ASLUAV autopilot -- http://www.asl.ethz.ch */
	MAV_AUTOPILOT_ASLUAV MavAutopilot = 17

	/* SmartAP Autopilot - http://sky-drones.com */
	MAV_AUTOPILOT_SMARTAP MavAutopilot = 18

	/* AirRails - http://uaventure.com */
	MAV_AUTOPILOT_AIRRAILS MavAutopilot = 19
)

/* MAVLINK system type. All components in a system should report this type in their HEARTBEAT. */
type MavType uint32

const (
	/* Generic micro air vehicle. */
	MAV_TYPE_GENERIC MavType = 0

	/* Fixed wing aircraft. */
	MAV_TYPE_FIXED_WING MavType = 1

	/* Quadrotor */
	MAV_TYPE_QUADROTOR MavType = 2

	/* Coaxial helicopter */
	MAV_TYPE_COAXIAL MavType = 3

	/* Normal helicopter with tail rotor. */
	MAV_TYPE_HELICOPTER MavType = 4

	/* Ground installation */
	MAV_TYPE_ANTENNA_TRACKER MavType = 5

	/* Operator control unit / ground control station */
	MAV_TYPE_GCS MavType = 6

	/* Airship, controlled */
	MAV_TYPE_AIRSHIP MavType = 7

	/* Free balloon, uncontrolled */
	MAV_TYPE_FREE_BALLOON MavType = 8

	/* Rocket */
	MAV_TYPE_ROCKET MavType = 9

	/* Ground rover */
	MAV_TYPE_GROUND_ROVER MavType = 10

	/* Surface vessel, boat, ship */
	MAV_TYPE_SURFACE_BOAT MavType = 11

	/* Submarine */
	MAV_TYPE_SUBMARINE MavType = 12

	/* Hexarotor */
	MAV_TYPE_HEXAROTOR MavType = 13

	/* Octorotor */
	MAV_TYPE_OCTOROTOR MavType = 14

	/* Tricopter */
	MAV_TYPE_TRICOPTER MavType = 15

	/* Flapping wing */
	MAV_TYPE_FLAPPING_WING MavType = 16

	/* Kite */
	MAV_TYPE_KITE MavType = 17

	/* Onboard companion controller */
	MAV_TYPE_ONBOARD_CONTROLLER MavType = 18

	/* Two-rotor VTOL using control surfaces in vertical operation in addition. Tailsitter. */
	MAV_TYPE_VTOL_DUOROTOR MavType = 19

	/* Quad-rotor VTOL using a V-shaped quad config in vertical operation. Tailsitter. */
	MAV_TYPE_VTOL_QUADROTOR MavType = 20

	/* Tiltrotor VTOL */
	MAV_TYPE_VTOL_TILTROTOR MavType = 21

	/* VTOL reserved 2 */
	MAV_TYPE_VTOL_RESERVED2 MavType = 22

	/* VTOL reserved 3 */
	MAV_TYPE_VTOL_RESERVED3 MavType = 23

	/* VTOL reserved 4 */
	MAV_TYPE_VTOL_RESERVED4 MavType = 24

	/* VTOL reserved 5 */
	MAV_TYPE_VTOL_RESERVED5 MavType = 25

	/* Gimbal (standalone) */
	MAV_TYPE_GIMBAL MavType = 26

	/* ADSB system (standalone) */
	MAV_TYPE_ADSB MavType = 27

	/* Steerable, nonrigid airfoil */
	MAV_TYPE_PARAFOIL MavType = 28

	/* Dodecarotor */
	MAV_TYPE_DODECAROTOR MavType = 29

	/* Camera (standalone) */
	MAV_TYPE_CAMERA MavType = 30

	/* Charging station */
	MAV_TYPE_CHARGING_STATION MavType = 31

	/* FLARM collision avoidance system (standalone) */
	MAV_TYPE_FLARM MavType = 32
)

/* These flags encode the MAV mode. */
type MavModeFlag uint32

const (
	/* 0b10000000 MAV safety set to armed. Motors are enabled / running / can start. Ready to fly. Additional note: this flag is to be ignore when sent in the command MAV_CMD_DO_SET_MODE and MAV_CMD_COMPONENT_ARM_DISARM shall be used instead. The flag can still be used to report the armed state. */
	MAV_MODE_FLAG_SAFETY_ARMED MavModeFlag = 128

	/* 0b01000000 remote control input is enabled. */
	MAV_MODE_FLAG_MANUAL_INPUT_ENABLED MavModeFlag = 64

	/* 0b00100000 hardware in the loop simulation. All motors / actuators are blocked, but internal software is full operational. */
	MAV_MODE_FLAG_HIL_ENABLED MavModeFlag = 32

	/* 0b00010000 system stabilizes electronically its attitude (and optionally position). It needs however further control inputs to move around. */
	MAV_MODE_FLAG_STABILIZE_ENABLED MavModeFlag = 16

	/* 0b00001000 guided mode enabled, system flies waypoints / mission items. */
	MAV_MODE_FLAG_GUIDED_ENABLED MavModeFlag = 8

	/* 0b00000100 autonomous mode enabled, system finds its own goal positions. Guided flag can be set or not, depends on the actual implementation. */
	MAV_MODE_FLAG_AUTO_ENABLED MavModeFlag = 4

	/* 0b00000010 system has a test mode enabled. This flag is intended for temporary system tests and should not be used for stable implementations. */
	MAV_MODE_FLAG_TEST_ENABLED MavModeFlag = 2

	/* 0b00000001 Reserved for future use. */
	MAV_MODE_FLAG_CUSTOM_MODE_ENABLED MavModeFlag = 1
)

/* These values encode the bit positions of the decode position. These values can be used to read the value of a flag bit by combining the base_mode variable with AND with the flag position value. The result will be either 0 or 1, depending on if the flag is set or not. */
type MavModeFlagDecodePosition uint32

const (
	/* First bit:  10000000 */
	MAV_MODE_FLAG_DECODE_POSITION_SAFETY MavModeFlagDecodePosition = 128

	/* Second bit: 01000000 */
	MAV_MODE_FLAG_DECODE_POSITION_MANUAL MavModeFlagDecodePosition = 64

	/* Third bit:  00100000 */
	MAV_MODE_FLAG_DECODE_POSITION_HIL MavModeFlagDecodePosition = 32

	/* Fourth bit: 00010000 */
	MAV_MODE_FLAG_DECODE_POSITION_STABILIZE MavModeFlagDecodePosition = 16

	/* Fifth bit:  00001000 */
	MAV_MODE_FLAG_DECODE_POSITION_GUIDED MavModeFlagDecodePosition = 8

	/* Sixt bit:   00000100 */
	MAV_MODE_FLAG_DECODE_POSITION_AUTO MavModeFlagDecodePosition = 4

	/* Seventh bit: 00000010 */
	MAV_MODE_FLAG_DECODE_POSITION_TEST MavModeFlagDecodePosition = 2

	/* Eighth bit: 00000001 */
	MAV_MODE_FLAG_DECODE_POSITION_CUSTOM_MODE MavModeFlagDecodePosition = 1
)

type MavState uint32

const (
	/* Uninitialized system, state is unknown. */
	MAV_STATE_UNINIT MavState = 0

	/* System is booting up. */
	MAV_STATE_BOOT MavState = 2

	/* System is calibrating and not flight-ready. */
	MAV_STATE_CALIBRATING MavState = 3

	/* System is grounded and on standby. It can be launched any time. */
	MAV_STATE_STANDBY MavState = 4

	/* System is active and might be already airborne. Motors are engaged. */
	MAV_STATE_ACTIVE MavState = 5

	/* System is in a non-normal flight mode. It can however still navigate. */
	MAV_STATE_CRITICAL MavState = 6

	/* System is in a non-normal flight mode. It lost control over parts or over the whole airframe. It is in mayday and going down. */
	MAV_STATE_EMERGENCY MavState = 7

	/* System just initialized its power-down sequence, will shut down now. */
	MAV_STATE_POWEROFF MavState = 8

	/* System is terminating itself. */
	MAV_STATE_FLIGHT_TERMINATION MavState = 9
)

/* Component ids (values) for the different types and instances of onboard hardware/software that might make up a MAVLink system (autopilot, cameras, servos, GPS systems, avoidance systems etc.).        Components must use the appropriate ID in their source address when sending messages. Components can also use IDs to determine if they are the intended recipient of an incoming message. The MAV_COMP_ID_ALL value is used to indicate messages that must be processed by all components.       When creating new entries, components that can have multiple instances (e.g. cameras, servos etc.) should be allocated sequential values. An appropriate number of values should be left free after these components to allow the number of instances to be expanded. */
type MavComponent uint32

const (
	/* Used to broadcast messages to all components of the receiving system. Components should attempt to process messages with this component ID and forward to components on any other interfaces. */
	MAV_COMP_ID_ALL MavComponent = 0

	/* System flight controller component ("autopilot"). Only one autopilot is expected in a particular system. */
	MAV_COMP_ID_AUTOPILOT1 MavComponent = 1

	/* Camera #1. */
	MAV_COMP_ID_CAMERA MavComponent = 100

	/* Camera #2. */
	MAV_COMP_ID_CAMERA2 MavComponent = 101

	/* Camera #3. */
	MAV_COMP_ID_CAMERA3 MavComponent = 102

	/* Camera #4. */
	MAV_COMP_ID_CAMERA4 MavComponent = 103

	/* Camera #5. */
	MAV_COMP_ID_CAMERA5 MavComponent = 104

	/* Camera #6. */
	MAV_COMP_ID_CAMERA6 MavComponent = 105

	/* Servo #1. */
	MAV_COMP_ID_SERVO1 MavComponent = 140

	/* Servo #2. */
	MAV_COMP_ID_SERVO2 MavComponent = 141

	/* Servo #3. */
	MAV_COMP_ID_SERVO3 MavComponent = 142

	/* Servo #4. */
	MAV_COMP_ID_SERVO4 MavComponent = 143

	/* Servo #5. */
	MAV_COMP_ID_SERVO5 MavComponent = 144

	/* Servo #6. */
	MAV_COMP_ID_SERVO6 MavComponent = 145

	/* Servo #7. */
	MAV_COMP_ID_SERVO7 MavComponent = 146

	/* Servo #8. */
	MAV_COMP_ID_SERVO8 MavComponent = 147

	/* Servo #9. */
	MAV_COMP_ID_SERVO9 MavComponent = 148

	/* Servo #10. */
	MAV_COMP_ID_SERVO10 MavComponent = 149

	/* Servo #11. */
	MAV_COMP_ID_SERVO11 MavComponent = 150

	/* Servo #12. */
	MAV_COMP_ID_SERVO12 MavComponent = 151

	/* Servo #13. */
	MAV_COMP_ID_SERVO13 MavComponent = 152

	/* Servo #14. */
	MAV_COMP_ID_SERVO14 MavComponent = 153

	/* Gimbal component. */
	MAV_COMP_ID_GIMBAL MavComponent = 154

	/* Logging component. */
	MAV_COMP_ID_LOG MavComponent = 155

	/* Automatic Dependent Surveillance-Broadcast (ADS-B) component. */
	MAV_COMP_ID_ADSB MavComponent = 156

	/* On Screen Display (OSD) devices for video links. */
	MAV_COMP_ID_OSD MavComponent = 157

	/* Generic autopilot peripheral component ID. Meant for devices that do not implement the parameter microservice. */
	MAV_COMP_ID_PERIPHERAL MavComponent = 158

	/* Gimbal ID for QX1. */
	MAV_COMP_ID_QX1_GIMBAL MavComponent = 159

	/* FLARM collision alert component. */
	MAV_COMP_ID_FLARM MavComponent = 160

	/* Component that supports the Mission microservice. */
	MAV_COMP_ID_MISSIONPLANNER MavComponent = 190

	/* Component that finds an optimal path between points based on a certain constraint (e.g. minimum snap, shortest path, cost, etc.). */
	MAV_COMP_ID_PATHPLANNER MavComponent = 195

	/* Component that plans a collision free path between two points. */
	MAV_COMP_ID_OBSTACLE_AVOIDANCE MavComponent = 196

	/* Component that provides position estimates using VIO techniques. */
	MAV_COMP_ID_VISUAL_INERTIAL_ODOMETRY MavComponent = 197

	/* Inertial Measurement Unit (IMU) #1. */
	MAV_COMP_ID_IMU MavComponent = 200

	/* Inertial Measurement Unit (IMU) #2. */
	MAV_COMP_ID_IMU_2 MavComponent = 201

	/* Inertial Measurement Unit (IMU) #3. */
	MAV_COMP_ID_IMU_3 MavComponent = 202

	/* GPS #1. */
	MAV_COMP_ID_GPS MavComponent = 220

	/* GPS #2. */
	MAV_COMP_ID_GPS2 MavComponent = 221

	/* Component to bridge MAVLink to UDP (i.e. from a UART). */
	MAV_COMP_ID_UDP_BRIDGE MavComponent = 240

	/* Component to bridge to UART (i.e. from UDP). */
	MAV_COMP_ID_UART_BRIDGE MavComponent = 241

	/* Component for handling system messages (e.g. to ARM, takeoff, etc.). */
	MAV_COMP_ID_SYSTEM_CONTROL MavComponent = 250
)

// The Dialect factory function constructs the proper empty message given the message ID.
func Dialect(mid int) mavlink.Message {
	switch mid {
	case 0:
		return &Heartbeat{}
	}
	return nil
}

/* The heartbeat message shows that a system or component is present and responding. The type and autopilot fields (along with the message component id), allow the receiving system to treat further messages from this system appropriately (e.g. by laying out the user interface based on the autopilot). */
type Heartbeat struct {
	/* A bitfield for use for autopilot-specific flags */
	CustomMode uint32

	/* Type of the system (quadrotor, helicopter, etc.). Components use the same type as their associated system. */
	Type MavType // byte

	/* Autopilot type / class. */
	Autopilot MavAutopilot // byte

	/* System mode bitmap. */
	BaseMode MavModeFlag // byte

	/* System status flag. */
	SystemStatus MavState // byte

	/* MAVLink version, not writable by user, gets added by protocol because of magic data type: uint8_t_mavlink_version */
	MavlinkVersion byte
}

func (m *Heartbeat) ID() int        { return 0 }
func (m *Heartbeat) CRCExtra() byte { return 50 }

func (m *Heartbeat) MarshalV1(buf []byte) []byte {
	buf = marshalUint32(buf, (m.CustomMode))
	buf = marshalByte(buf, byte(m.Type))
	buf = marshalByte(buf, byte(m.Autopilot))
	buf = marshalByte(buf, byte(m.BaseMode))
	buf = marshalByte(buf, byte(m.SystemStatus))
	buf = marshalByte(buf, (m.MavlinkVersion))

	return buf
}

func (m *Heartbeat) MarshalV2(buf []byte) []byte {
	buf = m.MarshalV1(buf)

	return buf
}

func (m *Heartbeat) UnmarshalV1(buf []byte) []byte {

	buf, m.CustomMode = unmarshalUint32(buf)

	{
		var v byte
		buf, v = unmarshalByte(buf)
		m.Type = MavType(v)
	}

	{
		var v byte
		buf, v = unmarshalByte(buf)
		m.Autopilot = MavAutopilot(v)
	}

	{
		var v byte
		buf, v = unmarshalByte(buf)
		m.BaseMode = MavModeFlag(v)
	}

	{
		var v byte
		buf, v = unmarshalByte(buf)
		m.SystemStatus = MavState(v)
	}

	buf, m.MavlinkVersion = unmarshalByte(buf)

	return buf
}

func (m *Heartbeat) UnmarshalV2(buf []byte) []byte {
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
