#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "thread_tree.h"

// https://blog.csdn.net/zimuzi2019/article/details/126273611

void visit(ThreadNode* p) {
    printf("%d ", p->data);
}

// ====找后继结点=====

// 找到以p为根的子树中，第一个被中序遍历的结点
ThreadNode* firstNode(ThreadNode* p) {
    // 循环找到最左下结点（不一定是叶结点）
    while (p->ltag == 0) {
        p = p->lchild;
    }

    return p;
}

// 在中序线索二叉树中找到结点找到 p 的后继结点
ThreadNode* nextNode(ThreadNode* p) {
    // 右子树中的最左下结点
    if (p->rtag == 0) {
        return firstNode(p->lchild);
    }

    return p->rchild;
}

// 中序线索二叉树进行中序遍历。这是利用线索实现的非递归算法，空间复杂度仅为O(1)
void inOrderScan(ThreadTree t) {
    for (ThreadNode* cur = firstNode(t); cur != NULL; cur = nextNode(cur)) {
        visit(cur);
    }
}

// ====找后继结点=====

// ====找前驱结点=====

// 找到以p为根的子树中，最后一个被中序遍历的结点
ThreadNode* lastNode(ThreadNode* p) {
    while (p->ltag == 0) {
        p = p->rchild;
    }
    return p;
}

ThreadNode* preNode(ThreadNode* p) {
    if (p->ltag == 0) {
        return lastNode(p->lchild);
    }
    return p->lchild;
}

// 中序线索二叉树进行逆向中序遍历。这是利用线索实现的非递归算法，空间复杂度仅为O(1)
void inOrderReverseScan(ThreadTree t) {
    for (ThreadNode* cur = lastNode(t); cur != NULL; cur = preNode(cur)) {
        visit(cur);
    }
}

// ====找前驱结点=====
