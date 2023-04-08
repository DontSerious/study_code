#include <stdio.h>
#include <assert.h>
#include <limits.h>

/* Determine whether arguments can be subtracted without overflow */
int tSub_ok(int x, int y) {
    int res = 1;

    (y == INT_MIN) && (res = 0);
    // if (y == INT_MIN) res = 0;               因为int_maX的模一定小于int_min的模

    // 当x>0 y<0 以及 x < 0 y > 0 才有可能导致溢出
    int sub = x - y;
    int pos_over = x > 0 && y < 0 && sub < 0;
    int neg_over = x < 0 && y > 0 && sub > 0;

    res = res && !(pos_over || neg_over);

    return res;
}

int main(int argc, char* argv[]) {
  assert(!tSub_ok(0x00, INT_MIN));
  assert(tSub_ok(0x00, 0x00));
  return 0;
}