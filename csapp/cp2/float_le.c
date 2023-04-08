#include <stdio.h>
#include <assert.h>

unsigned f2u(float x) {
  return *(unsigned*)&x;
}

// 判断 x 是否小于 y
int float_le(float x, float y) {
  unsigned ux = f2u(x);
  unsigned uy = f2u(y);

  unsigned sx = ux >> 31;
  unsigned sy = uy >> 31;

  return (ux << 1 == 0 && uy << 1 == 0) || /* 去掉符号位都为 0 时 */
    (sx && !sy) ||                         /* 限制 x 为负数时，y一定为正数 */
    (!sx && !sy && ux <= uy) ||
    (sx && sy && ux >= uy);
}

int main(int argc, char* argv[]) {
  assert(float_le(-0, +0));
  assert(float_le(+0, -0));
  assert(float_le(0, 3));
  assert(float_le(-4, -0));
  assert(float_le(-4, 4));
  return 0;
}


