#include "二叉树-链式存储.c"

void visit(ElemType e) {
    printf("%d ", e);
}

// 先序遍历，可以得出前缀表达式
void preOrderDfs(BiTree t) {
    if (t == NULL) {
        return;
    }

    visit(t->data);
    preOrderDfs(t->lchild);
    preOrderDfs(t->rchild);
}

// 中序遍历，可以得出中缀表达式（但是缺少界限符）
void inOrderDfs(BiTree t) {
    if (t == NULL) {
        return;
    }

    inOrderDfs(t->lchild);
    visit(t->data);
    inOrderDfs(t->rchild);
}

// 后序遍历，可以得出后缀表达式
void postOrderDfs(BiTree t) {
    if (t == NULL) {
        return;
    }

    postOrderDfs(t->lchild);
    postOrderDfs(t->rchild);
    visit(t->data);
}

// 求树高/深度
int treeDeepth(BiTree t) {
    if (t == NULL) {
        return 0;
    }

    int ld = treeDeepth(t->lchild);
    int rd = treeDeepth(t->rchild);
    return ld > rd ? ld + 1 : rd + 1;
}