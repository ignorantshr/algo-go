#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include "sequence_table.h"

typedef struct DNode {
    ElemType data;
    struct DNode* pre;
    struct DNode* next;
} DNode, *DLinkList;

// 初始化空的循环双链表
DLinkList InitDLinkList(int len, int vals[]) {
    DLinkList l = (DLinkList)malloc(sizeof(DNode));  // 分配一个头结点
    if (l == NULL)  // 内存不足，分配失败
        return false;
    l->next = l;
    l->pre = l;
    DNode* r = l;
    for (int i = 0; i < len; i++) {
        DNode* n = (DNode*)malloc(sizeof(DNode));
        n->data = vals[i];
        r->next = n;
        n->pre = r;
        r = n;
    }
    r->next = l;
    l->pre = r;
    return l;
}

bool Empty(DLinkList l) {
    return l->next == l;
}

// 前插
bool InsertPriorNode(DNode* p, DNode* s) {
    if (p == NULL || s == NULL) {
        return false;
    }

    s->pre = p->pre;
    p->pre->next = s;
    s->next = p;
    p->pre = s;
    return true;
}

// 后插
bool InsertNextNode(DNode* p, DNode* s) {
    if (p == NULL || s == NULL) {
        return false;
    }

    s->next = p->next;
    p->next->pre = s;
    p->next = s;
    s->pre = p;
    return true;
}

bool delete(DNode* n) {
    if (n == NULL) {
        return false;
    }

    n->pre->next = n->next;
    n->next->pre = n->pre;
    free(n);
    return true;
}

void printlist(DLinkList l) {
    if (Empty(l)) {
        printf("[]\n");
        return;
    }

    printf("[");
    DNode* n = l->next;
    while (n != l) {
        printf("%d ", n->data);
        n = n->next;
    }
    printf("]\n");
}

int main(int argc, char const* argv[]) {
    int vals[2] = {1, 2};
    DLinkList l = InitDLinkList(2, vals);
    printlist(l);
    // delete (l->next);
    // printlist(l);
    // delete (l->next);
    // printlist(l);

    DNode* n = (DNode*)malloc(sizeof(DNode));
    n->data = 4;
    n->next = NULL;
    // InsertNextNode(l->next->next, n);
    InsertPriorNode(l->next->next->next, n);
    printlist(l);
    return 0;
}
