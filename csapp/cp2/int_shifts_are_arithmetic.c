#include <stdio.h>
#include <assert.h>

// 测试当前机器是否为算术右移
int int_shifts_are_arithmetic() {
    int num = -1;
    return !(num ^ (num >> 1));
}

int main(int argc, char* argv[]) {
    assert(int_shifts_are_arithmetic());
    return 0;
}
