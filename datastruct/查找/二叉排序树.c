#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// 二叉排序树结点
typedef struct BSTNode {
    int key;
    struct BSTNode *lchild, *rchild;
} BSTNode, *BSTree;

// 在二叉排序树中查找值为key的结点
// 最坏空间复杂度为 O(1), 递归版本的最坏空间复杂度为 O(h)
BSTNode* BST_Search(BSTree T, int key) {
    while (T != NULL && key != T->key) {  // 若树空或等于根结点值，则结束循环
        if (key < T->key) {
            T = T->lchild;  // 小于，则在左子树上查找
        } else {
            T = T->rchild;  // 大于，则在右子树上查找
        }
    }
    return T;
}

// 在二叉排序树中插入关键字为k的新结点（递归实现）
// 空间复杂度为 O(h)
int BST_Insert_DFS(BSTree T, int k) {
    if (T == NULL) {  // 原树为空，新插入的结点为根结点
        T = (BSTree)malloc(sizeof(BSTNode));
        (T)->key = k;
        (T)->lchild = (T)->rchild = NULL;
        return 1;                // 返回1，插入成功
    } else if (k == (T)->key) {  // 树中存在相同关键字的结点，插入失败
        return 0;
    } else if (k < (T)->key) {  // 插入到T的左子树
        return BST_Insert_DFS((T)->lchild, k);
    } else {  // 插入到T的右子树
        return BST_Insert_DFS((T)->rchild, k);
    }
}

// 在二叉排序树中插入关键字为k的新结点（循环实现）
// 空间复杂度为 O(1)
int BST_Insert(BSTree T, int k) {
    BSTNode* parent = NULL;
    BSTNode* cur = T;
    bool left;

    while (cur != NULL) {
        parent = cur;
        if (k < (cur)->key) {
            left = true;
            cur = (cur)->lchild;
        } else if (k > (cur)->key) {
            left = false;
            cur = (cur)->rchild;
        } else {
            // k == (cur)->key
            return 0;
        }
    }

    cur = (BSTree)malloc(sizeof(BSTNode));
    (cur)->key = k;
    (cur)->lchild = (cur)->rchild = NULL;
    if (parent != NULL) {
        if (left) {
            parent->lchild = cur;
        } else {
            parent->rchild = cur;
        }
    } else {
        T = cur;
    }
    return 1;  // 返回1，插入成功
}

// 按照str[]中的关键字序列建立二叉排序树
// 不同的关键字序列可能得到同款二叉排序树，也可能得到不同款二叉排序树
void Creat_BST(BSTree* T, int str[], int n) {
    int i = 0;
    if (*T == NULL) {
        *T = (BSTree)malloc(sizeof(BSTNode));
        (*T)->key = str[0];
        (*T)->lchild = (*T)->rchild = NULL;
        i++;
    }
    while (i < n) {
        BST_Insert(*T, str[i]);  // 依次将每个关键字插入到二叉排序树中
        i++;
    }
}

/* 二叉排序树的删除
先搜索找到目标结点

1.若被删除结点z是叶结点，则直接删除，不会破坏二叉排序树的性质

2.若结点z只有一棵左子树或右子树，则让z的子树成为z父结点的子树，替代z的位置

3.若结点z有左右两棵子树，则令z的直接后继（或直接前驱）替代z，然后从二叉排序树中删去这个直接后继（或直接前驱），这样就转换成了第一种或第二种情况
 */

BSTNode* BST_Remove(BSTree T, int k) {
    if (T == NULL) {
        return NULL;
    }

    if (T->key > k) {
        T->lchild = BST_Remove(T->lchild, k);
        return T;
    }

    if (T->key < k) {
        T->rchild = BST_Remove(T->rchild, k);
        return T;
    }

    // 被包含在只有一个子树的情况里面了
    // if (t->lchild == NULL && t->rchild == NULL) {
    //     free(t);
    //     return NULL;
    // }

    if (T->lchild == NULL) {
        BSTNode* tmp = T->rchild;
        free(T);
        return tmp;
    }

    if (T->rchild == NULL) {
        BSTNode* tmp = T->lchild;
        free(T);
        return tmp;
    }

    // 找前驱结点
    BSTNode* pre = T->lchild;
    while (pre->rchild != NULL) {
        pre = pre->rchild;
    }

    T->key = pre->key;
    T->lchild = BST_Remove(T->lchild, pre->key);
    return T;
}

void walk(BSTree t) {
    if (t == NULL) {
        return;
    }

    printf("%d ", t->key);
    walk(t->lchild);
    walk(t->rchild);
}

void testing() {
    int arr[10] = {5, 2, 8, 1, 3, 4, 6, 7, 9, 10};
    BSTree t = NULL;
    Creat_BST(&t, arr, 10);
    walk(t);
    printf("\n");

    t = BST_Remove(t, 8);
    walk(t);
    printf("\n");
    t = BST_Remove(t, 5);
    walk(t);
    printf("\n");
    t = BST_Remove(t, 7);
    walk(t);
    printf("\n");
    t = BST_Remove(t, 9);
    walk(t);
    printf("\n");
    t = BST_Remove(t, 4);
    walk(t);
    printf("\n");
}

int main(int argc, char const* argv[]) {
    testing();
    return 0;
}
