#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "sequence_table.h"

#define MaxSize 10

typedef struct Elem {
    ElemType value;
    int next;
} Elem, SLinklist[MaxSize + 1];

// typedef struct SNode {
//     Elem* datas[MaxSize + 1];
//     int len;
// } SNode, *SLinklist;

void initSLinklist(SLinklist l) {
    // SLinklist l = (SLinklist)malloc(sizeof(SNode));
    Elem head = *(Elem*)malloc(sizeof(Elem));
    head.next = -1;  // 链表尾部设置一个特殊值
    l[0] = head;
    for (int i = 1; i <= MaxSize; i++) {
        Elem* tmp = (Elem*)malloc(sizeof(Elem));
        tmp->next = -2;  // 空位都设置一个特殊值
        l[i] = *tmp;
    }
}

int getPrior(SLinklist l, int idx) {
    Elem pre = l[0];
    int res = 0;
    int i = 1;
    for (; i < idx; i++) {
        res = pre.next;
        pre = l[pre.next];
    }
    return res;
}

bool insert4Idx(SLinklist l, int idx, int val) {
    if (idx < 1 || idx > MaxSize) {
        return false;
    }

    // find a space
    int empty = 1;
    while (empty <= MaxSize) {
        if (l[empty].next == -2) {
            break;
        }
        empty++;
    }

    Elem* pre = &l[getPrior(l, idx)];  // O(n)

    Elem* cur = &l[empty];
    cur->value = val;
    cur->next = pre->next;
    pre->next = empty;
    // l[empty] = *cur;
    return true;
}

bool delete(SLinklist l, int idx) {
    if (idx < 1 || idx > MaxSize) {
        return false;
    }

    Elem* pre = &l[getPrior(l, idx)];  // O(n)

    Elem* cur = &l[pre->next];
    pre->next = cur->next;
    cur->next = -2;
    return true;
}

void printlist(SLinklist l) {
    Elem c = l[0];
    printf("[");
    while (true) {
        if (c.next == -1) {
            break;
        }
        c = l[c.next];
        printf("%d ", c.value);
    }
    printf("]\n");
}

int main(int argc, char const* argv[]) {
    SLinklist l;
    initSLinklist(l);
    insert4Idx(l, 1, 3);
    insert4Idx(l, 2, 2);
    insert4Idx(l, 3, 1);
    printlist(l);
    delete (l, 1);
    printlist(l);
    delete (l, 2);
    printlist(l);
    delete (l, 1);
    printlist(l);

    return 0;
}
