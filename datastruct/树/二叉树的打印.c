#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "tree.h"

// 横向展开打印
void printTreeHorizontal(BiTree root, int level) {
    // if (root == NULL) {
    //     return;
    // }

    if (root != NULL) {
        printTreeHorizontal(root->rchild, level + 1);
    }

    for (int i = 0; i < level; i++) {
        printf("   ");
    }

    if (root != NULL) {
        printf("%d\n", root->data);
    } else {
        printf("null\n");
    }

    if (root != NULL) {
        printTreeHorizontal(root->lchild, level + 1);
    }
}

void writeToFile(const char* filename, const char* content) {
    FILE* file = fopen(
        filename,
        "a");  // 打开文件，以写入模式打开（如果文件不存在则创建，如果存在则清空内容）
    if (file == NULL) {
        printf("无法打开文件 %s\n", filename);
        return;
    }

    fprintf(file, "%s", content);  // 将内容写入文件

    fclose(file);  // 关闭文件
}

int lvls[10];       // 每层已遍历结点的数量
char* lines[10];    // 每层的结点信息
int height;         // 树高
int pos[10][100];   // i 层 j 个结点的位置
const int gap = 3;  // 结点直接的间隔
char null[gap + 1];

int treeHeight(BiTree t) {
    if (t == NULL) {
        return 0;
    }

    return max(treeHeight(t->lchild), treeHeight(t->rchild)) + 1;
}

int bottomNum(BiTree node, int level) {
    int n = pow(2, height - 1);  // 底层结点总数
    int splitn = pow(2, level);  // 当前层结点总数
    int per = n / splitn;        // 划分
    int i = lvls[level];         // 当前层第几个结点
    return i * per +
           per / 2;  // 计算底层结点在自己左边的结点数 i * per + per/2, i from 0
}

int brotherPos(int level) {
    if (lvls[level] == 0) {
        return 0;
    }
    return pos[level][lvls[level] - 1];
}

// 纵向展开打印
void printTreeVertical(BiTree root, int level, int h) {
    int n = powl(2, level);
    char pre[200];
    pre[0] = '\0';
    int bnn = bottomNum(root, level);
    bnn = level == height - 1 ? bnn * 2 : bnn * 2 - 1;
    int preCharNum = gap * bnn;

    // 找同层左兄弟结点的尾巴的位置
    int bropos = brotherPos(level);
    preCharNum -= bropos;
    // 防止过长异常
    if (preCharNum <= 0 && lvls[level] != 0) {
        preCharNum = gap / 2;
    }

    for (int i = 0; i < preCharNum; i++) {
        strcat(pre, " ");
    }

    char tmp[300];
    int l;
    if (root == NULL) {
        l = sprintf(tmp, "%s%s", pre, null);
    } else {
        l = sprintf(tmp, "%s%d", pre, root->data);
    }
    tmp[l] = '\0';
    pos[level][lvls[level]] = bropos + l;

    // printf("[%d] line[%d] space:%d, l:%d\n", root == NULL ? -1 : root->data,
    //        level, preCharNum, l);

    if (lines[level] == NULL) {
        lines[level] = (char*)malloc(preCharNum + l);
        strcpy(lines[level], tmp);
    } else {
        size_t newSize =
            strlen(lines[level]) + strlen(tmp) + 1;  // 计算新的内存大小
        char* newBuffer =
            (char*)realloc(lines[level], newSize);  // 重新分配内存
        if (newBuffer != NULL) {
            lines[level] = newBuffer;
            strcat(lines[level], tmp);
        } else {
            // 内存分配失败的处理逻辑
            printf("realloc error\n");
            exit(1);
        }
    }

    lvls[level]++;
    if (lvls[level] == n) {
        // writeToFile("tree.log", strcat(lines[level], "\n"));
        printf("%s\n", lines[level]);
    }

    if (h == height) {
        return;
    }
    if (root == NULL) {
        printTreeVertical(NULL, level + 1, h + 1);
        printTreeVertical(NULL, level + 1, h + 1);
    } else {
        printTreeVertical(root->lchild, level + 1, h + 1);
        printTreeVertical(root->rchild, level + 1, h + 1);
    }
}

void initPrint(BiTree root) {
    height = treeHeight(root);

    for (int i = 0; i < gap; i++) {
        null[i] = 'n';
    }
    null[gap] = '\0';
}
