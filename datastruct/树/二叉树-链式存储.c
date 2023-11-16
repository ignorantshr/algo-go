#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "tree.h"

// 若有n个结点，则有2n个指针域，而除了根结点每个节点头上都会连一个指针，故n个结点的二叉链表共有n+1个空链域。

void destroy(BiTree t) {
    if (t == NULL) {
        return;
    }

    if (t->lchild != NULL) {
        destroy(t->lchild);
    }
    if (t->rchild != NULL) {
        destroy(t->rchild);
    }
    free(t);
}

void testing() {
    BiTree root = (BiTree)malloc(sizeof(BiTNode));
    root->data = 1;
    root->lchild = NULL;
    root->rchild = NULL;

    BiTNode* ln = (BiTNode*)malloc(sizeof(BiTNode));
    ln->data = 2;
    ln->lchild = NULL;
    ln->rchild = NULL;
    root->lchild = ln;

    destroy(root);
}
