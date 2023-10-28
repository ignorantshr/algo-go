#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include "sequence_table.h"

#define MaxSize 10

typedef int ElemType;

typedef struct {
    ElemType data[MaxSize];
    int length;
} SqList;

void InitList(SqList* l) {
    l->length = 0;
}

void DestoryList(SqList* l) {
    free(l);
}

// i 位序，from 1
bool ListInsert(SqList* l, int i, int e) {
    if (i > l->length + 1 || i < 1) {
        return false;
    }
    if (l->length == MaxSize) {
        return false;
    }

    for (int j = l->length; j > i - 1; j--) {
        l->data[j] = l->data[j - 1];
    }
    l->data[i - 1] = e;
    l->length++;

    return true;
}

bool ListDelete(SqList* l, int i, int* e) {
    if (i < 0 || i > l->length) {
        return false;
    }

    *e = l->data[i - 1];
    for (int j = i - 1; j < l->length; j--) {
        l->data[j] = l->data[j];
    }
    l->length--;
    return true;
}

ElemType GetElem(SqList l, int i) {
    if (i < 1 || i > l.length) {
        return -1;
    }

    return l.data[i - 1];
    // return *(l.data + i - 1);
}

ElemType LocateElem(SqList l, ElemType e) {
    for (int i = 0; i < l.length; i++) {
        if (l.data[i] == e) {
            return l.data[i];
        }
    }

    return -1;
}

void printList(SqList l) {
    printf("len: %d [", l.length);
    for (int i = 0; i < l.length; i++) {
        printf("%d ", l.data[i]);
    }
    printf("]\n");
}

int main(int argc, char const* argv[]) {
    SqList l;
    InitList(&l);
    ListInsert(&l, 1, 1);
    ListInsert(&l, 1, 2);
    ListInsert(&l, 1, 3);
    printf("%d\n", GetElem(l, 1));
    printf("%d\n", LocateElem(l, 1));
    // printf("%p, %d, %d\n", &l, l.length, l.data[4]);
    printList(l);
    return 0;
}
