#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "thread_tree.h"

// 二叉树找某个结点的前驱和后继结点不方便，由此有了线索二叉树
// 二叉树可以被称为二叉链表，线索化后又可以被称为线索链表
// 中序线索二叉树
// 先序线索二叉树
// 后序线索二叉树

ThreadNode* pre;

void visit(ThreadNode* cur) {
    if (cur->lchild == NULL) {
        cur->ltag = 1;
        cur->lchild = pre;
    }
    if (pre != NULL && pre->rchild == NULL) {
        pre->rtag = 1;
        pre->rchild = cur;
    }
    pre = cur;
}

// 二叉树中序线索化
void inOrderThread(ThreadTree t) {
    if (t == NULL) {
        return;
    }

    inOrderThread(t->lchild);
    visit(t);
    inOrderThread(t->rchild);
}

void createInOrderThread(ThreadTree t) {
    if (t == NULL) {
        return;
    }

    pre = NULL;
    inOrderThread(t);

    // 最后 pre 一定和 cur 指向遍历序列的最终结点
    pre->rchild = NULL;  // 中序遍历的最后结点的右孩子一定是 null
    pre->rtag = 1;       // 处理遍历的最后一个结点
}

// 二叉树先序线索化
void preOrderThread(ThreadTree t) {
    if (t == NULL) {
        return;
    }

    visit(t);
    // 只有先序遍历会出现转圈的问题，因为在访问完根结点之后会访问左子树，
    // 访问左子结点之后已经与根结点建立了前驱的关系，如果不加判断则又会访问根结点
    // 同理，中序遍历和后续遍历则不会产生这种问题
    if (t->ltag == 0) {  // ⚠️lchild不是前驱线索
        preOrderThread(t->lchild);
    }
    preOrderThread(t->rchild);
}

void createPreOrderThread(ThreadTree t) {
    if (t == NULL) {
        return;
    }

    pre = NULL;
    preOrderThread(t);

    // 最后 pre 一定和 cur 指向遍历序列的最终结点
    if (pre->rchild ==
        NULL) {  // todo 先序遍历的最后结点的右孩子不一定是 null ？？？
        pre->rtag = 1;  // 处理遍历的最后一个结点
    }
}

// 二叉树后序线索化
void postOrderThread(ThreadTree t) {
    if (t == NULL) {
        return;
    }

    postOrderThread(t->lchild);
    postOrderThread(t->rchild);
    visit(t);
}

void createPeostOrderThread(ThreadTree t) {
    if (t == NULL) {
        return;
    }

    pre = NULL;
    postOrderThread(t);

    // 最后 pre 一定和 cur 指向遍历序列的最终结点
    if (pre->rchild == NULL) {  // 后序遍历的最后结点的右孩子不一定是 null
        pre->rtag = 1;  // 处理遍历的最后一个结点
    }
}