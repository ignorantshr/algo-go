#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include "stack.h"

typedef struct Node {
    ElemType data;
    struct Node* next;
} Node, *ListStack;

// 无头节点，比较方便
void initListStack(ListStack s) {
    return;
}

void initHListStack(ListStack s) {
    Node* n = (Node*)malloc(sizeof(Node));
    n->data = 0;
    n->next = NULL;
    *s = *n;
}

void destroy(ListStack s) {
    for (Node* n = s; n != NULL;) {
        Node* next = s->next;
        free(n);
        n = next;
    }
}

/*
 *在
 *C语言中，指针本身也是可以传递的。当需要修改指针本身的值时，我们可以通过传递指针的指针或指针的引用来实现。
 */
// 头插法，比较方便
bool push(ListStack* s, ElemType e) {
    Node* tmp = (Node*)malloc(sizeof(Node));
    tmp->data = e;
    if (*s == NULL) {
        tmp->next = NULL;
        *s = tmp;
    } else {
        tmp->next = *s;
        *s = tmp;
    }

    return true;
}

bool pushH(ListStack s, ElemType e) {
    Node* tmp = (Node*)malloc(sizeof(Node));
    tmp->data = e;
    tmp->next = (s)->next;
    (s)->next = tmp;
    return true;
}

bool pop(ListStack* s, ElemType* e) {
    if (*s == NULL) {
        return false;
    }

    Node* tmp = (*s);
    *e = tmp->data;
    *s = tmp->next;
    tmp->next = NULL;
    free(tmp);
    return true;
}

bool popH(ListStack s, ElemType* e) {
    if ((s)->next == NULL) {
        return false;
    }

    Node* tmp = (s)->next;
    *e = tmp->data;
    (s)->next = tmp->next;
    tmp->next = NULL;
    free(tmp);
    return true;
}

// popH(s, &e); 这种写法的函数形式
// bool popH(ListStack* s, ElemType* e) {
//     if ((*s)->next == NULL) {
//         return false;
//     }

//     Node* tmp = (*s)->next;
//     *e = tmp->data;
//     (*s)->next = tmp->next;
//     tmp->next = NULL;
//     free(tmp);
//     return true;
// }

bool top(ListStack s, ElemType* e) {
    if (s == NULL) {
        return false;
    }

    *e = s->data;
    return true;
}

bool topH(ListStack s, ElemType* e) {
    if (s->next == NULL) {
        return false;
    }

    *e = s->next->data;
    return true;
}

void test() {
    ListStack s = NULL;
    initListStack(s);
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

    pop(&s, &e);
    top(s, &e);
    printf("%d\n", e);

    pop(&s, &e);
    top(s, &e);
    printf("%d\n", e);
}

void testHead() {
    ListStack s = (Node*)malloc(sizeof(Node));
    initHListStack(s);
    ElemType e;

    pushH(s, 1);
    topH(s, &e);
    printf("%d\n", e);

    pushH(s, 2);
    topH(s, &e);
    printf("%d\n", e);

    popH(s, &e);
    topH(s, &e);
    printf("%d\n", e);

    popH(s, &e);
    topH(s, &e);
    printf("%d\n", e);

    popH(s, &e);
    topH(s, &e);
    printf("%d\n", e);
}

// int main(int argc, char const* argv[]) {
//     test();
//     // testHead();
//     return 0;
// }
