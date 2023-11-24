#include <stdio.h>
#include <stdlib.h>

#define max(x, y) ((x) > (y) ? (x) : (y))

typedef struct AVLNode {
    int key;      // 数据域
    int balance;  // 平衡因子
    // int height;   // 节点高度
    struct AVLNode *lchild, *rchild;
} AVLNode, *AVLTree;

// 获取节点高度
// int getHeight2(AVLNode* node) {
//     if (node == NULL) {
//         return 0;
//     }
//     return node->height;
// }

// 更新节点高度
// void updateHeight(AVLNode* node) {
//     if (node == NULL) {
//         return;
//     }
//     int leftHeight = getHeight2(node->lchild);
//     int rightHeight = getHeight2(node->rchild);
//     node->height = (leftHeight > rightHeight ? leftHeight : rightHeight) + 1;
// }

// 获取节点的高度
int getHeight(AVLTree root) {
    if (root == NULL) {
        return 0;
    }
    // return 1 + max(getHeight2(root->lchild), getHeight2(root->rchild));
    return 1 + max(getHeight(root->lchild), getHeight(root->rchild));
}

// 获取节点的平衡因子
int getBalance(AVLTree node) {
    if (node == NULL) {
        return 0;
    }
    return getHeight(node->lchild) - getHeight(node->rchild);
}

// 左旋操作
AVLTree leftRotate(AVLTree root) {
    /*
   a            b
    \          / \
     b    ->  a   d
    / \        \
   c   d        c

 */
    AVLTree newRoot = root->rchild;
    root->rchild = newRoot->lchild;
    newRoot->lchild = root;
    // updateHeight(root);
    // updateHeight(newRoot);
    return newRoot;
}

// 右旋操作
AVLTree rightRotate(AVLTree root) {
    /*
       a        b
      /        / \
     b    ->  c   a
    / \          /
   c   d        d

     */
    AVLTree newRoot = root->lchild;
    root->lchild = newRoot->rchild;
    newRoot->rchild = root;

    // updateHeight(root);
    // updateHeight(newRoot);
    return newRoot;
}

/* 插入节点
为了便于记忆，可以总结出
只有左孩子才能右上旋，只有右孩子才能左上旋，每一次旋转都能导致它和它的父结点父子关系互换的规律。
在LR中处理A的左孩子的右孩子，它首先是一个右孩子，所以第一步只能让它左上旋替代A以前的左孩子，然后作为左孩子它只能右旋，RL也遵循相同的规律
插入操作导致“最小不平衡子树”高度+1，经过调整后高度恢复，所以在插入操作中只要将最小不平衡树调整平衡则其他祖先结点的平衡因子都会恢复
 */
AVLTree insertNode(AVLTree root, int key) {
    if (root == NULL) {
        AVLTree newNode = (AVLTree)malloc(sizeof(AVLNode));
        newNode->key = key;
        newNode->balance = 0;
        newNode->lchild = newNode->rchild = NULL;
        // updateHeight(newNode);
        return newNode;
    }

    if (key < root->key) {
        root->lchild = insertNode(root->lchild, key);
    } else {
        root->rchild = insertNode(root->rchild, key);
    }

    // updateHeight(root);

    // 更新平衡因子
    root->balance = getBalance(root);

    // 平衡调整
    if (root->balance > 1) {
        // 左边高
        if (key < root->lchild->key) {
            // LL 在A的左孩子（L）的左子树（L）上插入新结点
            root = rightRotate(root);
        } else {
            // LR 在A的左孩子（L）的右子树（R）上插入新结点
            // 先左孩子左旋，后根结点右旋
            root->lchild = leftRotate(root->lchild);
            root = rightRotate(root);
        }
    } else if (root->balance < -1) {
        // 右边高
        if (key > root->rchild->key) {
            // RR
            root = leftRotate(root);
        } else {
            // RL
            root->rchild = rightRotate(root->rchild);
            root = leftRotate(root);
        }
    }

    return root;
}

// 找到后继节点
AVLTree findPostNode(AVLTree root) {
    if (root == NULL) {
        return NULL;
    }
    while (root->lchild != NULL) {
        root = root->lchild;
    }
    return root;
}

// 删除节点
// 和二叉排序树一样，但是多了平衡性调整的步骤
// 删除结点之后可能会触发多次平衡调整
AVLTree deleteNode(AVLTree root, int key) {
    if (root == NULL) {
        return NULL;
    }

    if (key < root->key) {
        root->lchild = deleteNode(root->lchild, key);
    } else if (key > root->key) {
        root->rchild = deleteNode(root->rchild, key);
    } else {
        // 找到了要删除的节点

        if (root->lchild == NULL && root->rchild == NULL) {
            // 情况1：要删除的节点为叶子节点
            free(root);
            return NULL;
        } else if (root->lchild == NULL) {
            // 情况2：要删除的节点只有右子节点
            AVLTree temp = root->rchild;
            free(root);
            return temp;
        } else if (root->rchild == NULL) {
            // 情况3：要删除的节点只有左子节点
            AVLTree temp = root->lchild;
            free(root);
            return temp;
        } else {
            // 情况4：要删除的节点有两个子节点

            // 找到右子树中的最小值节点
            AVLTree postNode = findPostNode(root->rchild);

            // 将最小值节点的值赋给要删除的节点，并在右子树中删除最小值节点
            root->key = postNode->key;
            root->rchild = deleteNode(root->rchild, postNode->key);
        }
    }

    // updateHeight(root);

    // 更新平衡因子
    root->balance = getBalance(root);

    // 平衡调整
    if (root->balance > 1) {
        // 左子树高
        if (getBalance(root->lchild) >= 0) {
            // LL
            root = rightRotate(root);
        } else {
            // LR
            root->lchild = leftRotate(root->lchild);
            root = rightRotate(root);
        }
    } else if (root->balance < -1) {
        if (getBalance(root->rchild) <= 0) {
            // RR
            root = leftRotate(root);
        } else {
            // RL
            root->rchild = rightRotate(root->rchild);
            root = leftRotate(root);
        }
    }

    return root;
}

// 中序遍历打印树节点
void inorderTraversal(AVLTree root) {
    if (root == NULL) {
        return;
    }
    inorderTraversal(root->lchild);
    printf("%d ", root->key);
    inorderTraversal(root->rchild);
}

void printTreeShape(AVLTree root, int level) {
    if (root == NULL) {
        return;
    }

    printTreeShape(root->rchild, level + 1);

    for (int i = 0; i < level; i++) {
        printf("   ");
    }

    // printf("%d[%d]\n", root->key, root->height);
    printf("%d[%d]\n", root->key, root->balance);

    printTreeShape(root->lchild, level + 1);
}

int main() {
    AVLTree root = NULL;

    // root = insertNode(root, 5);
    // root = insertNode(root, 3);
    // root = insertNode(root, 7);
    // root = insertNode(root, 1);
    // root = insertNode(root, 4);
    // root = insertNode(root, 2);
    root = insertNode(root, 1);
    root = insertNode(root, 2);
    root = insertNode(root, 3);
    root = insertNode(root, 4);
    root = insertNode(root, 5);
    root = insertNode(root, 6);
    root = insertNode(root, 7);
    root = insertNode(root, 8);

    printf("中序遍历结果：");
    inorderTraversal(root);
    printf("\n");
    printTreeShape(root, 0);

    return 0;
}
