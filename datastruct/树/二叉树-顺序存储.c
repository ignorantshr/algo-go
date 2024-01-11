#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "tree.h"
/* ⚠️重点
度为 0 的结点数 = 度为 2 的结点数 + 1
结点总数 = 总度数 + 1

i的父节点——[i/2]
i所在的层次—— [log_2(n+1)]（[]表示向上取整） 或 [log_2n]+1
（[]表示向下取整）

完全二叉树中共有 n个结点，则
判断 i 是否有左孩子？——2i ≤ n ?
判断 i 是否有右孩子？——2i + 1 ≤ n?
判断 i 是否是叶子结点？——i > [n/2] ?
*/

/* 给出完全二叉树的结点数 n 或结点编号 i，计算层高：

h 层高的满二叉树有 2^(h)-1 个结点
h-1 层高的满二叉树有 2^(h-1)-1 个结点

2^(h-1)-1 < i <=2^(h)-1
2^(h-1) < i+1 <=2^(h)
h >= log(i+1)
h < log(i+1)+1
h = log(i+1) 向上取整

h 层高的完全二叉树最多有 2^(h)-1 个结点，则
h-1 层高的满二叉树有 2^(h-1)-1 个结点
h 层高的完全二叉树最少有 2^(h-1) 个结点
2^(h-1)-1+1 <= i <= 2^(h)-1
2^(h-1) <= i < 2^(h)
h > log(i)
h <= log(i)+1
h = log(i)+1 向下取整
 */

#define MaxSize 10

typedef struct node {
    ElemType data;
    bool isFull;  // 指示是否有值
} node;

int main(int argc, char const* argv[]) {
    // 如果是非完全二叉树，那么这样存储会极其浪费空间
    node tree[MaxSize];
    // tree[0] 不用，从 1 开使存储
    int i = 1;
    tree[i].data = 1;
    tree[i].isFull = true;
    // left node
    node ln = {2, true};
    tree[i * 2] = ln;
    // right node
    tree[i * 2 + 1].data = 3;
    tree[i * 2 + 1].isFull = true;
    return 0;
}
