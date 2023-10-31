#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "queue.h"

typedef struct Node {
    ElemType data;
    struct Node* next;
} Node;

typedef struct LinkQueue {
    Node *front, *rear;
} LinkQueue;

void destory(LinkQueue* q) {
    for (Node* n = 0; n != NULL;) {
        Node* tmp = n->next;
        free(n);
        n = tmp;
    }
}

// 不带头结点
void InitQueue1(LinkQueue* q) {
    q->front = NULL;
    q->rear = NULL;
    return;
}

bool empty1(LinkQueue q) {
    return q.front == NULL;
}

bool enQueue1(LinkQueue* q, ElemType e) {
    Node* tmp = (Node*)malloc(sizeof(Node));
    tmp->data = e;
    tmp->next = NULL;
    if (empty1(*q)) {
        q->front = tmp;
    } else {
        q->rear->next = tmp;
    }
    q->rear = tmp;
    return true;
}

bool deQueue1(LinkQueue* q, ElemType* e) {
    if (empty1(*q)) {
        return false;
    }

    Node* tmp = q->front;
    *e = tmp->data;
    q->front = tmp->next;
    if (q->front == NULL) {
        q->rear = NULL;
    }
    free(tmp);
    return true;
}

bool front1(LinkQueue* q, ElemType* e) {
    if (empty1(*q)) {
        return false;
    }

    *e = q->front->data;
    return true;
}

void printqueue1(LinkQueue q, char* s) {
    s[0] = '\0';
    for (Node* t = q.front; t != NULL; t = t->next) {
        char tmp[10];
        sprintf(tmp, "%d-", t->data);
        strcat(s, tmp);
    }
    printf("%s\n", s);
}

// 带头结点
void InitQueue2(LinkQueue* q) {
    Node* head = (Node*)malloc(sizeof(Node));
    head->next = NULL;
    q->front = head;
    q->rear = head;
}

bool empty2(LinkQueue q) {
    return q.front->next == NULL;
}

bool enQueue2(LinkQueue* q, ElemType e) {
    Node* tmp = (Node*)malloc(sizeof(Node));
    tmp->data = e;
    tmp->next = NULL;
    q->rear->next = tmp;
    q->rear = tmp;
    return true;
}

bool deQueue2(LinkQueue* q, ElemType* e) {
    if (empty2(*q)) {
        return false;
    }

    Node* old = q->front->next;
    *e = old->data;
    q->front->next = old->next;
    if (q->front->next == NULL) {
        q->rear = q->front;
    }
    free(old);
    return true;
}

bool front2(LinkQueue* q, ElemType* e) {
    if (empty2(*q)) {
        return false;
    }

    *e = q->front->next->data;
    return true;
}

void printqueue2(LinkQueue q, char* s) {
    s[0] = '\0';
    for (Node* n = q.front->next; n != NULL; n = n->next) {
        char tmp[10];
        sprintf(tmp, "%d-", n->data);
        strcat(s, tmp);
    }
    printf("%s\n", s);
}

typedef struct queue {
    void (*InitQueue)(LinkQueue* q);
    void (*destory)(LinkQueue* q);
    bool (*empty)(LinkQueue q);
    bool (*enQueue)(LinkQueue* q, ElemType e);
    bool (*deQueue)(LinkQueue* q, ElemType* e);
    bool (*front)(LinkQueue* q, ElemType* e);
    void (*printqueue)(LinkQueue q, char* s);
} queue;

void testing(queue qq) {
    LinkQueue q;
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
    assert(qq.enQueue(&q, 1));
    qq.printqueue(q, s);
    assert(!strcmp(s, "3-2-1-1-"));

    qq.destory(&q);
}

void test1() {
    queue qq;

    qq.InitQueue = InitQueue1;
    qq.empty = empty1;
    qq.enQueue = enQueue1;
    qq.deQueue = deQueue1;
    qq.front = front1;
    qq.printqueue = printqueue1;
    qq.destory = destory;

    testing(qq);
}

void test2() {
    queue qq;

    qq.InitQueue = InitQueue2;
    qq.empty = empty2;
    qq.enQueue = enQueue2;
    qq.deQueue = deQueue2;
    qq.front = front2;
    qq.printqueue = printqueue2;
    qq.destory = destory;

    testing(qq);
}

int main(int argc, char const* argv[]) {
    // test1();
    test2();
    return 0;
}
