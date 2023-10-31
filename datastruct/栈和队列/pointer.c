#include <stdio.h>

// 总之，如果你想在函数内部修改值，并影响到函数外部，那么传参必须是一个合法的地址，
// 而不能是 NULL。否则就需要使用 指针的指针 来作为传参类型

typedef int* mint;

void change1(int x) {
    x = 10;
}

void change2(int* x) {
    *x = 10;
}

void change4(mint x) {
    *x = 10;
}

void change3(int** x) {
    **x = 10;
}

int main(int argc, char const* argv[]) {
    int i = 1;
    int j = 1;
    int m = 1;

    int x1 = 1;
    int* x2 = &i;
    int* x3 = &j;
    // mint x4; // 这样是不行的，会造成给空指针赋值的错误
    mint x4 = &m;

    change1(x1);
    printf("%d\n", x1);

    change2(x2);
    printf("%d, %d\n", *x2, i);

    change3(&x3);
    printf("%d, %d\n", *x3, j);

    change4(x4);
    printf("%d, %d\n", *x4, m);
    return 0;
}
