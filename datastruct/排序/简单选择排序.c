#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "check.c"
/*
1.空间复杂度： O(1)

2.时间复杂度： O(n^2) 无论有序、逆序、还是乱序，一定需要 n−1 趟处理
总共需要对比关键字 (n−1)+(n−2)+...+1= n(n−1)/2 次
元素交换次数 < n−1

3.稳定性：不稳定

4.适用性：既可以用于顺序表，也可以用于链表
 */
void select_sort(int a[], int n) {
    int minidx;
    for (int m = 0; m < n - 1; m++) {  // n-1 轮
        minidx = m;
        for (int i = m + 1; i < n; i++) {
            if (a[minidx] > a[i]) {
                minidx = i;
            }
        }
        if (minidx != m) {
            int tmp = a[minidx];
            a[minidx] = a[m];
            a[m] = tmp;
        }
    }
}

void testing() {
    tests("select_sort", select_sort);
}

int main(int argc, char const* argv[]) {
    testing();
    return 0;
}
