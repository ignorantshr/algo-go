#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ElemType int

// 索引表
typedef struct {
    ElemType maxValue;
    int low, high;
} Index;

// 顺序表存储实际元素
ElemType List[100];

/* 若用折半查找查索引
若索引表中不包含目标关键字，则折半查找索引表最终停在 low > high，要在 low
所指分块中查找。
*/