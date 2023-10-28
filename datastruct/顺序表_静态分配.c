#include <sequence_table.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#define MaxSize 10

typedef int ElemType;

typedef struct {
    ElemType data[MaxSize];
    int length;
} SqList;

void InitList(SqList* l) {
    l->length = 0;
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

    e = l->data[i - 1];
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
    // return l.data + i - 1;
}

int main(int argc, char const* argv[]) {
    SqList l;
    InitList(&l);
    printf("%p, %d, %d\n", &l, l.length, l.data[4]);
    return 0;
}
