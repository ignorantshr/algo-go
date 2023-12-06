#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#if !defined(_TREE_H_)
#define _TREE_H_
#include "../tree.h"
#endif  // _TREE_H_

#include "堆.c"

/*
结点的权：有某种现实含义的数值（如：表示结点的重要性等）

结点的带权路径长度：从树的根到该结点的路径长度（经过的边数）与该结点上权值的乘积

树的带权路径长度：树中所有叶结点的带权路径长度之和（WPL）

在含有n个带权叶结点的二叉树中，其中带权路径长度（WPL）最小的二叉树称为哈夫曼树，也称最优二叉树
 */

/*哈夫曼 树的性质
1.每个初始结点最终都成为叶结点，且权值越小的结点到根结点的路径长度最大

2.哈夫曼树的结点总数为2n−1
    原有n个结点，结合n-1次，每次结合产生一个新的结点，故哈夫曼树的结点总数为2n−1

3.哈夫曼树中不存在度为1的结点

4.哈夫曼树并不唯一，但WPL必然相同且为最优
 */

/* 哈夫曼编码，必须得有哈夫曼树作为对应

用二进制来表示字符，有下面两种方案：
- 固定长度编码——每个字符用相等长度的二进制位表示
- 可变长度编码——允许对不同字符用不等长的二进制位表示

由哈夫曼树得到哈夫曼编码——字符集中的每个字符作为一个叶子结点，各个字符出现的频度作为结点的权值，根据之前介绍的方法构造哈夫曼树。
哈夫曼树不唯一，因此哈夫曼编码不唯一。哈夫曼编码可用于数据压缩，哈夫曼编码中，一个编码不能是任何其它编码的前缀。

若没有一个编码是另一个编码的前缀，则称这样的编码为前缀编码
 */

/* 哈夫曼树的构造
给定n个权值分别为 w_1,w_2,...,w_n 的结点，构造哈夫曼树的算法描述如下：

1.将这n个结点分别作为n棵仅含一个结点的二叉树，构成森林F
2.构造一个新结点，从F中选取两棵根结点权值最小的树作为新结点的左、右子树，并且将新结点的权值置为左、右子树上根结点的权值之和
3.从F中删除刚才选出的两棵树，同时将新得到的树加入F中
4.重复步骤2和3，直至F中只剩下一棵树为止
 */
BiTree createHuffmanTree(int vals[], int len) {
    // 可以用优先队列或堆来寻找最小的两个结点
    MinHeap forests;
    InitMinHeap(&forests);
    for (int i = 0; i < len; i++) {
        BiTNode* node = (BiTNode*)malloc(sizeof(BiTNode));
        node->data = vals[i];
        node->lchild = NULL;
        node->rchild = NULL;
        PushHeap(&forests, node);
    }

    while (SizeHeap(&forests) > 1) {
        BiTNode* first = PopHeap(&forests);
        BiTNode* second = PopHeap(&forests);
        BiTNode* node = (BiTNode*)malloc(sizeof(BiTNode));
        node->data = first->data + second->data;
        node->lchild = first;
        node->rchild = second;
        PushHeap(&forests, node);
    }

    return PopHeap(&forests);
}

void scan(BiTree t) {
    if (t == NULL) {
        return;
    }

    if (t->lchild == NULL && t->rchild == NULL) {
        printf("[%d] ", t->data);
    } else {
        printf("%d ", t->data);
    }
    scan(t->lchild);
    scan(t->rchild);
}

void testing() {
    const int len = 5;
    int arr[len] = {1, 2, 3, 4, 5};
    BiTree t = createHuffmanTree(arr, len);
    scan(t);
}

int main(int argc, char const* argv[]) {
    testing();
    return 0;
}
