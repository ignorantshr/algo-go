#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "check.c"

#define max(a, b) (a > b ? a : b)

/*
先追求表中元素部分有序，再逐渐逼近全局有序

先将待排序表分割成若干形如 L[i,i+d,i+2d,...,i+kd] 的 “特殊”
子表，对各个子表分别进行直接插入排序。缩小增量 d，重复上述过程，直到 d=1 为止
 */

/*
1.空间复杂度： O(1)
2.时间复杂度：和增量序列 d_1,d_2,d_3...
的选择有关，目前无法用数学手段证明确切的时间复杂度。最坏时间复杂度为 O(n^2)
，当n在某个范围内时，可达 O(n^{1.3})
3.稳定性：不稳定
4.适用性：仅适用于顺序表，不适用于链表
 */

// 希尔排序
// 局部有序 -> 整体有序
void shell_sort(int a[], int n) {
    int d = n / 2;
    int tmp, i, j;

    while (d >= 1) {
        for (i = d; i < n; i++) {
            if (a[i] < a[i - d]) {
                tmp = a[i];
                for (j = i; j - d >= 0 && tmp < a[j - d]; j -= d) {
                    a[j] = a[j - d];
                }
                a[j] = tmp;
            }
        }

        // printf("%d | ", d);
        // parray(a, n);
        d = d / 2;
    }
}

void testing() {
    tests("shell_sort", shell_sort);
}

int main(int argc, char const* argv[]) {
    testing();
    return 0;
}
