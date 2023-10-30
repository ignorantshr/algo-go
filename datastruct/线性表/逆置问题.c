#include <stdio.h>

void printarray(int a[], int n);

// 将长度为n的数组的前端k个元素逆序后移动到数组后端，要求原数组中的数据不丢失，其余元素位置无关紧要
void reverse(int a[], int left, int right, int k) {
    int temp;
    for (int i = left, j = right; i < left + k && i < j; ++i, --j) {
        temp = a[i];
        a[i] = a[j];
        a[j] = temp;
    }
}

// 将长度为n的数组的前端k个元素保持原序移动到数组后端，要求原数组中的数据不丢失，其余元素位置无关紧要
void moveToEnd(int a[], int k, int n) {
    reverse(a, 0, k - 1, k);  // 先反转前k个数字
    reverse(a, 0, n - 1, k);  // 再反转+移动位置
}

// 循环左移p个位置
void moveP(int a[], int n, int p) {
    reverse(a, 0, p - 1, p);  // 先反转前p个数字
    // printarray(a, n);
    reverse(a, p, n - 1, n - p);  // 再反转剩余的数字
    // printarray(a, n);
    reverse(a, 0, n - 1, n);  // 反转整个数组
}

void printarray(int a[], int n) {
    for (int i = 0; i < n; i++) {
        printf("%d->", a[i]);
    }
    printf("\n");
}

int main(int argc, char** argv) {
    int vals[10] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    int k = 3;
    int n = 10;
    int p = 2;
    // reverse(vals, 0, n - 1, k);  // 10->9->8->4->5->6->7->3->2->1
    // moveToEnd(vals, k, n);  // 10->9->8->4->5->6->7->1->2->3
    moveP(vals, n, p);  // 3->4->5->6->7->8->9->10->1->2
    printarray(vals, 10);
    return 0;
}
