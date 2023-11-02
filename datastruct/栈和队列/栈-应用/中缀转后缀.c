#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../链栈.c"

// 也叫逆波兰式
// 按照“左优先”原则确定的这些运算符的生效顺序和后缀表达式中各个运算符从左到右出现的次序是相同的

void infix2suffix(char* s, char* res) {
    res[0] = '\0';
    ListStack opstack = NULL;
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
                strcat(res, numStr);
            }
            ElemType e;
            ElemType topv;
            switch (cur) {
                case '+':
                case '-':
                    // 弹出优先级 >=自己的，然后入栈。
                    // 因为之前入栈的肯定比自己的运行顺序靠前
                    while (top(opstack, &topv) && topv == '-' || topv == '+') {
                        pop(&opstack, &topv);
                        char str[20];  // 存储转换后的字符串
                        snprintf(str, sizeof(str), "%c", topv);
                        strcat(res, str);
                    }
                    push(&opstack, cur);
                    break;
                case '*':
                case '/':
                    while (top(opstack, &topv) && topv == '/' || topv == '*') {
                        pop(&opstack, &topv);
                        char str[20];  // 存储转换后的字符串
                        snprintf(str, sizeof(str), "%c", topv);
                        strcat(res, str);
                    }
                    push(&opstack, cur);
                    break;
                case '(':
                    push(&opstack, cur);
                    break;
                case ')':
                    while (pop(&opstack, &e) && e != '(') {
                        char str[20];  // 存储转换后的字符串
                        snprintf(str, sizeof(str), "%c", e);
                        strcat(res, str);
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
        strcat(res, numStr);
    }
    for (Node* n = opstack; n != NULL; n = n->next) {
        if (n != NULL) {
            char str[20];  // 存储转换后的字符串
            snprintf(str, sizeof(str), "%d", n->data);
            strcat(res, str);
        }
    }
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
