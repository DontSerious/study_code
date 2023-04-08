#include <stdio.h>
#include <assert.h>
#include <limits.h>

// 计算3/4x 向0舍入 使它不会溢出
int threeFourths(int x) {
    // 因为要求除四，需要向右移两位计算，会损失精度，所以需要分开计算
    int is_neg = x & INT_MIN;
    int f = x & ~0x3;
    int l = x & 0x3;
    
    // 高位先乘会导致溢出，如果先进行除法再进行乘法就不会溢出
    int fd4 = f >> 2;
    int fd4m3 = (fd4 << 1) + fd4;

    // 低位计算需要考虑bias
    // `( x < 0 ? (x + (1 << k) - 1) : x) >> k`
    int lm3 = (l << 1) + l;
    int bias = (1 << 2) - 1;
    (is_neg && (lm3 += bias));
    int lm3d4 = lm3 >> 2;

    return fd4m3 + lm3d4;
}

int main(int argc, char* argv[]) {
  assert(threeFourths(8) == 6);
  assert(threeFourths(9) == 6);
  assert(threeFourths(10) == 7);
  assert(threeFourths(11) == 8);
  assert(threeFourths(12) == 9);

  assert(threeFourths(-8) == -6);
  assert(threeFourths(-9) == -6);
  assert(threeFourths(-10) == -7);
  assert(threeFourths(-11) == -8);
  assert(threeFourths(-12) == -9);
  return 0;
}