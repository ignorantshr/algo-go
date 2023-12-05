#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "check.c"

#define MaxSize 100

// 树 目录下中也有小顶堆的实现

typedef struct heap {
    int data[MaxSize + 1];  // data[0] 不使用
    int size;
    bool (*compare)(int data[], int i, int j);
} heap;

void upFloat(heap* h, int i) {
    while (i > 1 && h->compare(h->data, i, i / 2)) {
        int tmp = h->data[i];
        h->data[i] = h->data[i / 2];
        h->data[i / 2] = tmp;
        i = i / 2;
    }
}

void sink(heap* h, int i) {
    int j = i * 2;
    while (j <= h->size) {
        if (j + 1 <= h->size && h->compare(h->data, j + 1, j)) {
            j++;
        }

        if (!h->compare(h->data, i, j)) {
            int tmp = h->data[i];
            h->data[i] = h->data[j];
            h->data[j] = tmp;
            i = j;
            j = i * 2;
        } else {
            break;
        }
    }
}

void insert(heap* h, int x) {
    if (h->size == MaxSize) {
        return;
    }

    h->data[++h->size] = x;

    upFloat(h, h->size);
}

int pop(heap* h) {
    if (h->size == 0) {
        return -1;
    }

    int res = h->data[1];
    h->data[1] = h->data[h->size];
    --h->size;
    if (h->size > 1) {
        sink(h, 1);
    }

    return res;
}

heap createHeap(int a[], int n, bool (*less)(int data[], int i, int j)) {
    heap h = {.size = 0, .compare = less};
    for (int i = 0; i < n; i++) {
        insert(&h, a[i]);
    }
    return h;
}

bool smaller(int data[], int i, int j) {
    return data[i] < data[j];
}

bool bigger(int data[], int i, int j) {
    return data[i] > data[j];
}

void testHeap() {
    srand(time(NULL));
    const int n = 10;
    int a[n];
    for (int i = 0; i < n; i++) {
        a[i] = randn(30);
    }

    heap h = createHeap(a, n, smaller);

    int res[n];
    for (int i = 0; i < n; i++) {
        res[i] = pop(&h);
    }
    parray(res, n);
    assert(checkAscend(res, n));
}

int main(int argc, char const* argv[]) {
    testHeap();
    return 0;
}
