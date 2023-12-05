#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "check.c"

/*
堆排序的空间复杂度=O(1)

建堆的过程，关键字对比次数不超过 4n，建堆时间复杂度=O(n)

每趟交换和建堆过程，根节点最多“下坠” h−1层，
每下坠一层最多只需对比关键字2次，因此每一趟排序时间复杂度不超过
O(h)=O(log_2n) 。共 n-1 趟，总的时间复杂度=O(nlog_2n)

堆排序的时间复杂度=O(n)+O(nlog_2n)=O(nlog_2n)

堆排序是不稳定的
*/

// 将以 k 为根的子树调整为大根堆
void adjustHeap(int a[], int k, int len) {
    a[0] = a[k];
    for (int i = 2 * k; i <= len; i *= 2) {
        if (i < len && a[i + 1] > a[i]) {
            i++;
        }
        if (a[i] > a[0]) {
            a[k] = a[i];
            k = i;
        } else {
            break;
        }
    }
    a[k] = a[0];
}

void buildMaxHeap(int a[], int len) {
    for (int i = len / 2; i >= 1; i--) {
        adjustHeap(a, i, len);
    }
}

// 堆排序
// a[0] 不存储元素
// 注意：这种写法将
// 基于“大根堆”的堆排序得到“递增序列”，而基于“小根堆”的堆排序得到“递减序列”
void heap_sort(int a[], int len) {
    buildMaxHeap(a, len);

    for (int i = len; i > 0; i--) {
        // 把最大的放到最后
        a[0] = a[i];
        a[i] = a[1];
        a[1] = a[0];
        // 调整最大的数字到数组第一位
        adjustHeap(a, 1, i - 1);
    }
}

void testing() {
    testsSentinel("heap_sort", heap_sort);
}

int main(int argc, char const* argv[]) {
    testing();
    return 0;
}
