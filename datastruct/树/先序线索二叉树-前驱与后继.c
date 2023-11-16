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

// 在先序线索二叉树中找到结点找到 p 的后继结点
// 必有右孩子
ThreadNode* nextNode(ThreadNode* p) {
    if (p->rtag == 0) {
        if (p->lchild != NULL) {
            return p->lchild;
        }
        return p->rchild;
    }

    return p->rchild;
}

// 先序线索二叉树进行先序遍历。这是利用线索实现的非递归算法，空间复杂度仅为O(1)
void preOrderScan(ThreadTree t) {
    for (ThreadNode* cur = t; cur != NULL; cur = nextNode(cur)) {
        visit(cur);
    }
}

// ====找后继结点=====

// ====找前驱结点=====

// 先序遍历的顺序为【根，左，右】，如果 p 没有前驱线索，
// 在子结点中是无法找到前驱结点的。有两种办法解决这个问题：
//  1. 从头重新先序遍历二叉树
//  2. 把二叉链表改为三叉链表，给各个结点设置一个指向它父结点的指针

ThreadNode* preNode2(ThreadNode* p) {
    ThreadNode* parent = p->parent;

    // 1.找不到父节点，则没有前驱结点
    if (parent == NULL) {
        return NULL;
    }

    // 2.找到父节点且是其左孩子结点
    if (parent->lchild == p) {
        return parent;
    }

    // 3.找到父节点且是其右孩子结点且左兄弟结点为空
    if (parent->rchild == p && parent->lchild == NULL) {
        return parent;
    }

    // 4.找到父节点且是其右孩子结点且左兄弟结点非空
    // 找到左兄弟子树的最后一个先序遍历的结点，即是 p 的前驱结点
    // todo 以下代码待验证
    ThreadNode* pre = parent->lchild;
    while (pre->rtag == 0) {
        if (pre->rchild != NULL) {  // 优先遍历右子树
            pre = pre->rchild;
        } else if (pre->lchild != NULL) {
            pre = pre->lchild;
        }
    }
    return pre;
}

ThreadNode* preNode(ThreadNode* p) {
    if (p->ltag == 0) {
        return preNode2(p);
    }
    return p->lchild;
}

// 先序线索二叉树进行逆向先序遍历。这是利用线索实现的非递归算法，空间复杂度仅为O(1)
void preOrderReverseScan(ThreadTree t) {
    for (ThreadNode* cur = t; cur != NULL; cur = preNode(cur)) {
        visit(cur);
    }
}

// ====找前驱结点=====
