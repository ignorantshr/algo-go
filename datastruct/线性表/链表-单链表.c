#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include "sequence_table.h"

typedef struct LNode {
    ElemType data;
    struct LNode* next;
} LNode, *LinkList;

// 头插法，注意！！！：数据倒序
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

// 头插法，注意！！！：数据倒序
// 不带头结点
LinkList InitHeadList2(int len, int vals[]) {
    LinkList l;
    for (int i = 0; i < len; i++) {
        LNode* node = (LNode*)malloc(sizeof(LNode));
        node->data = vals[i];
        node->next = l;
        l = node;
    }
    return l;
}

// 尾插法
// 有头结点
LinkList InitTailList1(int len, int vals[]) {
    LinkList l = (LinkList)malloc(sizeof(LNode));
    LNode* r = l;

    for (int i = 0; i < len; i++) {
        LNode* n = (LNode*)malloc(sizeof(LNode));
        n->data = vals[i];
        r->next = n;
        r = n;
    }
    r->next = NULL;
    return l;
}

// 尾插法
// 没有头结点
LinkList InitTailList2(int len, int vals[]) {
    LinkList l;
    LinkList r = l;

    for (int i = 0; i < len; i++) {
        LNode* n = (LNode*)malloc(sizeof(LNode));
        n->data = vals[i];
        if (i == 0) {
            l = n;
        } else {
            r->next = n;
        }
        r = n;
    }
    r->next = NULL;
    return l;
}

// ---------------------------------------

// 带头结点

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
LinkList ListInsert1_pre(LinkList l, int i, int e) {
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

// 无头结点

LNode* GetElem2(LinkList l, int i) {
    if (i < 1) {
        return NULL;
    }

    LNode* p = l;
    while (p != NULL && i > 1) {
        p = p->next;
    }

    if (p == NULL) {
        return NULL;
    }

    return p;
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

void printlist(LinkList l) {
    printf("[");
    for (LNode* n = l; n != NULL; n = n->next) {
        printf("%d ", n->data);
    }
    printf("]\n");
}

int main(int argc, char const* argv[]) {
    int vals[] = {1, 2, 3};
    LinkList hl1 = InitHeadList1(3, vals);
    printlist(hl1);

    // int vals[] = {1, 2, 3};
    LinkList hl2 = InitHeadList2(3, vals);
    printlist(hl2);

    // int vals[] = {1};
    LinkList l1 = InitTailList1(1, vals);

    printlist(l1);
    LinkList l2 = InitTailList2(1, vals);
    printlist(l2);
    return 0;
}
