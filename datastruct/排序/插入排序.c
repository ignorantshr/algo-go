#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include "check.c"

/*
1.空间复杂度： O(1)
2.时间复杂度:
    最好情况：原本就有序， O(n)
    最坏情况：原本为逆序， O(n^2)
    平均时间复杂度 O(n^2)
3.稳定性：稳定
*/

// 直接插入排序
void insert_sort(int a[], int n) {
    int i, j, tmp;
    for (i = 1; i < n; i++) {
        if (a[i] < a[i - 1]) {
            tmp = a[i];
            for (j = i; j > 0 && tmp < a[j - 1]; j--) {
                a[j] = a[j - 1];
            }
            a[j] = tmp;
        }
    }
}

// 直接插入排序(带哨兵)
// a[0] 不存放元素
void insert_sort_sentinel(int a[], int n) {
    int i, j, tmp;
    for (i = 2; i < n; i++) {
        if (a[i] < a[i - 1]) {
            a[0] = a[i];
            for (j = i - 1; a[0] < a[j]; j--) {  // 不再需要判断 j 是否越界
                a[j + 1] = a[j];
            }
            a[j + 1] = a[0];
        }
    }
}

// 折半插入排序
// 比起“直接插入排序”，比较关键字的次数减少了，但是移动元素的次数没变，整体来看时间复杂度依然是
// O(n^2)
void insert_sort_binary(int a[], int n) {
    int i, j, mid, tmp, low, high;
    for (i = 1; i < n; i++) {
        if (a[i] < a[i - 1]) {
            tmp = a[i];
            low = 0;
            high = i - 1;
            while (low <= high) {
                mid = (low + high) / 2;
                if (tmp < a[mid]) {
                    high = mid - 1;
                } else {
                    low = mid + 1;
                }
            }
            // [low:i-1] 右移一位
            for (j = i - 1; j >= low; j--) {
                a[j + 1] = a[j];
            }
            a[low] = tmp;
        }
    }
}

// 对链表进行插入排序 略
// 比起“直接插入排序”，移动元素的次数减少了，但是比较关键字的次数没变，整体来看时间复杂度依然是
// O(n^2)
void insert_sort_list(int a[], int n);

void testing() {
    tests("insert_sort", insert_sort);
    testsSentinel("insert_sort_sentinel", insert_sort_sentinel);
    tests("insert_sort_binary", insert_sort_binary);
}

int main(int argc, char const* argv[]) {
    testing();
    return 0;
}
