#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include "sequence_table.h"

typedef struct DNode {
    ElemType data;
    struct DNode* pre;
    struct DNode* next;
} DNode, *DLinkList;

// 带头结点的双链表
DLinkList initHeadDList(int len, int vals[]) {
    DLinkList l = (DLinkList)malloc(sizeof(DNode));
    l->pre = NULL;
    l->next = NULL;
    DNode* tail = l;

    for (int i = 0; i < len; i++) {
        DNode* tmp = (DNode*)malloc(sizeof(DNode));
        tmp->data = vals[i];
        tail->next = tmp;
        tmp->pre = tail;
        tail = tmp;
    }
    tail->next = NULL;
    return l;
}

// 不带头结点的双链表
DLinkList initHeadDList(int len, int vals[]) {
    DLinkList l;
    DNode* tail = l;

    for (int i = 0; i < len; i++) {
        DNode* tmp = (DNode*)malloc(sizeof(DNode));
        tmp->data = vals[i];
        if (i == 0) {
            l = tmp;
        } else {
            tail->next = tmp;
            tmp->pre = tail;
        }
        tail = tmp;
    }
    tail->next = NULL;
    return l;
}

bool insert_before(DNode* cur, DNode* n) {
    if (cur == NULL || n == NULL) {
        return false;
    }

    n->next = cur;
    if (cur->pre != NULL) {
        cur->pre->next = n;
    }
    n->pre = cur->pre;
    cur->pre = n;
    return true;
}

bool insert_after(DNode* cur, DNode* n) {
    if (cur == NULL || n == NULL) {
        return false;
    }

    n->next = cur->next;
    if (cur->next != NULL) {
        cur->next->pre = n;
    }
    cur->next = n;
    n->pre = cur;
    return true;
}
