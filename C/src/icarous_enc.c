// Generated enums and structures for Mavlink dialect icarous #0 version 0
// Generated by gomavlink, DO NOT EDIT.

#include <assert.h>
#include <string.h>

#include "icarous.h"

// defined in icarous_crc.h
uint8_t icarous_crcextra(enum ICAROUS_MESSAGE mid);

inline static void append_byte(uint8_t **p, uint8_t v) { *(*p)++ = v; }
inline static void append_int8(uint8_t **p, int8_t v) { *(*p)++ = v; }
inline static void append_int16(uint8_t **p, int16_t v) {
  *(*p)++ = v;
  *(*p)++ = v >> 8;
}
inline static void append_uint16(uint8_t **p, uint16_t v) {
  *(*p)++ = v;
  *(*p)++ = v >> 8;
}
inline static void append_int32(uint8_t **p, int32_t v) {
  *(*p)++ = v;
  *(*p)++ = v >> 8;
  *(*p)++ = v >> 16;
  *(*p)++ = v >> 24;
}
inline static void append_uint32(uint8_t **p, uint32_t v) {
  *(*p)++ = v;
  *(*p)++ = v >> 8;
  *(*p)++ = v >> 16;
  *(*p)++ = v >> 24;
}
inline static void append_int64(uint8_t **p, int64_t v) {
  *(*p)++ = v;
  *(*p)++ = v >> 8;
  *(*p)++ = v >> 16;
  *(*p)++ = v >> 24;
  *(*p)++ = v >> 32;
  *(*p)++ = v >> 40;
  *(*p)++ = v >> 48;
  *(*p)++ = v >> 56;
}
inline static void append_uint64(uint8_t **p, uint64_t v) {
  *(*p)++ = v;
  *(*p)++ = v >> 8;
  *(*p)++ = v >> 16;
  *(*p)++ = v >> 24;
  *(*p)++ = v >> 32;
  *(*p)++ = v >> 40;
  *(*p)++ = v >> 48;
  *(*p)++ = v >> 56;
}

inline static void append_float32(uint8_t **p, float v) {
  append_uint32(p, *(uint32_t *)&v);
}
inline static void append_float64(uint8_t **p, double v) {
  append_uint64(p, *(uint64_t *)&v);
}

inline static uint16_t x25(uint16_t crc, uint8_t v) {
  v ^= crc;
  v ^= v << 4;
  return (crc >> 8) ^ (v << 8) ^ (v << 3) ^
         (v >> 4); // (uint8_t << int) is promoted to int (!!)
}

static void append_payload_base(uint8_t **p, struct icarous_message *msg) {
  switch (msg->msg_id) {

  case ICAROUS_MESSAGE_ICAROUS_HEARTBEAT:
    append_byte(p, msg->icarous_heartbeat.status);
    break;

  case ICAROUS_MESSAGE_ICAROUS_KINEMATIC_BANDS:
    append_float32(p, msg->icarous_kinematic_bands.min1);
    append_float32(p, msg->icarous_kinematic_bands.max1);
    append_float32(p, msg->icarous_kinematic_bands.min2);
    append_float32(p, msg->icarous_kinematic_bands.max2);
    append_float32(p, msg->icarous_kinematic_bands.min3);
    append_float32(p, msg->icarous_kinematic_bands.max3);
    append_float32(p, msg->icarous_kinematic_bands.min4);
    append_float32(p, msg->icarous_kinematic_bands.max4);
    append_float32(p, msg->icarous_kinematic_bands.min5);
    append_float32(p, msg->icarous_kinematic_bands.max5);
    append_int8(p, msg->icarous_kinematic_bands.numBands);
    append_byte(p, msg->icarous_kinematic_bands.type1);
    append_byte(p, msg->icarous_kinematic_bands.type2);
    append_byte(p, msg->icarous_kinematic_bands.type3);
    append_byte(p, msg->icarous_kinematic_bands.type4);
    append_byte(p, msg->icarous_kinematic_bands.type5);
    break;
  };
}

static void append_payload_ext(uint8_t **p, struct icarous_message *msg) {
  switch (msg->msg_id) {

  default:; // nix
  };
}

size_t icarous_message_serialize_v1(struct icarous_message *msg, uint8_t *buf,
                                    size_t size) {

  if (msg->msg_id > 255)
    return 0;

  // it's faster to alloc and copy at end than to test buf[size] bound at every
  // byte
  uint8_t tmp[264];
  uint8_t *p = tmp;

  *p++ = 0xFE; // V1 STX
  *p++ = 0;    // placeholder for len
  *p++ = msg->seq_nr;
  *p++ = msg->sys_id;
  *p++ = msg->comp_id;
  *p++ = msg->msg_id;

  append_payload_base(&p, msg);

  tmp[1] = p - tmp - 6; // set len

  uint16_t crc = 0xffff;
  for (uint8_t *pp = tmp; pp < p; ++pp) {
    crc = x25(crc, *pp);
  }
  crc = x25(crc, icarous_crcextra(msg->msg_id));
  *p++ = crc; // overwrite crc_extra with the lower byte
  *p++ = crc >> 8;

  size_t len = p - tmp;
  assert(len < sizeof tmp);
  if (len < size) {
    memmove(buf, tmp, len);
    ++msg->seq_nr;
    msg->len = len;
    msg->crc =
        crc; // may be useful for testing idempotency of serialize/deserialize
  }
  return len;
}

size_t icarous_message_serialize_v2(struct icarous_message *msg, uint8_t *buf,
                                    size_t size) {
  // it's faster to alloc and copy at end than to test buf[size] bound at every
  // byte
  uint8_t tmp[280];
  uint8_t *p = tmp;

  *p++ = 0xFD;                     // V2 STX
  *p++ = 0;                        // placeholder for len
  *p++ = (msg->signature) ? 1 : 0; // incompat flag (0x1 means signed)
  *p++ = 0;                        // compat flags
  *p++ = msg->seq_nr;
  *p++ = msg->sys_id;
  *p++ = msg->comp_id;
  *p++ = msg->msg_id;
  *p++ = msg->msg_id >> 8;
  *p++ = msg->msg_id >> 16; // 10 bytes

  append_payload_base(&p, msg);
  append_payload_ext(&p, msg);

  // trim trailing zeroes, a v2 only feature
  while (*--p == 0)
    if (p <= tmp + 10)
      break;
  // now p points to the last non-zero element of tmp;
  ++p;

  tmp[1] = p - tmp - 10; // fill in len

  uint16_t crc = 0xffff;
  for (uint8_t *pp = tmp; pp < p; ++pp) {
    crc = x25(crc, *pp);
  }
  crc = x25(crc, icarous_crcextra(msg->msg_id));

  *p++ = crc;
  *p++ = crc >> 8;

  // zero is never a valid signature. if your hash turns out to zero you've
  // mined a bitcoin.
  if (msg->signature) {
    *p++ = msg->link_id;
    append_uint16(&p, msg->timestamp);
    append_uint32(&p, msg->timestamp >> 16);
    append_uint16(&p, msg->signature);
    append_uint32(&p, msg->signature >> 16);
  }

  size_t len = p - tmp;
  assert(len < sizeof tmp);
  if (len < size) {
    memmove(buf, tmp, len);
    ++msg->seq_nr;
    msg->len = len;
    msg->crc =
        crc; // may be useful for testing idempotency of serialize/deserialize
  }
  return len;
}
