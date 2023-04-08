#include <stdio.h>
#include <assert.h>
#include <limits.h>

/* Addition that saturates to TMin or TMax */
int saturating_add(int x, int y) {
    int sum = x + y;
    int sig_mask = INT_MIN;

    /*
        当x>0 y>0时 如果sum小于0则正溢出
        当x<0 y<0时 如果sum大于0则负溢出
    */
    int pos_over = !(x & sig_mask) && !(y & sig_mask) && (sum & sig_mask);
    int neg_over = (x & sig_mask) && (y & sig_mask) && !(sum &sig_mask);

    (pos_over && (sum = INT_MAX) || neg_over && (sum = INT_MIN));

    return sum;
}

int main(int argc, char* argv[]) {
  assert(INT_MAX == saturating_add(INT_MAX, 0x1234));
  assert(INT_MIN == saturating_add(INT_MIN, -0x1234));
  assert(0x11 + 0x22 == saturating_add(0x11, 0x22));
  return 0;
}