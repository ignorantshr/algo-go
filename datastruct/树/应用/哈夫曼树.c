#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#if !defined(_TREE_H_)
#define _TREE_H_
#include "../tree.h"
#endif  // _TREE_H_

#include "heap.c"

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
哈夫曼树不唯一，因此哈夫曼编码不唯一。哈夫曼编码可用于数据压缩

若没有一个编码是另一个编码的前缀，则称这样的编码为前缀编码
 */

// 给定一些结点，创建 哈夫曼 树
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
