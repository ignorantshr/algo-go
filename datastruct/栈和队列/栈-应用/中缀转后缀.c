#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../链栈.c"

void infix2suffix(char* s, char* res) {
    ListStack numstack = NULL;
    ListStack opstack = NULL;
    initListStack(numstack);
    initListStack(opstack);

    int len = strlen(s);
    char numStr[20];
    int j = 0;

    for (int i = 0; i < len; i++) {
        char cur = s[i];
        if (cur >= '0' && cur <= '9') {
            numStr[j++] = cur;
        } else {
            if (j > 0) {
                numStr[j] = '\0';
                j = 0;
                push(&numstack, numStr[0]);
            }
            ElemType e;
            switch (cur) {
                case '(':
                case '+':
                case '-':
                case '*':
                case '/':
                    push(&opstack, cur);
                    break;
                case ')':
                    while (pop(&opstack, &e) && e != '(') {
                        push(&numstack, e);
                    };
                    break;
                default:
                    break;
            }
        }
    }

    if (j > 0) {
        numStr[j] = '\0';
        j = 0;
        push(&numstack, numStr[0]);
    }
    int i = 0;
    for (Node* n = numstack; n != NULL; n = n->next) {
        if (n != NULL) {
            char c = n->data;
            res[i++] = c;
            // if (c < '0' || c > '9') {
            // } else {
            //     res[i++] = '0' + n->data;
            // }
        }
    }
    for (int j = 0; j < i >> 1; j++) {
        char tmp = res[i - j - 1];
        res[i - j - 1] = res[j];
        res[j] = tmp;
    }
    for (Node* n = opstack; n != NULL; n = n->next) {
        if (n != NULL) {
            res[i++] = n->data;
        }
    }
    res[i] = '\0';
}

int main(int argc, char const* argv[]) {
    char res[100];

    infix2suffix("1+1", res);
    assert(!strcmp("11+", res));

    infix2suffix("1-1", res);
    assert(!strcmp("11-", res));

    infix2suffix("1*1", res);
    assert(!strcmp("11*", res));

    infix2suffix("1/1", res);
    assert(!strcmp("11/", res));

    infix2suffix("1+1*(2-3)", res);
    assert(!strcmp("1123-*+", res));

    infix2suffix("((5/(7-(1+1)))*3)-(2+(1+1))", res);
    assert(!strcmp("5711+-/3*211++-", res));
    return 0;
}
