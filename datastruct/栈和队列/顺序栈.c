#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include "stack.h"

#define MaxSize 10

// 创、增、删、查 都是 O(1)
typedef struct SeqStack {
    ElemType data[MaxSize];
    int top;
} SeqStack;

void init(SeqStack* s) {
    s->top = -1;
}

void destroy(SeqStack* s) {
    // 系统自动回收
    s->top = -1;
}

bool push(SeqStack* s, ElemType e) {
    if (s->top == MaxSize - 1) {
        return false;
    }

    s->data[++s->top] = e;
    return true;
}

bool pop(SeqStack* s, ElemType* e) {
    if (s->top == -1) {
        return false;
    }

    *e = s->data[s->top--];
    return true;
}

bool top(SeqStack s, ElemType* e) {
    if (s.top == -1) {
        return false;
    }

    *e = s.data[s.top];
    return true;
}

int main(int argc, char const* argv[]) {
    SeqStack s;
    init(&s);
    ElemType e;

    push(&s, 1);
    top(s, &e);
    printf("%d\n", e);

    push(&s, 2);
    top(s, &e);
    printf("%d\n", e);

    pop(&s, &e);
    top(s, &e);
    printf("%d\n", e);
    return 0;
}
