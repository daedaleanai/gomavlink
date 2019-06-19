/*
    test driver for mavgen code: read ardupilot logs (which have a timestamp)
    from stdin and printf them on stdout.
    gcc -O3 dump.c lib.a -o dump
*/

#include <errno.h>
#include <fcntl.h>
#include <stdarg.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/uio.h>
#include <time.h>
#include <unistd.h>

#include "include/ardupilotmega.h"

static void crash(const char *fmt, ...) {
    va_list ap;
    char buf[1000];
    va_start(ap, fmt);
    vsnprintf(buf, sizeof(buf), fmt, ap);
    fprintf(stderr, "%s%s%s\n", buf, (errno) ? ": " : "", (errno) ? strerror(errno) : "");
    exit(1);
    va_end(ap);
    return;
}

int main(int argc, char *argv[]) {
    uint8_t buf[4096];
    size_t end = 0;

    int fd = fileno(stdin);
    if (argc == 2) {
        fd = open(argv[1], O_RDONLY);
    }

    for (;;) {
        ssize_t r = read(fd, buf + end, (sizeof buf) - end);
        if (r < 0)
            crash("read");
        end += r;

        if (end < 8)
            break;

        uint64_t ts = (uint64_t)buf[0] << 56;
        ts |= (uint64_t)buf[1] << 48;
        ts |= (uint64_t)buf[2] << 40;
        ts |= (uint64_t)buf[3] << 32;
        ts |= (uint64_t)buf[4] << 24;
        ts |= (uint64_t)buf[5] << 16;
        ts |= (uint64_t)buf[6] << 8;
        ts |= (uint64_t)buf[7];
        time_t tt = ts / 1000000;

        struct ardupilotmega_message msg;

        size_t s = ardupilotmega_message_deserialize(buf + 8, end - 8, &msg);
        if (s == 0)
            crash("incomplete message");
        end -= 8 + s;
        memmove(buf, buf + 8 + s, end);

        // TODO better format timestamp
        fprintf(stdout, "%s.%0llu:", ctime(&tt), ts % 1000000);
        char out[65536];
        ardupilotmega_message_snprintf(out, sizeof out, &msg);
        puts(out);
    }

    return 0;
}
