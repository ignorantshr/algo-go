#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "check.c"

/*
1.时间复杂度=O(n×递归层数)
最好时间复杂度：O(nlog_2n)
最坏时间复杂度：O(n^2)
平均时间复杂度：O(nlog_2n)

2.空间复杂度= O(递归深度)
最好空间复杂度：O(log_2n)
最坏空间复杂度：O(n)

若每次选中的“枢轴”将待排序序列划分为均匀的两个部分，则递归深度最小，算法效率最高

若每次选中的“枢轴”将待排序序列划分为很不均匀的两个部分，则会导致递归深度增加，算法效率变低

若初始序列有序或逆序，则快速排序的性能最差（因为每次选择的都是最靠边的元素，可在此优化）

快速排序是所有内部排序算法中平均性能最优的排序算法

3.稳定性：不稳定

 */

void _sort(int a[], int l, int r);

// 快速排序
void quick_sort(int a[], int n) {
    _sort(a, 0, n - 1);
}

void _sort(int a[], int l, int r) {
    if (l >= r) {
        return;
    }

    int base = a[l];
    int low = l;
    int high = r;

    // 这种写法必须先找右边更小的元素
    while (low < high) {
        // find smaller on the right
        while (low < high && a[high] >= base) {
            high--;
        }
        a[low] = a[high];
        // 利用可覆盖的位置移动元素，此时 high 变成了可覆盖的位置

        // find bigger on the left
        while (low < high && a[low] <= base) {
            low++;
        }
        a[high] = a[low];
        // 因为 a[high]
        // 已经移动到前面去了，所以这个位置可以直接用来存储更大的元素，然后 low
        // 又变成了可覆盖的位置
    }

    a[low] = base;

    _sort(a, l, low - 1);
    _sort(a, low + 1, r);
}

void testing() {
    tests("quick_sort", quick_sort);
}

int main(int argc, char const* argv[]) {
    testing();
    return 0;
}
