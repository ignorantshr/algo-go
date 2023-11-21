#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ElemType int

// 以顺序表为例，查找表是单链表或者双链表当然也行
typedef struct {     // 查找表的数据结构(顺序表)
    ElemType* elem;  // 动态数组基址
    int TableLen;    // 表的长度
} SSTable;

// 顺序查找
int Search_Seq(SSTable ST, ElemType key) {
    int i;
    for (i = 0; i < ST.TableLen && ST.elem[i] != key; ++i) {
    }
    // 查找成功，则返回元素下标;查找失败，则返回-1
    return i == ST.TableLen ? -1 : i;
}

// 顺序查找，哨兵模式
// 这种写法在每一轮for循环的时候，只需要判断当前指向的元素与要查找的关键字是否相等而无需判断是否越界，效率更高一点点
int Search_Seq(SSTable ST, ElemType key) {
    ST.elem[0] = key;  // 哨兵
    int i;
    for (i = ST.TableLen; ST.elem[i] != key; --i) {
    }
    // 从后往前找
    return i;  // 查找成功，则返回元素下标;查找失败，则返回0
}
