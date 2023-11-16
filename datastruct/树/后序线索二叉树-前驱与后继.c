#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "thread_tree.h"

// https://blog.csdn.net/zimuzi2019/article/details/126273611

void visit(ThreadNode* p) {
    printf("%d ", p->data);
}

// ====找前驱结点=====

// 在后序线索二叉树中找到结点找到 p 的前驱结点
// 必有左孩子
ThreadNode* preNode(ThreadNode* p) {
    if (p->ltag == 0) {
        if (p->rchild != NULL) {  // 右孩子优先
            return p->rchild;
        }
        return p->lchild;
    }

    return p->lchild;
}

void reverse() {}

// 后序线索二叉树进行后序遍历。这是利用线索实现的非递归算法，空间复杂度仅为O(1)
void postOrderScan(ThreadTree t) {
    for (ThreadNode* cur = t; cur != NULL; cur = preNode(cur)) {
        visit(cur);
    }
    // 后序遍历需要翻转
    reverse();
}

// 后序线索二叉树进行逆向后序遍历。这是利用线索实现的非递归算法，空间复杂度仅为O(1)
void postOrderReverseScan(ThreadTree t) {
    for (ThreadNode* cur = t; cur != NULL; cur = preNode(cur)) {
        visit(cur);
    }
}

// ====找前驱结点=====

// ====找后继结点=====

// 后序遍历的顺序为【左，右，根】，如果 p 没有后继线索，
// 在子结点中是无法找到后继结点的。有两种办法解决这个问题：
//  1. 从头重新后序遍历二叉树
//  2. 把二叉链表改为三叉链表，给各个结点设置一个指向它父结点的指针

ThreadNode* nextNode2(ThreadNode* p) {
    ThreadNode* parent = p->parent;

    // 1.找不到父节点，则没有后继结点
    if (parent == NULL) {
        return NULL;
    }

    // 2.找到父节点且是其右孩子结点
    if (parent->rchild == p) {
        return parent;
    }

    // 3.找到父节点且是其左孩子结点且右兄弟结点为空
    if (parent->lchild == p && parent->rchild == NULL) {
        return parent;
    }

    // 4.找到父节点且是其左孩子结点且右兄弟结点非空
    // 找到右兄弟子树的第一个后序遍历的结点，即是 p 的后继结点
    // todo 以下代码待验证
    ThreadNode* post = parent->rchild;
    while (post->ltag == 0) {
        if (post->lchild != NULL) {  // 优先遍历左子树
            post = post->lchild;
        } else if (post->rchild != NULL) {
            post = post->rchild;
        }
    }
    return post;
}

ThreadNode* nextNode(ThreadNode* p) {
    if (p->rtag == 0) {
        return nextNode2(p->lchild);
    }
    return p->rchild;
}

// ====找后继结点=====
