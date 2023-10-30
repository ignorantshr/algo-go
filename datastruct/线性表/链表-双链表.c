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
DLinkList initDList(int len, int vals[]) {
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

void DestroyDlist(DLinkList l) {
    while (l != NULL) {
        DNode* next = l->next;
        free(l);
        l = next;
    }
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

// 删除带头结点的链表
bool deleteNodeHead(DNode* n) {
    if (n == NULL) {
        return false;
    }

    n->pre->next = n->next;
    if (n->next != NULL) {
        n->next->pre = n->pre;
    }
    free(n);
    return true;
}

// 删除不带头结点的链表
bool deleteNode(DNode* n) {
    if (n == NULL) {
        return false;
    }

    if (n->pre != NULL) {
        n->pre->next = n->next;
    }
    if (n->next != NULL) {
        n->next->pre = n->pre;
    }
    free(n);
    return true;
}

void printlist(DLinkList l) {
    printf("[");
    for (DNode* n = l; n != NULL; n = n->next) {
        printf("%d ", n->data);
    }
    printf("]\n");
}

int main(int argc, char const* argv[]) {
    int vals[] = {1, 2, 3};
    DLinkList hl1 = initHeadDList(3, vals);
    printlist(hl1);
    DNode* n = (DNode*)malloc(sizeof(DNode));
    n->data = 4;
    n->next = NULL;
    // insert_after(hl1, n);
    // deleteNodeHead(hl1->next);
    deleteNodeHead(hl1->next->next->next);
    printlist(hl1);

    DLinkList l2 = initDList(3, vals);
    printlist(l2);
    // insert_before(l2->next, n);
    insert_before(l2, n);
    l2 = n;
    printlist(l2);
    // deleteNode(l2->next);
    // n = l2->next;
    // deleteNode(l2);
    // l2 = n;
    deleteNode(l2->next->next->next);
    printlist(l2);
}
