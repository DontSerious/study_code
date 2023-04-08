#include <stdio.h>

// 看当前机器是大端小端
int is_little_endian () {
    short s = 1;
    unsigned char * p = (unsigned char *) &s;
    return p[0] > p[1] ? 1 : 0;
}