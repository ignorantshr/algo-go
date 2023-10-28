#include <sequence_table.h>
#include <stdlib.h>

#define InitSize 10

typedef struct {
    ElemType* data;
    int length;
    int maxSize;
} SeqList;

void InitList(SeqList* l) {
    l->data = (ElemType*)malloc(sizeof(ElemType) * InitSize);
    l->length = 0;
    l->maxSize = InitSize;
}

void IncreaseSize(SeqList* l, int len) {
    ElemType* p = l->data;
    l->data = (ElemType*)malloc(sizeof(ElemType) * (l->maxSize + len));

    for (int i = 0; i < l->maxSize; i++) {
        l->data[i] = p[i];
    }

    l->maxSize += len;
    free(p);
}

void DestoryList(SeqList* l) {
    free(l);
}
