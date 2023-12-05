#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "check.c"

/* “2路”归并 —— 每选出一个小元素就需对比关键字1次。“4路”归并 ——
每选出一个小元素就需对比关键字3次。故 m 路归并，每选出一个元素需要对比关键字 m-1
 */

/*
2路归并的“归并树”——形态上就是一棵倒立的二叉树

n个元素进行2路归并排序，归并趟数=[log_2n]
每趟归并时间复杂度为 O(n) ,则算法时间复杂度为 O(nlog_2n)

空间复杂度=O(n) ，来自于辅助数组B

两个元素相等时，优先使用靠前的那个。归并排序是稳定的
 */

void _merge_sort(int a[], int low, int high);
void _merge(int a[], int low, int mid, int high);

void merge_sort(int a[], int len) {
    _merge_sort(a, 0, len - 1);
}

void _merge_sort(int a[], int low, int high) {
    if (low >= high) {  // 必须放在这里，不可下放到 _merge
        return;
    }

    // printf("before [%d:%d] ", low, high);
    // parray(a, 10);

    int mid = (low + high) / 2;
    _merge_sort(a, low, mid);
    _merge_sort(a, mid + 1, high);
    _merge(a, low, mid, high);

    // printf("after [%d:%d] ", low, high);
    // parray(a, 10);
}

// merge [low,mid] [mid+1,high]
void _merge(int a[], int low, int mid, int high) {
    int b[high - low + 1];
    for (int i = low; i <= high; i++) {
        b[i - low] = a[i];
    }

    // 也可以直接使用同一个辅助数组，不用每次新建一个数组，也省去了下标对应的麻烦
    // int *B=(int *)malloc(n*sizeof(int)); //辅助数组B

    // printf("merge [%d:%d:%d] ", low, mid, high);
    // parray(b, high - low + 1);

    int i, j;  // 对应 b 的下标
    int k;     // 对应 a 的下标
    for (i = 0, j = mid + 1 - low, k = low; i <= mid - low && j <= high - low;
         k++) {
        if (b[i] < b[j]) {
            a[k] = b[i];
            i++;
        } else {
            a[k] = b[j];
            j++;
        }
    }

    while (i <= mid - low) {
        a[k++] = b[i++];
    }
    while (j <= high - low) {
        a[k++] = b[j++];
    }
}

void testing() {
    tests("merge_sort", merge_sort);
}

int main(int argc, char const* argv[]) {
    testing();
    return 0;
}
