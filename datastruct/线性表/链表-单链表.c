#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include "sequence_table.h"

typedef struct LNode {
    ElemType data;
    LNode* next;
} LNode, *LinkList;

// 头插法，数据倒序
// 带头结点
LinkList InitHeadList1(int len, int vals[]) {
    LinkList l = (LinkList)malloc(sizeof(LNode) * len);
    l->next = NULL;
    for (int i = 0; i < len; i++) {
        LNode* node = (LNode*)malloc(sizeof(LNode));
        node->data = vals[i];
        node->next = l->next;
        l->next = node;
    }
    return l;
}

LNode* GetElem1(LinkList l, int i) {
    if (i == 0) {
        return l;
    }

    if (i < 1) {
        return NULL;
    }

    for (LNode* node = l->next; node != NULL; node = node->next) {
        if (i == 1) {
            return node;
        }
        i--;
    }

    return NULL;
}

LNode* LocateElem1(LinkList l, ElemType e) {
    for (LNode* node = l->next; node != NULL; node = node->next) {
        if (l->data == e) {
            return node;
        }
    }
    return NULL;
}

// 前插法
LinkList* ListInsert1_pre(LinkList l, int i, int e) {
    LNode* p = GetElem1(l, i - 1);  // 对于前驱节点来说就是后插
    if (p == NULL) {
        return l;
    }

    LNode* node = (LNode*)malloc(sizeof(LNode));
    node->data = e;
    node->next = p->next;
    p->next = node;
    return l;
}

// 指定节点后插
bool ListInsert1_after(LNode* p, ElemType e) {
    if (p == NULL) {
        return false;
    }

    LNode* node = (LNode*)malloc(sizeof(LNode));
    node->data = e;
    node->next = p->next;
    p->next = node;
    return true;
}

LinkList ListDelete4Idx(LinkList l, int i, int* e) {
    LNode* p = GetElem1(l, i - 1);
    if (p == NULL || p->next == NULL) {
        return l;
    }

    LNode* cur = p->next;
    *e = cur->data;
    p->next = cur->next;
    free(cur);
    return l;
}

bool ListDelete4Spec(LNode* n) {
    if (n == NULL) {
        return false;
    }

    // 这里直接free并不会变成null，此时必须找到前驱节点然后删除
    // if (n->next == NULL) {
    //     free(n);
    //     return false;
    // }

    LNode* cur = n->next;
    n->data = n->next->data;
    n->next = n->next->next;
    free(cur);
    return true;
}

// ------------------------------------------------

// 头插法，数据倒序
// 无头结点
LinkList InitHeadList2(int len, int vals[]) {
    LinkList l;
    for (int i = 1; i < len; i++) {
        LNode* node = (LNode*)malloc(sizeof(LNode));
        node->data = vals[i];
        node->next = l;
        l = node;
    }
    return l;
}

// 前插法
bool ListInsert2_pre(LinkList l, int i, int e) {
    if (i == 1) {  // 第一个节点特殊处理
        LNode* node = (LNode*)malloc(sizeof(LNode));
        node->data = e;
        node->next = l;
        l = node;
        return true;
    }

    LNode* p = GetElem2(l, i - 1);  // 对于前驱节点来说就是后插
    if (p == NULL) {
        return false;
    }

    LNode* node = (LNode*)malloc(sizeof(LNode));
    node->data = e;
    node->next = p->next;
    p->next = node;
    return true;
}

// 指定节点后插
bool ListInsert2_after(LNode* p, ElemType e) {
    if (p == NULL) {
        return false;
    }

    LNode* node = (LNode*)malloc(sizeof(LNode));
    node->data = e;
    node->next = p->next;
    p->next = node;
    return true;
}

bool List2Delete4Idx(LinkList l, int i, int* e) {
    if (i < 1 || l == NULL) {
        return false;
    }
    if (i == 1) {
        LNode* cur = l;
        *e = cur->data;
        l = cur->next;
        free(cur);
        return true;
    }

    LNode* p = GetElem2(l, i - 1);  // O(n)
    if (p == NULL) {
        return false;
    }

    if (p->next == NULL) {
        return false;
    }

    LNode* cur = p->next;
    *e = cur->data;
    p->next = cur->next;
    free(cur);
    return l;
}

bool List2Delete4Spec(LNode* n) {
    if (n == NULL) {
        return false;
    }

    // 这里直接free并不会变成null，此时必须找到前驱节点然后删除
    // if (n->next == NULL) {
    //     free(n);
    //     return false;
    // }

    // O(1)
    LNode* cur = n->next;
    n->data = n->next->data;
    n->next = n->next->next;
    free(cur);
    return true;
}

LNode* GetElem2(LinkList l, int i) {
    if (i < 1) {
        return -1;
    }

    LNode* p = l;
    while (p != NULL && i > 1) {
        p = p->next;
    }

    if (p == NULL) {
        return -1;
    }

    return p->data;
}

void printlist(LinkList l) {
    printf("[\n");
    for (LNode* n = l; n != NULL; n = n->next) {
        printf("%d ", n->data);
    }
    printf("]\n");
}
