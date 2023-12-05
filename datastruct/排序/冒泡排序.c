#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "check.c"

/*
1.空间复杂度： O(1)

2.时间复杂度：

最好情况（有序）比较次数 = n−1 ；交换次数 = 0, 最好时间复杂度= O(n)
最坏情况（逆序）比较次数 = (n−1)+(n−2)+...+1 = n(n−1)2 = 交换次数

每次交换都需要移动元素3次

最坏时间复杂度= O(n^2)
平均时间复杂度= O(n^2)

3.稳定性： 稳定的

是否适用于链表？
可从前往后“冒泡”，每一趟将更大的元素“冒’'到链尾
 */

void bubble_sort(int a[], int n) {
    bool change;
    for (int m = 1; m < n; m++) {  // n-1 轮
        change = false;
        for (int i = n - 1; i >= m; i--) {
            if (a[i] < a[i - 1]) {
                int tmp = a[i];
                a[i] = a[i - 1];
                a[i - 1] = tmp;
                change = true;
            }
        }
        if (!change) {
            return;
        }
    }
}

void testing() {
    tests("bubble_sort", bubble_sort);
}

int main(int argc, char const* argv[]) {
    testing();
    return 0;
}
