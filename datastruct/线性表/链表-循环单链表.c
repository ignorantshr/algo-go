#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include "sequence_table.h"

typedef struct LNode {
    ElemType data;
    struct LNode* next;
} LNode, *LinkList;

LinkList initHeadList(int len, int vals[]) {
    LinkList l = (LinkList)malloc(sizeof(LNode));
    l->data = 0;
    l->next = NULL;
    LNode* r = l;

    for (int i = 0; i < len; i++) {
        LNode* node = (LNode*)malloc(sizeof(LNode));
        node->data = vals[i];
        r->next = node;
        r = node;
    }
    r->next = l;
    return l;
}

// 无头结点
LinkList initList(int len, int vals[]) {
    LinkList l;
    LNode* r = l;

    for (int i = 0; i < len; i++) {
        LNode* node = (LNode*)malloc(sizeof(LNode));
        node->data = vals[i];
        if (i == 0) {
            l = node;
        } else {
            r->next = node;
        }
        r = node;
    }
    r->next = l;
    return l;
}

bool EmptyHead(LinkList l) {
    return l->next == l;
}

bool Empty(LinkList l) {
    return l == NULL;
}

// 带头结点/不带头结点 后插
bool insert_after(LNode* cur, LNode* s) {
    if (cur == NULL || s == NULL) {
        return false;
    }

    s->next = cur->next;
    cur->next = s;
    return true;
}

// 对于循环单链表来说，如果我们让这个循环单链表的指针L不是指向头结点而是指向尾结点，从尾部找到头部时间复杂度仅为O(1)
// 同时由于L这个指针是指向尾部的，所以当我们对链表的尾部进行操作的时候，也可以在
// O(1) 的时间复杂度就找到我们要操作的位置，而不需要从头往后依次遍历

// l 指向链尾，插入头结点
bool insertAtHead(LinkList l, LNode* s) {
    if (l == NULL) {
        return false;
    }

    s->next = l->next;
    l->next = s;
    return true;
}

// l 指向链尾，插入尾结点
bool insertAtTail(LinkList l, LNode* s) {
    if (l == NULL) {
        return false;
    }

    s->next = l->next;
    l->next = s;
    l = s;
    return true;
}

// 带头结点删除
bool deleteHead(LNode* n) {
    if (n == NULL) {
        return false;
    }

    LNode* pre = n;
    while (pre->next != n) {
        pre = pre->next;
    }
    pre->next = n->next;
    free(n);
    return true;
}

// 不带头结点删除
bool delete(LNode* n) {
    if (n == NULL) {
        return false;
    }

    LNode* pre = n;
    while (pre->next != n) {
        pre = pre->next;
    }
    // if (pre != n) {  // 不只剩一个结点
    //     pre->next = n->next;
    // }
    pre->next = n->next;
    free(n);
    return true;
}

void printlist(LinkList l) {
    if (l == NULL) {
        printf("[]\n");
        return;
    }

    printf("[");
    LNode* n = l;
    do {
        printf("%d ", n->data);
        n = n->next;
    } while (n != l);
    printf("]\n");
}

int main(int argc, char const* argv[]) {
    int vals[2] = {1, 2};
    // LinkList l1 = initHeadList(2, vals);
    // printlist(l1);
    // LNode* n = (LNode*)malloc(sizeof(LNode));
    // n->data = 4;
    // n->next = NULL;
    // insert_after(l1->next, n);
    // printlist(l1);
    // deleteHead(l1->next->next->next);
    // printlist(l1);

    LinkList l2 = initList(2, vals);
    printlist(l2);
    // insert_after(l2->next, n2);
    // LinkList nl2 = l2->next;
    // delete (l2);
    // l2 = nl2;
    // printlist(l2);
    LNode* n2 = (LNode*)malloc(sizeof(LNode));
    n2->data = 4;
    // insertAtHead(l2->next, n2);
    insertAtTail(l2->next, n2);
    l2 = n2;
    printlist(l2);
}
