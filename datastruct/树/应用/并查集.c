#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MaxSize 5

/* 并查集使用双亲结点表示比较好实现
 */

char ufSets[MaxSize];  // 集合元素数组
// parent[] 父结点的索引数组

// 初始化并查集
void Initial(int parent[], int len) {
    for (int i = 0; i < len; i++) {
        parent[i] = -1;
    }
}

// Find “查”操作，找索引x所属集合（返回x所属根结点索引），最坏时间复杂度：O(n)
int Find(int parent[], int x) {
    while (parent[x] >= 0) {
        x = parent[x];
    }
    return x;
}

// 如果指定的两个元素不是根节点，要合并这两个元素从属的集合，需要先“查”确定两个元素各自的根节点，然后再对两个根节点进行“并”
// Union “并”操作，将两个集合合并为一个,时间复杂度：O(1)
// 将 n 个独立的元素通过多次 union 合并为一个集合——O(n^2)
void Union(int parent[], int root1, int root2) {
    if (root1 == root2) {
        return;
    }

    // 合并
    parent[root2] = root1;
}

void printidx(int parent[]) {
    printf("--------------\n");
    for (int i = 0; i < MaxSize; i++) {
        printf("%c\t", ufSets[i]);
    }
    printf("\n");
    for (int i = 0; i < MaxSize; i++) {
        printf("%d\t", parent[i]);
    }
    printf("\n");
}

/* Union操作的优化

优化思路：在每次Union操作构建树的时候，尽可能让树不长高

1.用根节点的 绝对值 表示树的结点总数
2.Union操作，结点总数小的树是小树，让小树合并到大树

Find的最坏时间复杂度从 O(n) 变为 O(logn)
该方法构造的树高不超过 [logn]+1,该结论可以用数学归纳法证明

// 将 n 个独立的元素通过多次 union 合并为一个集合——O(nlogn)
 */
void Union2(int parent[], int root1, int root2) {
    if (root1 == root2) {
        return;
    }

    if (parent[root1] > parent[root2]) {  // -2 > -5, root1 的节点更少
        parent[root2] += parent[root1];
        parent[root1] = root2;  // 小树合并到大树
    } else {
        parent[root1] += parent[root2];
        parent[root2] = root1;  // 小树合并到大树
    }
}

/* 并查集的进一步优化
Find操作的优化（压缩路径）
将查找路径上的各个结点全部挂到根节点A下面，让树尽量矮，这样后面再查找时这条路径上的结点可直接找到根结点
，可使树的高度不超过O(α(n))，最坏时间复杂度变为 O(α(n))

// 将 n 个独立的元素通过多次 union 合并为一个集合——O(nα(n))
 */
int Find2(int parent[], int x) {
    int root = x;
    while (parent[root] >= 0) {
        root = parent[root];
    }

    while (x != root) {
        int tmp = parent[x];
        parent[x] = root;  // 直接链接到根结点
        x = tmp;
    }

    return root;
}

void testing() {
    int elements[MaxSize] = {'a', 'b', 'c', 'd', 'e'};
    for (int i = 0; i < MaxSize; i++) {
        ufSets[i] = elements[i];
    }

    int parent[MaxSize];
    Initial(parent, MaxSize);
    printidx(parent);

    Union(parent, 0, 1);
    Union(parent, 0, 2);

    Union(parent, 3, 4);
    printidx(parent);

    printf("root: %d\n", Find(parent, 2));
    printf("root: %d\n", Find(parent, 3));

    Union(parent, 0, 3);
    printidx(parent);
    printf("root: %d\n", Find(parent, 4));
}

int main(int argc, char const* argv[]) {
    testing();
    return 0;
}
