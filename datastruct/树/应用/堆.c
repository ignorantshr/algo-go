#include <stdio.h>

#if !defined(_TREE_H_)
#define _TREE_H_
#include "../tree.h"
#endif  // _TREE_H_

#define MaxSize 30

typedef struct {
    BiTNode* data[MaxSize];  // 用数组存储堆的数据
    int size;                // 堆的当前大小
} MinHeap;

// 初始化一个空的小顶堆
void InitMinHeap(MinHeap* heap) {
    heap->size = 0;
}

int SizeHeap(MinHeap* heap) {
    if (heap == NULL) {
        return -1;
    }

    return heap->size;
}

// 插入一个新元素
void PushHeap(MinHeap* heap, BiTNode* node) {
    if (heap->size > MaxSize) {
        printf("Error: heap is full\n");
        return;
    }

    int i = ++(heap->size);  // 插入新元素后堆的大小加1
    for (; heap->data[i / 2]->data > node->data && i > 1; i /= 2) {
        heap->data[i] = heap->data[i / 2];  // 上浮操作
    }
    heap->data[i] = node;  // 将新元素插入到正确的位置
}

// 从指定位置开始进行下沉操作
void percolateDown(MinHeap* heap, int position) {
    int child;
    BiTNode* tmp = heap->data[position];

    for (; position * 2 <= heap->size; position = child) {
        child = position * 2;
        if (child != heap->size &&
            heap->data[child + 1]->data < heap->data[child]->data) {
            child++;
        }
        if (heap->data[child]->data < tmp->data) {
            heap->data[position] = heap->data[child];  // 下沉操作
        } else {
            break;
        }
    }
    heap->data[position] = tmp;
}

// 删除堆顶元素
BiTNode* PopHeap(MinHeap* heap) {
    BiTNode* minItem;
    if (heap->size < 1) {
        printf("Error: heap is empty\n");
        return NULL;
    }
    minItem = heap->data[1];
    heap->data[1] = heap->data[heap->size--];
    percolateDown(heap, 1);
    return minItem;
}
