// Generated enums and structures for Mavlink dialect icarous #0 version 0
// Generated by gomavlink, DO NOT EDIT.
#include <stdarg.h>
#include <stdio.h>

#include "icarous.h"

static inline size_t psnprintf(char **pbuf, size_t *psize, const char *fmt,
                               ...) {
  va_list ap;
  va_start(ap, fmt);
  size_t s = vsnprintf(*pbuf, *psize, fmt, ap);
  if (s < *psize) {
    *pbuf += s;
    *psize -= s;
  } else {
    *psize = 0;
  }
  return s;
}

static const char *
ICAROUS_TRACK_BAND_TYPES_str(enum ICAROUS_TRACK_BAND_TYPES val) {
  switch (val) {
  case ICAROUS_TRACK_BAND_TYPE_NONE:
    return "ICAROUS_TRACK_BAND_TYPE_NONE";
  case ICAROUS_TRACK_BAND_TYPE_NEAR:
    return "ICAROUS_TRACK_BAND_TYPE_NEAR";
  case ICAROUS_TRACK_BAND_TYPE_RECOVERY:
    return "ICAROUS_TRACK_BAND_TYPE_RECOVERY";
  }
  return NULL;
};

static size_t
psnprint_ICAROUS_TRACK_BAND_TYPES(char **p, size_t *sz,
                                  enum ICAROUS_TRACK_BAND_TYPES val) {
  const char *s = ICAROUS_TRACK_BAND_TYPES_str(val);
  if (s) {
    return psnprintf(p, sz, "%s", s);
  }
  return psnprintf(p, sz, "ICAROUS_TRACK_BAND_TYPES<%d>", (int)val);
}

static const char *ICAROUS_FMS_STATE_str(enum ICAROUS_FMS_STATE val) {
  switch (val) {
  case ICAROUS_FMS_STATE_IDLE:
    return "ICAROUS_FMS_STATE_IDLE";
  case ICAROUS_FMS_STATE_TAKEOFF:
    return "ICAROUS_FMS_STATE_TAKEOFF";
  case ICAROUS_FMS_STATE_CLIMB:
    return "ICAROUS_FMS_STATE_CLIMB";
  case ICAROUS_FMS_STATE_CRUISE:
    return "ICAROUS_FMS_STATE_CRUISE";
  case ICAROUS_FMS_STATE_APPROACH:
    return "ICAROUS_FMS_STATE_APPROACH";
  case ICAROUS_FMS_STATE_LAND:
    return "ICAROUS_FMS_STATE_LAND";
  }
  return NULL;
};

static size_t psnprint_ICAROUS_FMS_STATE(char **p, size_t *sz,
                                         enum ICAROUS_FMS_STATE val) {
  const char *s = ICAROUS_FMS_STATE_str(val);
  if (s) {
    return psnprintf(p, sz, "%s", s);
  }
  return psnprintf(p, sz, "ICAROUS_FMS_STATE<%d>", (int)val);
}

size_t icarous_message_snprintf(char *buf, size_t size,
                                const struct icarous_message *msg) {
  size_t r = 0;
  switch (msg->msg_id) {

  case ICAROUS_MESSAGE_ICAROUS_HEARTBEAT:

    r += psnprintf(&buf, &size, "ICAROUS_HEARTBEAT<");

    r += psnprintf(&buf, &size, " status:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " >");
    break;

  case ICAROUS_MESSAGE_ICAROUS_KINEMATIC_BANDS:

    r += psnprintf(&buf, &size, "ICAROUS_KINEMATIC_BANDS<");

    r += psnprintf(&buf, &size, " min1:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " max1:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " min2:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " max2:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " min3:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " max3:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " min4:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " max4:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " min5:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " max5:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " numBands:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " type1:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " type2:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " type3:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " type4:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " type5:");
    // TODO values: handle strings, arrays and enums
    r += psnprintf(&buf, &size, " >");
    break;
  };

  return r;
}
