#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define m 5

typedef struct BNode {
    int num;                 // 关键字数量
    int keys[m - 1];         // 关键字
    struct BNode* child[m];  // 孩子
} BNode, *BTree;

/* 下面是 ChatGPT 写的，然后修正了一下 */

// 在 B 树节点中插入关键字
void insertKey(BNode* node, int key) {
    int i = node->num - 1;

    // 找到合适的位置插入关键字
    while (i >= 0 && key < node->keys[i]) {
        node->keys[i + 1] = node->keys[i];
        i--;
    }

    node->keys[i + 1] = key;
    node->num++;
}

// 分裂节点的满子节点
void splitChild(BNode* parent, int index, BNode* child) {
    BNode* newNode = (BNode*)malloc(sizeof(BNode));
    // m=6,(m-1)/2=2
    // key   0,1, [2], 3,4
    // child 0,1, [2], 3,4,5
    // m=5,(m-1)/2=2
    // key   0,1, [2], 3
    // child 0,1, [2], 3,4

    // key[(m-1)/2] 飞升到父节点
    int s = (m - 1) / 2;
    newNode->num = m - 2 - s;  // key[(m-1)/2:]

    // 将子节点的后一半关键字复制到新节点中
    for (int i = 0; i < newNode->num; i++) {
        newNode->keys[i] = child->keys[i + s + 1];
    }

    // 如果子节点有孩子，也进行相应的复制
    if (child->child[0] != NULL) {
        for (int i = 0; i < newNode->num + 1; i++) {
            newNode->child[i] = child->child[i + s + 1];
        }
    }

    child->num = s;

    // 将父节点中的关键字和孩子后移
    for (int i = parent->num; i > index; i--) {
        parent->child[i + 1] = parent->child[i];
    }

    parent->child[index + 1] = newNode;

    for (int i = parent->num - 1; i >= index; i--) {
        parent->keys[i + 1] = parent->keys[i];
    }

    parent->keys[index] = child->keys[s];
    parent->num++;
}

// 在非满节点插入关键字
void insertNonFull(BNode* node, int key) {
    int i = node->num - 1;

    if (node->child[0] == NULL) {
        // 如果是终端节点，直接插入关键字
        insertKey(node, key);
    } else {
        // 找到合适的子节点继续插入
        while (i >= 0 && key < node->keys[i]) {
            i--;
        }

        i++;

        if (node->child[i]->num == m - 1) {
            // 如果子节点已满，进行分裂
            splitChild(node, i, node->child[i]);

            if (key > node->keys[i]) {
                i++;
            }
        }

        insertNonFull(node->child[i], key);
    }
}

// 在 B 树中插入关键字
void insertBTree(BTree* tree, int key) {
    if (*tree == NULL) {
        // 如果树为空，创建一个新的根节点
        BNode* newNode = (BNode*)malloc(sizeof(BNode));
        newNode->num = 1;
        newNode->keys[0] = key;

        for (int i = 0; i < m; i++) {
            newNode->child[i] = NULL;
        }

        *tree = newNode;
    } else {
        BNode* currentNode = *tree;

        if (currentNode->num == m - 1) {
            // 如果根节点已满，需要分裂出新的根
            BNode* newRoot = (BNode*)malloc(sizeof(BNode));
            newRoot->num = 0;
            newRoot->child[0] = currentNode;
            *tree = newRoot;

            splitChild(newRoot, 0, currentNode);
            insertNonFull(newRoot, key);
        } else {
            insertNonFull(currentNode, key);
        }
    }
}

// 打印 B 树按树形结构
void printBTree(BTree tree, int level) {
    if (tree == NULL) {
        return;
    }

    // for (int j = 0; j < level; j++) {
    //     printf("  ");
    // }
    // for (int i = 0; i < tree->num; i++) {
    //     printf("%d ", tree->keys[i]);
    // }
    // printf("\n");

    // for (int i = 0; i <= tree->num; i++) {
    //     printBTree(tree->child[i], level + 1);
    // }

    for (int i = 0; i < tree->num; i++) {
        if (tree->child[i] != NULL) {
            printBTree(tree->child[i], level + 1);
        }

        for (int j = 0; j < level; j++) {
            printf("  ");
        }

        printf("%d\n", tree->keys[i]);
    }
    if (tree->child[tree->num] != NULL) {
        printBTree(tree->child[tree->num], level + 1);
    }
}

/* 以下代码是 ChatGPT 编写，未验证 */

// 查找关键字在节点的索引位置
int searchIndex(BNode* node, int key) {
    int index = 0;
    while (index < node->num && key > node->keys[index]) {
        index++;
    }
    return index;
}

// 合并两个子节点
void mergeNodes(BNode* parent, int index, BNode* leftChild, BNode* rightChild) {
    leftChild->keys[leftChild->num] =
        parent->keys[index];  // 将父节点的关键字插入左子节点
    for (int i = 0; i < rightChild->num; i++) {
        leftChild->keys[leftChild->num + 1 + i] =
            rightChild->keys[i];  // 将右子节点的关键字插入左子节点
    }
    if (!leftChild->child[0]) {
        leftChild->num += rightChild->num + 1;  // 更新左子节点的关键字数量
    } else {
        leftChild->num += rightChild->num;  // 更新左子节点的关键字数量
    }
    // 将父节点的关键字和右子节点从父节点中删除
    for (int i = index; i < parent->num - 1; i++) {
        parent->keys[i] = parent->keys[i + 1];
        parent->child[i + 1] = parent->child[i + 2];
    }
    parent->child[parent->num--] = NULL;
    free(rightChild);  // 释放右子节点的内存空间
}

// 从左兄弟节点借关键字
void borrowFromPrevious(BNode* parent, int index, BNode* child) {
    BNode* leftSibling = parent->child[index - 1];
    // 将父节点的关键字插入到子节点的起始位置
    for (int i = child->num - 1; i >= 0; i--) {
        child->keys[i + 1] = child->keys[i];
    }
    if (!child->child[0]) {
        child->keys[0] = parent->keys[index - 1];
    } else {
        child->keys[0] = leftSibling->keys[leftSibling->num - 1];
    }
    // 如果子节点是内部节点，将左兄弟节点的最右孩子作为子节点的最左孩子
    if (child->child[0]) {
        for (int i = child->num; i >= 0; i--) {
            child->child[i + 1] = child->child[i];
        }
        child->child[0] = leftSibling->child[leftSibling->num];
    }
    child->num++;
    // 将左兄弟节点的最右关键字移动到父节点中
    if (!leftSibling->child[0]) {
        parent->keys[index - 1] = leftSibling->keys[leftSibling->num - 1];
    } else {
        parent->keys[index - 1] = leftSibling->keys[leftSibling->num - 2];
    }
    leftSibling->num--;
}

// 从右兄弟节点借关键字
void borrowFromNext(BNode* parent, int index, BNode* child) {
    BNode* rightSibling = parent->child[index + 1];
    // 将右兄弟节点的最左关键字插入到子节点的末尾
    if (!child->child[0]) {
        child->keys[child->num] = parent->keys[index];
    } else {
        child->keys[child->num] = rightSibling->keys[0];
    }
    if (child->child[0]) {
        child->child[child->num + 1] = rightSibling->child[0];
    }
    child->num++;
    // 将右兄弟节点的最左关键字移动到父节点中
    if (!rightSibling->child[0]) {
        parent->keys[index] = rightSibling->keys[0];
    } else {
        parent->keys[index] = rightSibling->keys[1];
    }
    // 将右兄弟节点中剩余的关键字左移一位
    for (int i = 1; i < rightSibling->num; i++) {
        rightSibling->keys[i - 1] = rightSibling->keys[i];
    }
    if (rightSibling->child[0]) {
        for (int i = 1; i <= rightSibling->num; i++) {
            rightSibling->child[i - 1] = rightSibling->child[i];
        }
    }
    rightSibling->num--;
}

// 删除关键字
void deleteKey(BNode* node, int key) {
    int index = searchIndex(node, key);

    // 如果关键字在当前节点中
    if (index < node->num && key == node->keys[index]) {
        // 如果当前节点是叶子节点
        if (!node->child[0]) {
            for (int i = index + 1; i < node->num; i++) {
                node->keys[i - 1] = node->keys[i];
            }
            node->num--;
        }
        // 如果当前节点是内部节点
        else {
            BNode* child = node->child[index];  // 关键字所在子树
            if (child->num >= m / 2) {
                int predecessor = child->keys[child->num - 1];  // 前驱关键字
                node->keys[index] = predecessor;  // 替换关键字
                deleteKey(child, predecessor);    // 递归删除前驱关键字
            } else {
                BNode* sibling;
                int borrowIndex;
                // 如果左兄弟节点有足够的关键字，从左兄弟节点借一个
                if (index > 0 && node->child[index - 1]->num >= m / 2) {
                    borrowFromPrevious(node, index, child);
                    borrowIndex = index - 1;
                    sibling = node->child[borrowIndex];
                }
                // 如果右兄弟节点有足够的关键字，从右兄弟节点借一个
                else if (index < node->num &&
                         node->child[index + 1]->num >= m / 2) {
                    borrowFromNext(node, index, child);
                    borrowIndex = index + 1;
                    sibling = node->child[borrowIndex];
                }
                // 否则合并关键字和子节点
                else {
                    if (index < node->num) {
                        mergeNodes(node, index, child, node->child[index + 1]);
                        sibling = node->child[index + 1];
                        borrowIndex = index;
                    } else {
                        mergeNodes(node, index - 1, node->child[index - 1],
                                   child);
                        sibling = node->child[index - 1];
                        borrowIndex = index - 1;
                    }
                }
                // 递归删除关键字
                deleteKey(sibling, key);
            }
        }
    }
    // 如果关键字不在当前节点中
    else {
        BNode* child = node->child[index];  // 包含关键字的子树
        if (child->num >= m / 2) {
            deleteKey(child, key);  // 递归在子树中删除关键字
        } else {
            // 如果左兄弟节点有足够的关键字，从左兄弟节点借一个
            if (index > 0 && node->child[index - 1]->num >= m / 2) {
                borrowFromPrevious(node, index, child);
            }
            // 如果右兄弟节点有足够的关键字，从右兄弟节点借一个
            else if (index < node->num &&
                     node->child[index + 1]->num >= m / 2) {
                borrowFromNext(node, index, child);
            }
            // 否则合并关键字和子节点
            else {
                if (index < node->num) {
                    mergeNodes(node, index, child, node->child[index + 1]);
                } else {
                    mergeNodes(node, index - 1, node->child[index - 1], child);
                }
            }
            // 递归删除关键字
            deleteKey(node->child[index], key);
        }
    }
}

// 主函数调用
void bTreeDelete(BTree tree, int key) {
    if (tree == NULL) {
        return;
    }

    // 如果根节点为空，直接返回
    if (tree->num == 0) {
        return;
    }

    deleteKey(tree, key);

    // 如果根节点没有关键字，将子节点作为新的根节点
    if (tree->num == 0) {
        BNode* oldRoot = tree;
        if (tree->child[0] != NULL) {
            tree = tree->child[0];
        } else {
            tree = NULL;
        }
        free(oldRoot);
    }
}

int main() {
    BTree tree = NULL;

    // 插入关键字示例
    for (int i = 0; i < 50; i++) {
        insertBTree(&tree, i);
    }

    // insertBTree(&tree, 10);
    // insertBTree(&tree, 20);
    // insertBTree(&tree, 5);
    // insertBTree(&tree, 30);
    // insertBTree(&tree, 15);
    // insertBTree(&tree, 11);
    // insertBTree(&tree, 8);
    // insertBTree(&tree, 9);

    printBTree(tree, 0);

    return 0;
}