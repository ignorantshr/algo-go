#include <stdbool.h>
#include <stdio.h>
#include <time.h>

int randn(int n) {
    int max = n;  // 最大值
    return rand() % (max + 1);
}

void parray(int a[], int n) {
    for (int i = 0; i < n; i++) {
        printf("%d ", a[i]);
    }
    printf("\n");
}

bool checkAscend(int a[], int n) {
    for (int i = 0; i < n - 1; i++) {
        if (a[i] > a[i + 1]) {
            return false;
        }
    }
    return true;
}

void tests(char* name, void (*sort)(int a[], int n)) {
    int min = 0;    // 最小值
    int max = 100;  // 最大值
    srand(time(NULL));

    const int n = 10;
    int a[n];
    int t[n];  // 保存原始数组
    printf("-------%s--------\n", name);

    for (int m = 0; m < 100; m++) {
        for (int i = 0; i < n; i++) {
            a[i] = rand() % (max - min + 1) + min;
        }
        memccpy(t, a, n, sizeof(a));

        sort(a, n);

        if (!checkAscend(a, n)) {
            parray(t, n);
            parray(a, n);
            printf("[%s] fail\n\n", name);
            return;
        }
    }

    // 选择一个输出
    parray(t, n);
    parray(a, n);
    printf("[%s] ok\n\n", name);
}

bool checkAscendSentinel(int a[], int n) {
    for (int i = 1; i < n - 1; i++) {
        if (a[i] > a[i + 1]) {
            return false;
        }
    }
    return true;
}

void testsSentinel(char* name, void (*sort)(int a[], int n)) {
    int min = 0;    // 最小值
    int max = 100;  // 最大值
    srand(time(NULL));

    const int n = 10;
    int a[n];
    a[0] = -1;
    for (int i = 1; i < n; i++) {
        a[i] = rand() % (max - min + 1) + min;
    }

    printf("-------%s--------\n", name);
    parray(a, n);
    sort(a, n);

    if (!checkAscendSentinel(a, n)) {
        parray(a, n);
        printf("[%s] fail\n\n", name);
    } else {
        parray(a, n);
        printf("[%s] ok\n\n", name);
    }
}
