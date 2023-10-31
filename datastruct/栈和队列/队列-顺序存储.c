#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "queue.h"

// 存储空间 判满/判空 策略：
// 1. 牺牲一个存储单元
// 2. 使用 size 变量记录长度
// 3. 使用 tag 标记最近一次的操作是 0:出队, 1:入队

#define MaxSize 4  // 定义队列中元素的最大个数
typedef struct {
    ElemType data[MaxSize];  // 用静态数组存放队列元素
    int front, rear;         // 队头指针和队头指针
    int size;
    int tag;
} SqQueue;

// ------------1. 牺牲一个存储单元---------------
// 初始化队列
void InitQueue1(SqQueue* q) {
    q->front = 0;  // 队头指针指向队头元素的位置
    q->rear = 0;   // 队尾指针指向下一个存放的位置
    // q->rear = MaxSize - 1;  // 队尾指针指向队尾元素的写法
}

bool full1(SqQueue q) {
    // r 不存放元素
    // . . r f . .
    return q.front == (q.rear + 1) % MaxSize;
}

bool empty1(SqQueue q) {
    return q.rear == q.front;
}

bool enQueue1(SqQueue* q, ElemType e) {
    if (full1(*q)) {
        return false;
    }

    q->data[q->rear] = e;
    q->rear = (q->rear + 1) % MaxSize;
    return true;
}

bool deQueue1(SqQueue* q, ElemType* e) {
    if (empty1(*q)) {
        return false;
    }

    *e = q->data[q->front];
    q->front = (q->front + 1) % MaxSize;
    return true;
}

bool front1(SqQueue* q, ElemType* e) {
    if (empty1(*q)) {
        return false;
    }

    *e = q->data[q->front];
    return true;
}

int len1(SqQueue q) {
    return (q.rear + MaxSize - q.front) % MaxSize;
}

void printqueue1(SqQueue q, char* s) {
    s[0] = '\0';
    for (int i = q.front; i != q.rear; i = (i + 1) % MaxSize) {
        char tmp[10];
        sprintf(tmp, "%d-", q.data[i]);
        strcat(s, tmp);
    }
    printf("%s\n", s);
}

// ------------1. 牺牲一个存储单元---------------

// ------------2. 使用 size 变量记录长度---------------
void InitQueue2(SqQueue* q) {
    q->size = 0;
    q->front = 0;  // 队头指针指向队头元素的位置
    q->rear = 0;   // 队尾指针指向下一个存放的位置
}

bool full2(SqQueue q) {
    return q.size == MaxSize;
}

bool empty2(SqQueue q) {
    return q.size == 0;
}

bool enQueue2(SqQueue* q, ElemType e) {
    if (full2(*q)) {
        return false;
    }

    q->size++;
    q->data[q->rear] = e;
    q->rear = (q->rear + 1) % MaxSize;
    return true;
}

bool deQueue2(SqQueue* q, ElemType* e) {
    if (empty2(*q)) {
        return false;
    }

    q->size--;
    *e = q->data[q->front];
    q->front = (q->front + 1) % MaxSize;
    return true;
}

bool front2(SqQueue* q, ElemType* e) {
    if (empty2(*q)) {
        return false;
    }

    *e = q->data[q->front];
    return true;
}

void printqueue2(SqQueue q, char* s) {
    s[0] = '\0';
    for (int i = q.front; i < q.front + q.size; i++) {
        char tmp[10];
        sprintf(tmp, "%d-", q.data[(i) % MaxSize]);
        strcat(s, tmp);
    }
    printf("%s\n", s);
}

// ------------2. 使用 size 变量记录长度---------------

// ------------3. 使用 tag 标记最近一次的操作是 出队/入队---------------
void InitQueue3(SqQueue* q) {
    q->front = 0;  // 队头指针指向队头元素的位置
    q->rear = 0;   // 队尾指针指向下一个存放的位置
    q->tag = 0;
}

bool full3(SqQueue q) {
    return q.tag == 1 && q.front == q.rear;
}

bool empty3(SqQueue q) {
    return q.tag == 0 && q.rear == q.front;
}

bool enQueue3(SqQueue* q, ElemType e) {
    if (full3(*q)) {
        return false;
    }

    q->tag = 1;
    q->data[q->rear] = e;
    q->rear = (q->rear + 1) % MaxSize;
    return true;
}

bool deQueue3(SqQueue* q, ElemType* e) {
    if (empty3(*q)) {
        return false;
    }

    q->tag = 0;
    *e = q->data[q->front];
    q->front = (q->front + 1) % MaxSize;
    return true;
}

bool front3(SqQueue* q, ElemType* e) {
    if (empty3(*q)) {
        return false;
    }

    *e = q->data[q->front];
    return true;
}

void printqueue3(SqQueue q, char* s) {
    s[0] = '\0';
    if (full3(q)) {
        int i = q.front;
        do {
            char tmp[10];
            sprintf(tmp, "%d-", q.data[i]);
            strcat(s, tmp);
            i = (i + 1) % MaxSize;
        } while (i != q.rear);
    } else {
        for (int i = q.front; i != q.rear; i = (i + 1) % MaxSize) {
            char tmp[10];
            sprintf(tmp, "%d-", q.data[i]);
            strcat(s, tmp);
        }
    }
    printf("%s\n", s);
}
// ------------3. 使用 tag 标记最近一次的操作是 出队/入队---------------

typedef struct queue {
    void (*InitQueue)(SqQueue* q);
    bool (*full)(SqQueue q);
    bool (*empty)(SqQueue q);
    bool (*enQueue)(SqQueue* q, ElemType e);
    bool (*deQueue)(SqQueue* q, ElemType* e);
    bool (*front)(SqQueue* q, ElemType* e);
    void (*printqueue)(SqQueue q, char* s);
} queue;

void testing(queue qq) {
    SqQueue q;
    int e;
    char s[100];
    qq.InitQueue(&q);

    qq.enQueue(&q, 1);
    qq.printqueue(q, s);
    assert(!strcmp(s, "1-"));

    qq.enQueue(&q, 2);
    qq.printqueue(q, s);
    assert(!strcmp(s, "1-2-"));

    qq.deQueue(&q, &e);
    assert(e == 1);
    qq.printqueue(q, s);
    assert(!strcmp(s, "2-"));

    qq.enQueue(&q, 3);
    qq.printqueue(q, s);
    assert(!strcmp(s, "2-3-"));

    qq.front(&q, &e);
    assert(e == 2);

    qq.deQueue(&q, &e);
    qq.deQueue(&q, &e);
    assert(e == 3);
    assert(qq.empty(q));

    assert(!qq.deQueue(&q, &e));

    qq.enQueue(&q, 3);
    qq.printqueue(q, s);
    assert(!strcmp(s, "3-"));

    qq.enQueue(&q, 2);
    qq.printqueue(q, s);
    assert(!strcmp(s, "3-2-"));

    qq.enQueue(&q, 1);
    qq.printqueue(q, s);
    assert(!strcmp(s, "3-2-1-"));

    assert(!qq.enQueue(&q, 1));
}

void test1() {
    // only use MaxSize-1
    // #define MaxSize 4

    queue qq;

    qq.InitQueue = InitQueue1;
    qq.full = full1;
    qq.empty = empty1;
    qq.enQueue = enQueue1;
    qq.deQueue = deQueue1;
    qq.front = front1;
    qq.printqueue = printqueue1;

    testing(qq);
}

void test2() {
    // MaxSize = 3
    queue qq;

    qq.InitQueue = InitQueue2;
    qq.full = full2;
    qq.empty = empty2;
    qq.enQueue = enQueue2;
    qq.deQueue = deQueue2;
    qq.front = front2;
    qq.printqueue = printqueue2;

    testing(qq);
}

void test3() {
    // MaxSize = 3
    queue qq;

    qq.InitQueue = InitQueue3;
    qq.full = full3;
    qq.empty = empty3;
    qq.enQueue = enQueue3;
    qq.deQueue = deQueue3;
    qq.front = front3;
    qq.printqueue = printqueue3;

    testing(qq);
}

int main(int argc, char const* argv[]) {
    test1();
    // test2();
    // test3();
    return 0;
}
