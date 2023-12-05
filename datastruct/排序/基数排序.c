#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct LinkNode {
    int data;
    struct LinkNode* next;
} LinkNode, *LinkList;

typedef struct {             // 链式队列
    LinkNode *front, *rear;  // 队列的对头和队尾指针
} LinkQueue;

// todo
void radix_sort(int a[], int len) {}