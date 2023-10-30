#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include "stack.h"

#define MaxSize 3

// 共享栈，两个栈共用一个栈空间
typedef struct ShStack {
    ElemType data[MaxSize];
    int btop;  // 底部栈顶
    int ttop;  // 顶部栈顶
} ShStack;

void init(ShStack* s) {
    s->btop = -1;
    s->ttop = MaxSize;
}

bool full(ShStack s) {
    return s.btop + 1 == s.ttop;
}

bool push2b(ShStack* s, ElemType e) {
    if (full(*s)) {
        return false;
    }

    s->data[++s->btop] = e;
    return true;
}

bool push2t(ShStack* s, ElemType e) {
    if (full(*s)) {
        return false;
    }

    s->data[--s->ttop] = e;
    return true;
}

bool pop4b(ShStack* s, ElemType* e) {
    if (s->btop == -1) {
        return false;
    }

    *e = s->data[s->btop--];
    return true;
}

bool pop4t(ShStack* s, ElemType* e) {
    if (s->ttop == MaxSize) {
        return false;
    }

    *e = s->data[s->ttop++];
    return true;
}

bool top4b(ShStack s, ElemType* e) {
    if (s.btop == -1) {
        return false;
    }

    *e = s.data[s.btop];
    return true;
}

bool top4t(ShStack s, ElemType* e) {
    if (s.ttop == MaxSize) {
        return false;
    }

    *e = s.data[s.ttop];
    return true;
}

int main(int argc, char const* argv[]) {
    ShStack s;
    int e;
    init(&s);

    push2b(&s, 1);
    top4b(s, &e);
    printf("%d\n", e);

    push2t(&s, 2);
    top4t(s, &e);
    printf("%d\n", e);

    push2t(&s, 3);
    top4t(s, &e);
    printf("%d\n", e);

    if (push2t(&s, 4)) {
        top4t(s, &e);
        printf("%d\n", e);
    }

    pop4b(&s, &e);
    printf("%d\n", e);

    if (push2t(&s, 4)) {
        top4t(s, &e);
        printf("%d\n", e);
    }

    pop4t(&s, &e);
    pop4t(&s, &e);
    pop4t(&s, &e);
    printf("%d\n", e);
    return 0;
}
